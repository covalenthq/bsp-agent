package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/go-redis/redis/v7"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/covalenthq/mq-store-agent/internal/config"
	"github.com/covalenthq/mq-store-agent/internal/event"
	"github.com/covalenthq/mq-store-agent/internal/handler"
	"github.com/covalenthq/mq-store-agent/internal/utils"
)

var (
	waitGrp            sync.WaitGroup
	client             *redis.Client
	start              string = ">"
	consumeIdleTime    int64  = 30
	consumePendingTime int64  = 60
	consumeSleepTime   int64  = 2
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
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	client, err = utils.NewRedisClient(config.RedisConfig.Address, config.RedisConfig.Password, config.RedisConfig.DB)
	if err != nil {
		panic(err)
	}

	var consumerName string = uuid.NewV4().String()

	log.Printf("Initializing Consumer: %v\nConsumer Group: %v\nRedis Stream: %v\n", consumerName, config.RedisConfig.Group, config.RedisConfig.Key)

	createConsumerGroup(config)
	go consumeEvents(config, consumerName)
	go consumePendingEvents(config, consumerName)

	//Gracefully disconnect
	chanOS := make(chan os.Signal, 1)
	signal.Notify(chanOS, syscall.SIGINT, syscall.SIGTERM)
	<-chanOS

	waitGrp.Wait()
	client.Close()
}

func createConsumerGroup(config *config.Config) {
	if _, err := client.XGroupCreateMkStream(config.RedisConfig.Key, config.RedisConfig.Group, "0").Result(); err != nil {
		if !strings.Contains(fmt.Sprint(err), "BUSYGROUP") {
			log.Printf("Error on create Consumer Group: %v ...\n", config.RedisConfig.Group)
			panic(err)
		}
	}
}

func consumeEvents(config *config.Config, consumerName string) {
	for {
		log.Println("New round: ", time.Now().Format(time.RFC3339))
		streams, err := client.XReadGroup(&redis.XReadGroupArgs{
			Streams:  []string{config.RedisConfig.Key, start},
			Group:    config.RedisConfig.Group,
			Consumer: consumerName,
			Count:    config.GeneralConfig.ConsumeEvents,
			Block:    0,
		}).Result()

		if err != nil {
			log.Printf("err on consume events: %+v\n", err)
			return
		}

		for _, stream := range streams[0].Messages {
			waitGrp.Add(1)
			go processStream(config, stream, false, handler.HandlerFactory())
		}
		waitGrp.Wait()
	}
}

func consumePendingEvents(config *config.Config, consumerName string) {
	ticker := time.Tick(time.Second * time.Duration(consumePendingTime))
	for range ticker {
		var streamsRetry []string
		pendingStreams, err := client.XPendingExt(&redis.XPendingExtArgs{
			Stream: config.RedisConfig.Key,
			Group:  config.RedisConfig.Group,
			Start:  "0",
			End:    "+",
			Count:  config.GeneralConfig.ConsumeEvents,
		}).Result()

		if err != nil {
			panic(err)
		}

		for _, stream := range pendingStreams {
			streamsRetry = append(streamsRetry, stream.ID)
		}

		if len(streamsRetry) > 0 {
			streams, err := client.XClaim(&redis.XClaimArgs{
				Stream:   config.RedisConfig.Key,
				Group:    config.RedisConfig.Group,
				Consumer: consumerName,
				Messages: streamsRetry,
				MinIdle:  time.Duration(consumeIdleTime) * time.Second,
			}).Result()

			if err != nil {
				log.Printf("error on process pending: %+v\n", err)
				return
			}

			for _, stream := range streams {
				waitGrp.Add(1)
				go processStream(config, stream, true, handler.HandlerFactory())
			}
			waitGrp.Wait()
		}
		log.Println("process pending streams at: ", time.Now().Format(time.RFC3339))
	}
}

func processStream(config *config.Config, stream redis.XMessage, retry bool, handlerFactory func(t event.Type) handler.Handler) {
	defer waitGrp.Done()

	typeEvent := stream.Values["type"].(string)
	hash := stream.Values["hash"].(string)
	datetime := stream.Values["datetime"].(string)

	timeLayout := time.RFC3339
	parseDate, err := time.Parse(timeLayout, datetime)
	if err != nil {
		log.Printf("RFC format doesn't work: %+v\n", err)
	}

	newEvent, _ := event.New(event.Type(typeEvent))
	newEvent.SetID(stream.ID)

	h := handlerFactory(event.Type(typeEvent))
	err = h.Handle(config, newEvent, hash, parseDate, []byte(stream.Values["data"].(string)), retry)
	if err != nil {
		fmt.Printf("error on process event:%v\n", newEvent)
		fmt.Println(err)
		return
	}

	client.XAck(config.RedisConfig.Key, config.RedisConfig.Group, stream.ID)
	time.Sleep(time.Duration(consumeSleepTime) * time.Second) //to provide an interval for breaking (if necessary) between consumer threads
}
