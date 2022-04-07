package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"path"
	"strings"
	"sync"
	"syscall"
	"time"

	"cloud.google.com/go/storage"
	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v7"
	"github.com/golang/snappy"
	"github.com/linkedin/goavro/v2"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/ubiq/go-ubiq/rlp"
	"gopkg.in/avro.v0"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/covalenthq/bsp-agent/internal/config"
	"github.com/covalenthq/bsp-agent/internal/event"
	"github.com/covalenthq/bsp-agent/internal/handler"
	"github.com/covalenthq/bsp-agent/internal/proof"
	"github.com/covalenthq/bsp-agent/internal/types"
	"github.com/covalenthq/bsp-agent/internal/utils"
	"github.com/covalenthq/bsp-agent/internal/websocket"
)

var (
	waitGrp sync.WaitGroup
	// consts
	consumerEvents            int64 = 1
	consumerPendingIdleTime   int64 = 30
	consumerPendingTimeTicker int64 = 120

	// env flag vars
	consumerPendingTimeoutFlag = 60 // defaults to 1 min
	segmentLengthFlag          = 1  // defaults to 1 block per segment
	packerLengthFlag           = 10 // defaults to 10 block per pack tx
	avroCodecPathFlag          string
	redisURLFlag               string
	replicaBucketFlag          string
	gcpSvcAccountFlag          string
	proofChainFlag             string
	binaryFilePathFlag         string
	websocketURLsFlag          string
	logFolderFlag              = "./logs/"

	// stream processing vars
	start                 = ">"
	streamKey             string
	consumerGroup         string
	replicaSegmentName    string
	replicaSegmentIDBatch []string
	replicationSegment    event.ReplicationSegment
	packerTxBatch         []string
	blockReplica          types.BlockReplica
)

func parseFlags() {
	flag.StringVar(&redisURLFlag, "redis-url", utils.LookupEnvOrString("RedisURL", redisURLFlag), "redis consumer stream url")
	flag.StringVar(&avroCodecPathFlag, "avro-codec-path", utils.LookupEnvOrString("CodecPath", avroCodecPathFlag), "local path to AVRO .avsc files housing the specimen/result schemas")
	flag.StringVar(&binaryFilePathFlag, "binary-file-path", utils.LookupEnvOrString("BinaryFilePath", binaryFilePathFlag), "local path to AVRO encoded binary files that contain block-replicas")
	flag.StringVar(&gcpSvcAccountFlag, "gcp-svc-account", utils.LookupEnvOrString("GcpSvcAccount", gcpSvcAccountFlag), "local path to google cloud platform service account auth file")
	flag.StringVar(&replicaBucketFlag, "replica-bucket", utils.LookupEnvOrString("ReplicaBucket", replicaBucketFlag), "google cloud platform object store target for specimen")
	flag.StringVar(&proofChainFlag, "proof-chain-address", utils.LookupEnvOrString("ProofChain", proofChainFlag), "hex string address for deployed proof-chain contract")
	flag.StringVar(&websocketURLsFlag, "websocket-urls", utils.LookupEnvOrString("WebsocketURLs", websocketURLsFlag), "url to websockets clients separated by space")
	flag.IntVar(&segmentLengthFlag, "segment-length", utils.LookupEnvOrInt("SegmentLength", segmentLengthFlag), "number of block specimen/results within a single uploaded avro encoded object")
	flag.IntVar(&packerLengthFlag, "packer-length", utils.LookupEnvOrInt("PackerLength", packerLengthFlag), "number of block specimen within a packer tx object")
	flag.IntVar(&consumerPendingTimeoutFlag, "consumer-timeout", utils.LookupEnvOrInt("ConsumerPendingTimeout", consumerPendingTimeoutFlag), "number of seconds to wait before pending messages consumer timeout")
	flag.StringVar(&logFolderFlag, "log-folder", utils.LookupEnvOrString("LogFolder", logFolderFlag), "Location where the log files should be placed")
	flag.Parse()
}

