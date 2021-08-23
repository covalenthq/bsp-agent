package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/covalenthq/mq-store-agent/internal/config"
	"github.com/covalenthq/mq-store-agent/internal/event"
	"github.com/covalenthq/mq-store-agent/internal/handler"
	"github.com/covalenthq/mq-store-agent/internal/utils"
	"github.com/golang/snappy"

	"github.com/go-redis/redis/v7"
	uuid "github.com/satori/go.uuid"
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
	var err error
	client, err = utils.NewRedisClient()
	if err != nil {
		panic(err)
	}
}

func main() {

	// r := gin.Default()

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// r.POST("/cloud-storage-bucket", gcp.HandleFileUploadToBucket)

	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	consumerName = uuid.NewV4().String()
	streamName := config.RedisConfig.Key
	consumerGroup := config.RedisConfig.Group

	fmt.Printf("Initializing Consumer:%v\nConsumerGroup: %v \nStream: %v\n",
		consumerName, consumerGroup, streamName)

	createConsumerGroup(streamName, consumerGroup)
	go consumeEvents(streamName, consumerGroup)
	go consumePendingEvents(streamName, consumerGroup)

	//Gracefully disconection
	chanOS := make(chan os.Signal, 1)
	signal.Notify(chanOS, syscall.SIGINT, syscall.SIGTERM)
	<-chanOS

	waitGrp.Wait()
	client.Close()
}

func createConsumerGroup(streamName, consumerGroup string) {

	if _, err := client.XGroupCreateMkStream(streamName, consumerGroup, "0").Result(); err != nil {

		if !strings.Contains(fmt.Sprint(err), "BUSYGROUP") {
			fmt.Printf("Error on create Consumer Group: %v ...\n", consumerGroup)
			panic(err)
		}

	}
}

// start consume events
func consumeEvents(streamName, consumerGroup string) {

	for {
		func() {
			fmt.Println("New round ", time.Now().Format(time.RFC3339))

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
				//log.Printf("consuming events")
				waitGrp.Add(1)
				go processStream(stream, false, handler.HandlerFactory())
			}
			waitGrp.Wait()
		}()
	}

}

func consumePendingEvents(streamName, consumerGroup string) {

	ticker := time.Tick(time.Second * 30)
	for {
		select {
		case <-ticker:

			func() {

				var streamsRetry []string
				pendingStreams, err := client.XPendingExt(&redis.XPendingExtArgs{
					Stream: streamName,
					Group:  consumerGroup,
					Start:  "0",
					End:    "+",
					Count:  10,
					//Consumer string
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
						log.Printf("err on process pending: %+v\n", err)
						return
					}

					for _, stream := range streams {
						waitGrp.Add(1)
						go processStream(stream, true, handler.HandlerFactory())
					}
					waitGrp.Wait()
				}

				fmt.Println("process pending streams at ", time.Now().Format(time.RFC3339))

			}()

		}

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
		fmt.Println("RFC format doesn't work") // You shouldn't see this at all
	}

	decodeData, err := snappy.Decode(nil, []byte(stream.Values["data"].(string)))
	if err != nil {
		log.Fatal(err)
	}

	newEvent, _ := event.New(event.Type(typeEvent))
	//fmt.Println("event:", typeEvent, "hash:", hash, "date:", parseDate)
	// err = newEvent.UnmarshalBinary(decodeData)
	// if err != nil {
	// 	fmt.Printf("error on unmarshal stream:%v\n", stream.ID)
	// 	return
	// }

	newEvent.SetID(stream.ID)

	h := handlerFactory(newEvent.GetType())
	err = h.Handle(newEvent, hash, parseDate, decodeData, retry)
	if err != nil {
		fmt.Printf("error on process event:%v\n", newEvent)
		fmt.Println(err)
		return
	}

	//client.XDel(streamName, stream.ID)
	client.XAck(streamName, consumerGroup, stream.ID)

	time.Sleep(2 * time.Second)
}

func StringFrom(inter interface{}) string {
	str := ""
	s := reflect.ValueOf(inter)
	fmt.Println(s)
	for i := 0; i < s.Len(); i++ {
		str += string(s.Index(i).Interface().(uint8))
	}
	return str
}
