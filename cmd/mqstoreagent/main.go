package main

import (
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
	"gopkg.in/avro.v0"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/covalenthq/mq-store-agent/internal/config"
	"github.com/covalenthq/mq-store-agent/internal/event"
	"github.com/covalenthq/mq-store-agent/internal/handler"
	"github.com/covalenthq/mq-store-agent/internal/utils"
)

var (
	waitGrp sync.WaitGroup

	avroCodecs []*goavro.Codec

	ConsumeEvents       int64 = 1
	consumerIdleTime    int64 = 30
	consumerPendingTime int64 = 60

	start               string = ">"
	CodecPath           string = "./codec/"
	RedisURL            string
	streamKey           string
	consumerGroup       string
	specimenSegmentName string
	resultSegmentName   string

	specimenSegmentIdBatch []string
	resultSegmentIdBatch   []string

	specimenSegment event.SpecimenSegment
	resultSegment   event.ResultSegment
)

func init() {
	formatter := runtime.Formatter{ChildFormatter: &log.TextFormatter{
		FullTimestamp: true,
	}}
	formatter.Line = true
	log.SetFormatter(&formatter)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.WithFields(log.Fields{"file": "main.go"}).Info("Server is running...")
}

func main() {
	flag.StringVar(&RedisURL, "redis-url", utils.LookupEnvOrString("RedisURL", RedisURL), "redis consumer stream url")
	flag.Parse()

	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	log.Info("Agent config: ", utils.GetConfig(flag.CommandLine))

	redisClient, streamKey, consumerGroup, err := utils.NewRedisClient(utils.LookupEnvOrString("RedisURL", RedisURL))
	if err != nil {
		panic(err)
	}

	storageClient, err := utils.NewStorageClient(&config.GcpConfig)
	if err != nil {
		panic(err)
	}

	ethProofClient, err := utils.NewEthClient(config.EthConfig.ProofClient)
	if err != nil {
		panic(err)
	}

	specimenAvro, err := avro.ParseSchemaFile(CodecPath + "block-specimen.avsc")
	if err != nil {
		log.Fatalf("unable to parse avro schema for specimen: %v", err)
	}

	specimenCodec, err := goavro.NewCodec(specimenAvro.String())
	if err != nil {
		log.Fatalf("unable to gen avro codec for specimen: %v", err)
	}

	resultAvro, err := avro.ParseSchemaFile(CodecPath + "block-result.avsc")
	if err != nil {
		log.Fatalf("unable to parse avro schema for result: %v", err)
	}

	resultCodec, err := goavro.NewCodec(resultAvro.String())
	if err != nil {
		log.Fatalf("unable to parse avro schema for result: %v", err)
	}

	avroCodecs := append(avroCodecs, specimenCodec, resultCodec)

	var consumerName string = uuid.NewV4().String()

	log.Printf("Initializing Consumer: %v | Redis Stream: %v | Consumer Group: %v", consumerName, streamKey, consumerGroup)

	createConsumerGroup(redisClient, streamKey, consumerGroup)

	go consumeEvents(config, avroCodecs, redisClient, storageClient, ethProofClient, consumerName, streamKey, consumerGroup)
	go consumePendingEvents(config, avroCodecs, redisClient, storageClient, ethProofClient, consumerName, streamKey, consumerGroup)

	//Gracefully disconnect
	chanOS := make(chan os.Signal, 1)
	signal.Notify(chanOS, syscall.SIGINT, syscall.SIGTERM)
	<-chanOS

	waitGrp.Wait()
	redisClient.Close()
	storageClient.Close()
	ethProofClient.Close()
}

func createConsumerGroup(redisClient *redis.Client, streamKey, consumerGroup string) {
	if _, err := redisClient.XGroupCreateMkStream(streamKey, consumerGroup, "0").Result(); err != nil {
		if !strings.Contains(fmt.Sprint(err), "BUSYGROUP") {
			log.Printf("Error on create Consumer Group: %v ...\n", consumerGroup)
			panic(err)
		}
	}
}

func consumeEvents(config *config.Config, avroCodecs []*goavro.Codec, redisClient *redis.Client, storage *storage.Client, ethProof *ethclient.Client, consumerName, streamKey, consumerGroup string) {
	for {
		log.Debug("New sequential stream unit: ", time.Now().Format(time.RFC3339))
		streams, err := redisClient.XReadGroup(&redis.XReadGroupArgs{
			Streams:  []string{streamKey, start},
			Group:    consumerGroup,
			Consumer: consumerName,
			Count:    ConsumeEvents,
			Block:    0,
		}).Result()

		if err != nil {
			log.Error("err on consume events: ", err.Error())
			return
		}

		for _, stream := range streams[0].Messages {
			waitGrp.Add(1)
			go processStream(config, avroCodecs, redisClient, storage, ethProof, stream, false, handler.HandlerFactory())
		}
		waitGrp.Wait()
	}
}