func init() {
	parseFlags()

	// setup logger
	formatter := runtime.Formatter{ChildFormatter: &log.TextFormatter{
		FullTimestamp: true,
	}}
	formatter.Line = true
	log.SetFormatter(&formatter)

	var outWriter io.Writer
	logLocationURL, err := utils.GetLogLocationURL(logFolderFlag)
	if err != nil {
		log.Warn("error while setting up file logging: ", err)
		outWriter = os.Stdout
	} else {
		logFilePath := path.Join(logLocationURL.Path, "log.log")
		bspLogger := utils.NewLoggerOut(os.Stdout, &lumberjack.Logger{
			// logs folder created/searched in directory in which agent was started.
			Filename:   logFilePath,
			MaxSize:    100, // megabytes
			MaxBackups: 7,
			MaxAge:     10, // days
		})
		outWriter = &bspLogger
	}

	log.SetOutput(outWriter)
	log.SetLevel(log.InfoLevel)
	log.WithFields(log.Fields{"file": "main.go"}).Info("bsp-agent is running...")
}

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	log.Info("bsp-agent command line config: ", utils.GetConfig(flag.CommandLine))

	if binaryFilePathFlag == "" {
		log.Warn("--binary-file-path flag not provided to write block-replica avro encoded binary files to local path", binaryFilePathFlag)
	}

	redisClient, streamKey, consumerGroup, err := utils.NewRedisClient(utils.LookupEnvOrString("RedisURL", redisURLFlag), &config.RedisConfig)
	if err != nil {
		log.Fatalf("unable to get redis client from redis URL flag: %v", err)
	}
	storageClient, err := utils.NewStorageClient(gcpSvcAccountFlag)
	if err != nil {
		log.Printf("unable to get gcp storage client; --gcp-svc-account flag not set or set incorrectly: %v, storing BSP files locally: %v", err, utils.LookupEnvOrString("BinaryFilePath", binaryFilePathFlag))
	}
	ethClient, err := utils.NewEthClient(config.EthConfig.CqtRPCURL)
	if err != nil {
		log.Fatalf("unable to get ethereum client from Eth client flag: %v", err)
	}
	replicaAvro, err := avro.ParseSchemaFile(avroCodecPathFlag)
	if err != nil {
		log.Fatalf("unable to parse avro schema for specimen from codec path flag: %v", err)
	}
	replicaCodec, err := goavro.NewCodec(replicaAvro.String())
	if err != nil {
		log.Fatalf("unable to generate avro codec for block-replica: %v", err)
	}

	if websocketURLsFlag != "" {
		websocketsURLs := strings.Split(websocketURLsFlag, " ")
		for _, url := range websocketsURLs {
			go websocket.ConsumeWebsocketsEvents(&config.EthConfig, url, replicaCodec, ethClient, storageClient, binaryFilePathFlag, replicaBucketFlag, proofChainFlag)
		}
	} else {
		var consumerName string = uuid.NewV4().String()
		log.Printf("Initializing Consumer: %v | Redis Stream: %v | Consumer Group: %v", consumerName, streamKey, consumerGroup)
		createConsumerGroup(redisClient, streamKey, consumerGroup)
		go consumeEvents(config, replicaCodec, redisClient, storageClient, ethClient, consumerName, streamKey, consumerGroup)
		go consumePendingEvents(config, replicaCodec, redisClient, storageClient, ethClient, consumerName, streamKey, consumerGroup)
	}

	// Gracefully disconnect
	chanOS := make(chan os.Signal, 1)
	signal.Notify(chanOS, syscall.SIGINT, syscall.SIGTERM)
	<-chanOS

	log.Warn("Received interrupt. Flushing in-memory blocks...")
	waitGrp.Wait()
	log.Warn("waitgrp ended. Closing agent...")

	if redisClient != nil {
		err = redisClient.Close()
		if err != nil {
			log.Error("error in closing redis client: ", err)
		}
	}

	if storageClient != nil {
		err = storageClient.Close()
		if err != nil {
			log.Error("error in closing storage client: ", err)
		}
	}
	ethClient.Close()
}

func createConsumerGroup(redisClient *redis.Client, streamKey, consumerGroup string) {
	if _, err := redisClient.XGroupCreateMkStream(streamKey, consumerGroup, "0").Result(); err != nil {
		if !strings.Contains(fmt.Sprint(err), "BUSYGROUP") {
			log.Printf("Error on create Consumer Group: %v ...\n", consumerGroup)
			panic(err)
		}
	}
}

func consumeEvents(config *config.Config, avroCodecs *goavro.Codec, redisClient *redis.Client, storageClient *storage.Client, ethProof *ethclient.Client, consumerName, streamKey, consumerGroup string) {
	for {
		log.Debug("New sequential stream unit: ", time.Now().Format(time.RFC3339))
		streams, err := redisClient.XReadGroup(&redis.XReadGroupArgs{
			Streams:  []string{streamKey, start},
			Group:    consumerGroup,
			Consumer: consumerName,
			Count:    consumerEvents,
			Block:    0,
		}).Result()
		if err != nil {
			log.Error("error on consume events: ", err.Error())

			return
		}

		for _, stream := range streams[0].Messages {
			waitGrp.Add(1)
			go processStream(config, avroCodecs, redisClient, storageClient, ethProof, stream)
		}
		waitGrp.Wait()
	}
}

