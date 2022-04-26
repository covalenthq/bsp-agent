package node

import (
	"context"
	"sync"

	log "github.com/sirupsen/logrus"
	"gopkg.in/avro.v0"

	"github.com/covalenthq/bsp-agent/internal/config"
	"github.com/covalenthq/bsp-agent/internal/event"
	"github.com/covalenthq/bsp-agent/internal/proof"
	"github.com/covalenthq/bsp-agent/internal/storage"
	"github.com/covalenthq/bsp-agent/internal/utils"
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

	// codec
	ReplicaCodec *goavro.Codec
	EthClient    *ethclient.Client

	redisWaitGrp sync.WaitGroup

	// proochain
	proofchi *proof.ProofchainInteractor

	// redis derived settings
	streamKey     string
	consumerGroup string

	// stream processing
	segment event.ReplicaSegmentWrapped
}

func NewAgentNode(chainType ChainType, aconfig *config.AgentConfig) AgentNode {
	anode := agentNode{}
	anode.AgentConfig = aconfig
	anode.setupRedis()
	anode.setupEthClient()
	anode.setupReplicaCodec()
	anode.setupStorageManager()
	anode.setupProofchainInteractor()

	switch chainType {
	case Ethereum:
		return newEthAgentNode(anode)
	case Elrond:
		return newElrondAgentNode(anode)
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

	node.StorageManager.Close()
	node.EthClient.Close()
}

func (enode *agentNode) setupRedis() {
	redisClient, streamKey, consumerGroup, err := utils.NewRedisClient(&enode.AgentConfig.RedisConfig)
	if err != nil {
		log.Fatalf("unable to get redis client from redis URL flag: %v", err)
	}

	// setup redis client
	enode.RedisClient = redisClient
	enode.streamKey = streamKey
	enode.consumerGroup = consumerGroup
}

func (enode *agentNode) setupEthClient() {
	ethClient, err := utils.NewEthClient(enode.AgentConfig.ChainConfig.RPCURL)
	if err != nil {
		log.Fatalf("unable to get ethereum client from Eth client flag: %v", err)
	}

	enode.EthClient = ethClient
}

func (enode *agentNode) setupReplicaCodec() {
	replicaAvro, err := avro.ParseSchemaFile(enode.AgentConfig.CodecConfig.AvroCodecPath)
	if err != nil {
		log.Fatalf("unable to parse avro schema for specimen from codec path flag: %v", err)
	}
	replicaCodec, err := goavro.NewCodec(replicaAvro.String())
	if err != nil {
		log.Fatalf("unable to generate avro codec for block-replica: %v", err)
	}

	enode.ReplicaCodec = replicaCodec
}

func (enode *agentNode) setupStorageManager() {
	storageManager, err := storage.NewStorageManager(&enode.AgentConfig.StorageConfig)
	if err != nil {
		log.Fatalf("unable to setup storage manager: %v", err)
	}

	enode.StorageManager = storageManager
}

func (anode *agentNode) setupProofchainInteractor() {
	anode.proofchi = proof.NewProofchainInteractor(anode.AgentConfig, anode.EthClient)
}