func consumePendingEvents(config *config.Config, avroCodecs []*goavro.Codec, redisClient *redis.Client, storage *storage.Client, ethProof *ethclient.Client, consumerName, streamKey, consumerGroup string) {
	ticker := time.Tick(time.Second * time.Duration(consumerPendingTime))
	for range ticker {
		var streamsRetry []string
		pendingStreams, err := redisClient.XPendingExt(&redis.XPendingExtArgs{
			Stream: streamKey,
			Group:  consumerGroup,
			Start:  "0",
			End:    "+",
			Count:  ConsumeEvents,
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
				MinIdle:  time.Duration(consumerIdleTime) * time.Second,
			}).Result()

			if err != nil {
				log.Error("error on process pending: ", err.Error())
				return
			}

			for _, stream := range streams {
				waitGrp.Add(1)
				go processStream(config, avroCodecs, redisClient, storage, ethProof, stream, true, handler.HandlerFactory())
			}
			waitGrp.Wait()
		}
		log.Info("Process pending streams at: ", time.Now().Format(time.RFC3339))
	}
}

func processStream(config *config.Config, avroCodecs []*goavro.Codec, redisClient *redis.Client, storage *storage.Client, ethProof *ethclient.Client, stream redis.XMessage, retry bool, handlerFactory func(t event.Type) handler.Handler) {
	defer waitGrp.Done()

	ctx := context.Background()
	typeEvent := stream.Values["type"].(string)
	hash := stream.Values["hash"].(string)

	decodedData, err := snappy.Decode(nil, []byte(stream.Values["data"].(string)))
	if err != nil {
		log.Info("Failed to snappy decode: ", err.Error())
	}

	newEvent, _ := event.New(event.Type(typeEvent))
	newEvent.SetID(stream.ID)

	h := handlerFactory(event.Type(typeEvent))
	specimen, result, err := h.Handle(config, storage, ethProof, newEvent, hash, decodedData, retry)
	if err != nil {
		log.Fatalf("error: ", err.Error(), " on process event: ", newEvent)
	} else {
		if specimen == nil {
			// collect stream ids and block results
			resultSegmentIdBatch = append(resultSegmentIdBatch, stream.ID)
			resultSegment.BlockResult = append(resultSegment.BlockResult, result)
			if len(resultSegment.BlockResult) == 1 {
				resultSegment.StartBlock = result.Data.Header.Number.Uint64()
			}
			if len(resultSegment.BlockResult) == int(config.GeneralConfig.SegmentLength) {
				resultSegment.EndBlock = result.Data.Header.Number.Uint64()
				resultSegment.Elements = uint64(config.GeneralConfig.SegmentLength)
				resultSegmentName = fmt.Sprint(resultSegment.StartBlock) + "-" + fmt.Sprint(resultSegment.EndBlock)
				// encode, prove and upload
				_, err := handler.EncodeProveAndUploadResultSegment(ctx, config, avroCodecs[1], &resultSegment, resultSegmentName, storage, ethProof)
				if err != nil {
					log.Fatalf("failed to avro encode, proove and upload block-result segment: %v with err: %v", resultSegmentName, err)
				}
				//ack stream segment batch id
				err = utils.AckStreamSegment(config, redisClient, streamKey, consumerGroup, resultSegmentIdBatch)
				if err != nil {
					log.Fatalf("failed to match streamIDs length to segment length config: %v", err)
				}
				// reset segment and name
				resultSegment = event.ResultSegment{}
				resultSegmentName = ""
				resultSegmentIdBatch = []string{}
			}
		} else {
			// collect stream ids and block specimens
			specimenSegmentIdBatch = append(specimenSegmentIdBatch, stream.ID)
			specimenSegment.BlockSpecimen = append(specimenSegment.BlockSpecimen, specimen)
			if len(specimenSegment.BlockSpecimen) == 1 {
				specimenSegment.StartBlock = specimen.Data.Header.Number.Uint64()
				println(specimenSegment.StartBlock, "start block")
			}
			if len(specimenSegment.BlockSpecimen) == int(config.GeneralConfig.SegmentLength) {
				specimenSegment.EndBlock = specimen.Data.Header.Number.Uint64()
				specimenSegment.Elements = uint64(config.GeneralConfig.SegmentLength)
				specimenSegmentName = fmt.Sprint(specimenSegment.StartBlock) + "-" + fmt.Sprint(specimenSegment.EndBlock)
				println(specimenSegment.EndBlock, "end block")
				// encode, prove and upload
				_, err := handler.EncodeProveAndUploadSpecimenSegment(ctx, config, avroCodecs[0], &specimenSegment, specimenSegmentName, storage, ethProof)
				if err != nil {
					log.Fatalf("failed to avro encode, proove and upload block-specimen segment: %v with err: %v", resultSegmentName, err)
				}
				//ack stream segment batch id
				err = utils.AckStreamSegment(config, redisClient, streamKey, consumerGroup, specimenSegmentIdBatch)
				if err != nil {
					log.Fatalf("failed to match streamIDs length to segment length config: %v", err)
				}
				// reset segment and name
				specimenSegment = event.SpecimenSegment{}
				specimenSegmentName = ""
				specimenSegmentIdBatch = []string{}
			}
		}
	}
}
