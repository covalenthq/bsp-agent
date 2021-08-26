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

	"github.com/covalenthq/mq-store-agent/internal/config"
	"github.com/covalenthq/mq-store-agent/internal/event"
	"github.com/covalenthq/mq-store-agent/internal/handler"
	"github.com/covalenthq/mq-store-agent/internal/utils"
)

var (
	waitGrp       sync.WaitGroup
	client        *redis.Client
	start         string = ">"
	streamName    string
	consumerGroup string
	consumerName  string
)

func init() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	log.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true
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

	consumerName = uuid.NewV4().String()
	streamName := config.RedisConfig.Key
	consumerGroup := config.RedisConfig.Group

	log.Printf("Initializing Consumer: %v\nConsumer Group: %v\nRedis Stream: %v\n", consumerName, consumerGroup, streamName)

	createConsumerGroup(streamName, consumerGroup)
	go consumeEvents(streamName, consumerGroup)
	go consumePendingEvents(streamName, consumerGroup)

	//Gracefully disconnect
	chanOS := make(chan os.Signal, 1)
	signal.Notify(chanOS, syscall.SIGINT, syscall.SIGTERM)
	<-chanOS

	waitGrp.Wait()
	client.Close()
}

func createConsumerGroup(streamName, consumerGroup string) {
	if _, err := client.XGroupCreateMkStream(streamName, consumerGroup, "0").Result(); err != nil {
		if !strings.Contains(fmt.Sprint(err), "BUSYGROUP") {
			log.Printf("Error on create Consumer Group: %v ...\n", consumerGroup)
			panic(err)
		}
	}
}

func consumeEvents(streamName, consumerGroup string) {
	for {
		log.Println("New round: ", time.Now().Format(time.RFC3339))
		streams, err := client.XReadGroup(&redis.XReadGroupArgs{
			Streams:  []string{streamName, start},
			Group:    consumerGroup,
			Consumer: consumerName,
			Count:    10,
			Block:    0,
		}).Result()

		if err != nil {
			log.Printf("err on consume events: %+v\n", err)
			return
		}

		for _, stream := range streams[0].Messages {
			waitGrp.Add(1)
			go processStream(stream, false, handler.HandlerFactory())
		}
		waitGrp.Wait()
	}
}

func consumePendingEvents(streamName, consumerGroup string) {
	ticker := time.Tick(time.Second * 30)
	for range ticker {
		var streamsRetry []string
		pendingStreams, err := client.XPendingExt(&redis.XPendingExtArgs{
			Stream: streamName,
			Group:  consumerGroup,
			Start:  "0",
			End:    "+",
			Count:  10,
		}).Result()

		if err != nil {
			panic(err)
		}

		for _, stream := range pendingStreams {
			streamsRetry = append(streamsRetry, stream.ID)
		}

		if len(streamsRetry) > 0 {
			streams, err := client.XClaim(&redis.XClaimArgs{
				Stream:   streamName,
				Group:    consumerGroup,
				Consumer: consumerName,
				Messages: streamsRetry,
				MinIdle:  30 * time.Second,
			}).Result()

			if err != nil {
				log.Printf("error on process pending: %+v\n", err)
				return
			}

			for _, stream := range streams {
				waitGrp.Add(1)
				go processStream(stream, true, handler.HandlerFactory())
			}
			waitGrp.Wait()
		}
		log.Println("process pending streams at: ", time.Now().Format(time.RFC3339))
	}
}

func processStream(stream redis.XMessage, retry bool, handlerFactory func(t event.Type) handler.Handler) {
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
	err = h.Handle(newEvent, hash, parseDate, []byte(stream.Values["data"].(string)), retry)
	if err != nil {
		fmt.Printf("error on process event:%v\n", newEvent)
		fmt.Println(err)
		return
	}

	client.XAck(streamName, consumerGroup, stream.ID)
	time.Sleep(2 * time.Second) //break for testing
}
