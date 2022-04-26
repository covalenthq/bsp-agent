package main

import (
	"context"
	"flag"
	"io"
	"os"
	"os/signal"
	"path"
	"syscall"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/covalenthq/lumberjack/v3"
	log "github.com/sirupsen/logrus"

	"github.com/covalenthq/bsp-agent/internal/config"
	"github.com/covalenthq/bsp-agent/internal/node"
	"github.com/covalenthq/bsp-agent/internal/utils"
)

var (
	agconfig  *config.AgentConfig
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
	chainType := determineChainType(agconfig)
	agentNode = node.NewAgentNode(chainType, agconfig)

	agentNode.Start(context.Background())

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
