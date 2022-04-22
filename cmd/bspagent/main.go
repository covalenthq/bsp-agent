package main

import (
	"context"
	"flag"
	"io"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	"cloud.google.com/go/storage"
	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/covalenthq/lumberjack/v3"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v7"
	"github.com/linkedin/goavro/v2"
	log "github.com/sirupsen/logrus"

	"github.com/covalenthq/bsp-agent/internal/config"
	"github.com/covalenthq/bsp-agent/internal/node"
	"github.com/covalenthq/bsp-agent/internal/utils"
	pinapi "github.com/covalenthq/ipfs-pinner"
)

const (
	consumerEvents            int64  = 1
	consumerPendingIdleTime   int64  = 30
	consumerPendingTimeTicker int64  = 10
	start                     string = ">"
)

var (
	agconfig  config.AgentConfig
	agentNode node.AgentNode
)

func init() {
	agconfig = config.NewAgentConfig()
	agconfig.LoadConfig()

	// setup logger
	formatter := runtime.Formatter{ChildFormatter: &log.TextFormatter{
		FullTimestamp: true,
	}}
	formatter.Line = true
	log.SetFormatter(&formatter)

	var outWriter io.Writer
	logLocationURL, err := utils.GetLogLocationURL(agconfig.LogFolder)
	if err != nil {
		log.Warn("error while setting up file logging: ", err)
		outWriter = os.Stdout
	} else {
		logFilePath := path.Join(logLocationURL.Path, "log.log")
		bspLogger := utils.NewLoggerOut(os.Stdout, &lumberjack.Logger{
			// logs folder created/searched in directory in which agent was started.
			Filename:        logFilePath,
			MaxSize:         100,          // for a log file (in megabytes)
			MaxBackups:      10,           // maximum number of backup files to be created
			MaxAge:          60,           // days
			RollingInterval: 24 * 60 * 60, // 1 day (in seconds) -- one log file for each day
			Compress:        true,         // gzip rolled over files
		})
		outWriter = &bspLogger
	}

	log.SetOutput(outWriter)
	log.SetLevel(log.InfoLevel)
	log.WithFields(log.Fields{"file": "main.go"}).Info("bsp-agent is running...")
}

func main() {
	log.Info("bsp-agent command line config: ", utils.GetConfig(flag.CommandLine))
	chainType := determineChainType(&agconfig)
	agentNode = node.InitAgentNode(chainType, &agconfig)

	// webSockUrls := agconfig.ChainConfig.WebsocketURLs
	// if webSockUrls != "" {
	// 	// elrond chain
	// 	agentNode = node.NewAgentNode(node.Elrond, &agconfig)
	// 	// websocketsURLs := strings.Split(webSockUrls, " ")
	// 	// for _, url := range websocketsURLs {
	// 	// 	go websocket.ConsumeWebsocketsEvents(&agconfig, url, replicaCodec, ethClient, gcpStorageClient)
	// 	// }
	// } else {
	// 	// ethereum chain
	// 	agentNode = node.NewAgentNode(node.Ethereum, &agconfig)
	// 	// var consumerName string = uuid.NewV4().String()
	// 	// log.Printf("Initializing Consumer: %v | Redis Stream: %v | Consumer Group: %v", consumerName, streamKey, consumerGroup)
	// 	// createConsumerGroup(redisClient, streamKey, consumerGroup)
	// 	// go consumeEvents(config, replicaCodec, redisClient, pinnode, gcpStorageClient, ethClient, consumerName, streamKey, consumerGroup)
	// 	// go consumePendingEvents(config, replicaCodec, redisClient, pinnode, gcpStorageClient, ethClient, consumerName, streamKey, consumerGroup)
	// }
	agentNode.Start(context.TODO())

	// Gracefully disconnect
	chanOS := make(chan os.Signal, 1)
	signal.Notify(chanOS, syscall.SIGINT, syscall.SIGTERM)
	<-chanOS

	agentNode.StopProcessing()
	agentNode.Close()
}

func determineChainType(agconfig *config.AgentConfig) node.ChainType {
	webSockUrls := agconfig.ChainConfig.WebsocketURLs
	if webSockUrls != "" {
		return node.Elrond
	} else {
		return node.Ethereum
	}
}

// consume pending messages from redis
func consumePendingEvents(config *config.Config, avroCodecs *goavro.Codec, redisClient *redis.Client, pinnode pinapi.PinnerNode, gcpStorageClient *storage.Client, ethClient *ethclient.Client, consumerName, streamKey, consumerGroup string) {
	timeout := time.After(time.Second * time.Duration(consumerPendingTimeoutFlag))
	ticker := time.Tick(time.Second * time.Duration(consumerPendingTimeTicker))
	for {
		select {
		case <-timeout:
			// this would always timeout and exit. What's the point of this then?
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
					go processStream(config, avroCodecs, redisClient, gcpStorageClient, pinnode, ethClient, stream, streamKey, consumerGroup)
				}
				waitGrp.Wait()
			}
			log.Info("Process pending streams at: ", time.Now().Format(time.RFC3339))
		}
	}
}