func consumePendingEvents(config *config.Config, avroCodecs *goavro.Codec, redisClient *redis.Client, storageClient *storage.Client, ethClient *ethclient.Client, consumerName, streamKey, consumerGroup string) {
	timeout := time.After(time.Second * time.Duration(consumerPendingTimeoutFlag))
	ticker := time.Tick(time.Second * time.Duration(consumerPendingTimeTicker))
	for {
		select {
		case <-timeout:
			log.Info("Process pending streams stopped at: ", time.Now().Format(time.RFC3339), " after timeout: ", consumerPendingTimeoutFlag, " seconds")
			os.Exit(0)
		case <-ticker:
			var streamsRetry []string
			pendingStreams, err := redisClient.XPendingExt(&redis.XPendingExtArgs{
				Stream: streamKey,
				Group:  consumerGroup,
				Start:  "0",
				End:    "+",
				Count:  consumerEvents,
			}).Result()
			if err != nil {
				panic(err)
			}

			for _, stream := range pendingStreams {
				streamsRetry = append(streamsRetry, stream.ID)
			}
			if len(streamsRetry) > 0 {
				streams, err := redisClient.XClaim(&redis.XClaimArgs{
					Stream:   streamKey,
					Group:    consumerGroup,
					Consumer: consumerName,
					Messages: streamsRetry,
					MinIdle:  time.Duration(consumerPendingIdleTime) * time.Second,
				}).Result()
				if err != nil {
					log.Error("error on process pending: ", err.Error())

					return
				}
				for _, stream := range streams {
					waitGrp.Add(1)
					go processStream(config, avroCodecs, redisClient, storageClient, ethClient, stream)
				}
				waitGrp.Wait()
			}
			log.Info("Process pending streams at: ", time.Now().Format(time.RFC3339))
		}
	}
}
func processStream(config *config.Config, replicaCodec *goavro.Codec, redisClient *redis.Client, storageClient *storage.Client, ethClient *ethclient.Client, stream redis.XMessage) {
	ctx := context.Background()
	hash := stream.Values["hash"].(string)
	decodedData, err := snappy.Decode(nil, []byte(stream.Values["data"].(string)))
	if err != nil {
		log.Info("Failed to snappy decode: ", err.Error())
	}

	err = rlp.Decode(bytes.NewReader(decodedData), &blockReplica)
	if err != nil {
		log.Error("error decoding RLP bytes to block-replica: ", err)
	}
	defer waitGrp.Done()

	newEvent, _ := event.NewBlockReplicaEvent()
	replica, err := handler.ParseStreamToEvent(newEvent, hash, &blockReplica)
	objectType := blockReplica.Type[5:]
	objectReplica := &blockReplica
	if err != nil {
		log.Error("error on process event: ", err)
	} else {
		// collect stream ids and block replicas
		replicaSegmentIDBatch = append(replicaSegmentIDBatch, stream.ID)
		replicationSegment.BlockReplicaEvent = append(replicationSegment.BlockReplicaEvent, replica)
		if len(replicationSegment.BlockReplicaEvent) == 1 {
			replicationSegment.StartBlock = replica.Data.Header.Number.Uint64()
		}
		if len(replicationSegment.BlockReplicaEvent) == segmentLengthFlag {
			replicationSegment.EndBlock = replica.Data.Header.Number.Uint64()
			replicationSegment.Elements = uint64(segmentLengthFlag)
			replicaSegmentName = fmt.Sprint(replica.Data.NetworkId) + "-" + fmt.Sprint(replicationSegment.StartBlock) + objectType
			// avro encode, prove and upload
			specimenTxHash, err := handler.EncodeProveAndUploadReplicaSegment(ctx, &config.EthConfig, replicaCodec, &replicationSegment, objectReplica, storageClient, ethClient, binaryFilePathFlag, replicaBucketFlag, replicaSegmentName, proofChainFlag)
			if err != nil {
				log.Error("failed to avro encode, prove and upload block-result segment with err: ", err)
			}
			// ack stream segment batch id
			err = utils.AckStreamSegment(config, redisClient, segmentLengthFlag, streamKey, consumerGroup, replicaSegmentIDBatch)
			if err != nil {
				log.Error("failed to match streamIDs length to segment length config: ", err)
			}
			packerTxBatch = append(packerTxBatch, specimenTxHash)

			if len(packerTxBatch) == packerLengthFlag {
				fmt.Printf("\n---> Sending Tx to Moonbeam for %v Proof Txs %v <---\n", packerLengthFlag, packerTxBatch)

				_, err := proof.SendPackerProofTx(ctx, &config.EthConfig, ethClient, packerTxBatch)
				if err != nil {
					log.Error("failed to sent Packer Tx: ", err)
					panic(err)
				}
				packerTxBatch = []string{}
			}
			// reset segment, name, id batch stores
			replicationSegment = event.ReplicationSegment{}
			replicaSegmentName = ""
			replicaSegmentIDBatch = []string{}
		}
	}
}
