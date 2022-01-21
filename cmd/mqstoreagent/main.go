package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"cloud.google.com/go/storage"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v7"
	"github.com/golang/snappy"
	"github.com/linkedin/goavro/v2"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/ubiq/go-ubiq/rlp"
	"gopkg.in/avro.v0"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"

	"github.com/covalenthq/mq-store-agent/internal/config"
	"github.com/covalenthq/mq-store-agent/internal/event"
	"github.com/covalenthq/mq-store-agent/internal/handler"
	"github.com/covalenthq/mq-store-agent/internal/types"
	"github.com/covalenthq/mq-store-agent/internal/utils"
	"github.com/covalenthq/mq-store-agent/internal/websocket"
)

var (
	waitGrp sync.WaitGroup
	// consts
	consumerEvents            int64 = 1
	consumerPendingIdleTime   int64 = 30
	consumerPendingTimeTicker int64 = 120

	// env flags
	ConsumerPendingTimeoutFlag = 60 // defaults to 1 mins
	SegmentLengthFlag          = 5  // defaults to 5 blocks per segment

	CodecPathFlag      string
	RedisURLFlag       string
	ReplicaBucketFlag  string
	GcpSvcAccountFlag  string
	EthClientFlag      string
	ProofChainFlag     string
	BinaryFilePathFlag string
	WebsocketURLsFlag  string

	start                 = ">"
	streamKey             string
	consumerGroup         string
	replicaSegmentName    string
	replicaSegmentIDBatch []string

	replicationSegment event.ReplicationSegment
	blockReplica       types.BlockReplica
)

func init() {
	formatter := runtime.Formatter{ChildFormatter: &log.TextFormatter{
		FullTimestamp: true,
	}}
	formatter.Line = true
	log.SetFormatter(&formatter)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.WithFields(log.Fields{"file": "main.go"}).Info("mq-store-agent is running...")
}

