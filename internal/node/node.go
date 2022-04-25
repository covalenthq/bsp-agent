package node

import (
	"context"
	"sync"

	log "github.com/sirupsen/logrus"

	gcp "cloud.google.com/go/storage"
	"github.com/covalenthq/bsp-agent/internal/config"
	"github.com/covalenthq/bsp-agent/internal/event"
	"github.com/covalenthq/bsp-agent/internal/storage"
	pinner "github.com/covalenthq/ipfs-pinner"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v7"
	"github.com/linkedin/goavro/v2"
)

type ChainType string

const (
	Ethereum ChainType = "ethereum"
	Elrond   ChainType = "elrond"
)

type agentNode struct {
	AgentConfig *config.AgentConfig

	RedisClient *redis.Client

	// storage
	StorageManager *storage.StorageManager
	GcpStore       *gcp.Client
	LocalStore     *storage.LocalStoreClient
	IpfsStore      *pinner.PinnerNode

	// codec
	ReplicaCodec *goavro.Codec
	EthClient    *ethclient.Client

	redisWaitGrp sync.WaitGroup

	// redis derived settings
	streamKey     string
	consumerGroup string

	// stream processing
	// TODO: need to evaluate if all or some of these can be local
	segment event.ReplicaSegmentWrapped
}

func InitAgentNode(chainType ChainType, aconfig *config.AgentConfig) AgentNode {
	switch chainType {
	case Ethereum:
		return newEthAgentNode(aconfig)
	case Elrond:
		return newElrondAgentNode(aconfig)
	default:
		log.Fatalf("unknown chainType requested: %s", chainType)
	}

	// unreachable
	return nil
}

func (node *agentNode) NodeChainType() ChainType {
	log.Fatal("NodeChainType() shouldn't be called directly for AgentNode")
	// unreachable
	return ""
}

func (node *agentNode) Start(ctx context.Context) {
	log.Fatal("Start() shouldn't be called directly for AgentNode")
}

func (node *agentNode) StopProcessing() {
	log.Warn("Received interrupt. Flushing in-memory blocks...")
	node.redisWaitGrp.Wait()
	log.Warn("waitgrp ended. Closing node...")
}

func (node *agentNode) Close() {
	if node.RedisClient != nil {
		err := node.RedisClient.Close()
		if err != nil {
			log.Error("error in closing redis client: ", err)
		}
	}

	if node.GcpStore != nil {
		err := node.GcpStore.Close()
		if err != nil {
			log.Error("error in closing storage client: ", err)
		}
	}
	node.EthClient.Close()
}