func main() {
	flag.StringVar(&RedisURLFlag, "redis-url", utils.LookupEnvOrString("RedisURL", RedisURLFlag), "redis consumer stream url")
	flag.StringVar(&CodecPathFlag, "codec-path", utils.LookupEnvOrString("CodecPath", CodecPathFlag), "local path to AVRO .avsc files housing the specimen/result schemas")
	flag.StringVar(&BinaryFilePathFlag, "binary-file-path", utils.LookupEnvOrString("BinaryFilePath", BinaryFilePathFlag), "local path to AVRO encoded binary files that contain block-replicas")
	flag.StringVar(&GcpSvcAccountFlag, "gcp-svc-account", utils.LookupEnvOrString("GcpSvcAccount", GcpSvcAccountFlag), "local path to google cloud platform service account auth file")
	flag.StringVar(&ReplicaBucketFlag, "replica-bucket", utils.LookupEnvOrString("ReplicaBucket", ReplicaBucketFlag), "google cloud platform object store target for specimen")
	flag.StringVar(&EthClientFlag, "eth-client", utils.LookupEnvOrString("EthClient", EthClientFlag), "connection string for ethereum node on which proof-chain contract is deployed")
	flag.StringVar(&ProofChainFlag, "proof-chain-address", utils.LookupEnvOrString("ProofChain", ProofChainFlag), "hex string address for deployed proof-chain contract")
	flag.StringVar(&WebsocketURLsFlag, "websocket-urls", utils.LookupEnvOrString("WebsocketURLs", WebsocketURLsFlag), "url to websockets clients separated by space")
	flag.IntVar(&SegmentLengthFlag, "segment-length", utils.LookupEnvOrInt("SegmentLength", SegmentLengthFlag), "number of block specimen/results within a single uploaded avro encoded object")
	flag.IntVar(&ConsumerPendingTimeoutFlag, "consumer-timeout", utils.LookupEnvOrInt("ConsumerPendingTimeout", ConsumerPendingTimeoutFlag), "number of seconds to wait before pending messages consumer timeout")
	flag.Parse()

	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	log.Info("Agent command line config: ", utils.GetConfig(flag.CommandLine))

	CodecPathFlag = utils.LookupEnvOrString("CodecPath", CodecPathFlag)
	BinaryFilePathFlag = utils.LookupEnvOrString("BinaryFilePath", BinaryFilePathFlag)
	ReplicaBucketFlag = utils.LookupEnvOrString("ReplicaBucket", ReplicaBucketFlag)
	GcpSvcAccountFlag = utils.LookupEnvOrString("GcpSvcAccount", GcpSvcAccountFlag)
	EthClientFlag = utils.LookupEnvOrString("EthClient", EthClientFlag)
	ProofChainFlag = utils.LookupEnvOrString("ProofChain", ProofChainFlag)
	WebsocketURLsFlag = utils.LookupEnvOrString("WebsocketURLs", WebsocketURLsFlag)

	if BinaryFilePathFlag == "" {
		log.Warn("--binary-file-path flag not provided to write block-replica avro encoded binary files to local path", BinaryFilePathFlag)
	}

	redisClient, streamKey, consumerGroup, err := utils.NewRedisClient(utils.LookupEnvOrString("RedisURL", RedisURLFlag), &config.RedisConfig)
	if err != nil {
		log.Fatalf("unable to get redis client from redis URL flag : %v", err)
	}
	storageClient, err := utils.NewStorageClient(GcpSvcAccountFlag)
	if err != nil {
		log.Warn("unable to get gcp storage client from GCP Service account flag: ", err)
	}
	ethClient, err := utils.NewEthClient(EthClientFlag)
	if err != nil {
		log.Fatalf("unable to get ethereum client from Eth client flag: %v", err)
	}
	replicaAvro, err := avro.ParseSchemaFile(CodecPathFlag)
	if err != nil {
		log.Fatalf("unable to parse avro schema for specimen from codec path flag: %v", err)
	}
	replicaCodec, err := goavro.NewCodec(replicaAvro.String())
	if err != nil {
		log.Fatalf("unable to generate avro codec for block-replica: %v", err)
	}

	if WebsocketURLsFlag != "" {
		websocketsURLs := strings.Split(WebsocketURLsFlag, " ")
		for _, url := range websocketsURLs {
			go websocket.ConsumeWebsocketsEvents(&config.EthConfig, url, replicaCodec, ethClient, storageClient, BinaryFilePathFlag, ReplicaBucketFlag, ProofChainFlag)
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

	waitGrp.Wait()

	err = redisClient.Close()
	if err != nil {
		log.Error("error in closing redis client:", err)
	}
	err = storageClient.Close()
	if err != nil {
		log.Error("error in closing storage client:", err)
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
			log.Error("err on consume events: ", err.Error())

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
	timeout := time.After(time.Second * time.Duration(ConsumerPendingTimeoutFlag))
	ticker := time.Tick(time.Second * time.Duration(consumerPendingTimeTicker))
	for {
		select {
		case <-timeout:
			log.Info("Process pending streams stopped at: ", time.Now().Format(time.RFC3339), " after timeout: ", ConsumerPendingTimeoutFlag, " seconds")
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
		log.Fatalf("error decoding RLP bytes to block-replica: %v", err)
	}
	defer waitGrp.Done()

	newEvent, _ := event.New()
	replica, err := handler.Parse(newEvent, hash, &blockReplica)
	objectType := blockReplica.Type[5:]
	if err != nil {
		log.Error("error on process event: %w", err)
	} else {
		// collect stream ids and block replicas
		replicaSegmentIDBatch = append(replicaSegmentIDBatch, stream.ID)
		replicationSegment.BlockReplicaEvent = append(replicationSegment.BlockReplicaEvent, replica)
		if len(replicationSegment.BlockReplicaEvent) == 1 {
			replicationSegment.StartBlock = replica.Data.Header.Number.Uint64()
		}
		if len(replicationSegment.BlockReplicaEvent) == SegmentLengthFlag {
			replicationSegment.EndBlock = replica.Data.Header.Number.Uint64()
			replicationSegment.Elements = uint64(SegmentLengthFlag)
			replicaSegmentName = fmt.Sprint(replica.Data.NetworkId) + "-" + fmt.Sprint(replicationSegment.StartBlock) + "-" + fmt.Sprint(replicationSegment.EndBlock) + objectType + "-" + "segment"
			// avro encode, prove and upload
			_, err := handler.EncodeProveAndUploadReplicaSegment(ctx, &config.EthConfig, replicaCodec, &replicationSegment, storageClient, ethClient, BinaryFilePathFlag, ReplicaBucketFlag, replicaSegmentName, ProofChainFlag)
			if err != nil {
				log.Error("failed to avro encode, prove and upload block-result segment with err: ", err)
			}
			// ack stream segment batch id
			err = utils.AckStreamSegment(config, redisClient, SegmentLengthFlag, streamKey, consumerGroup, replicaSegmentIDBatch)
			if err != nil {
				log.Error("failed to match streamIDs length to segment length config: ", err)
			}
			// reset segment, name, id batch stores
			replicationSegment = event.ReplicationSegment{}
			replicaSegmentName = ""
			replicaSegmentIDBatch = []string{}
		}
	}
}
