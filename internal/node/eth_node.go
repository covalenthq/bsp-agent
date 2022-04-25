package node

import (
	"context"
	"fmt"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"

	"github.com/covalenthq/bsp-agent/internal/config"
	"github.com/covalenthq/bsp-agent/internal/event"
	"github.com/covalenthq/bsp-agent/internal/handler"
	"github.com/covalenthq/bsp-agent/internal/proof"
	st "github.com/covalenthq/bsp-agent/internal/storage"
	"github.com/covalenthq/bsp-agent/internal/utils"
	"github.com/go-redis/redis/v7"
	"github.com/linkedin/goavro/v2"
	"gopkg.in/avro.v0"
)

const (
	consumerEvents            int64  = 1
	consumerPendingIdleTime   int64  = 30
	consumerPendingTimeTicker int64  = 10
	start                     string = ">"
)

type ethAgentNode struct {
	agentNode
}

func newEthAgentNode(agconfig *config.AgentConfig) *ethAgentNode {
	enode := ethAgentNode{}
	enode.AgentConfig = agconfig

	enode.setupRedis()
	enode.setupEthClient()
	enode.setupReplicaCodec()
	enode.setupStorageManager()

	return &enode
}

func (node *ethAgentNode) NodeChainType() ChainType {
	return Ethereum
}

func (node *ethAgentNode) Start(ctx context.Context) {
	var consumerName string = uuid.NewV4().String()
	log.Printf("Initializing Consumer: %v | Redis Stream: %v | Consumer Group: %v", consumerName, node.streamKey, node.consumerGroup)
	createConsumerGroup(node.RedisClient, node.streamKey, node.consumerGroup)
	go node.consumeEvents(consumerName)
}

func (node *ethAgentNode) StopProcessing() {

}

func (node *ethAgentNode) Close() {

}

func (node *ethAgentNode) consumeEvents(consumerName string) {
	for {
		log.Debug("New sequential stream unit: ", time.Now().Format(time.RFC3339))
		streams, err := node.RedisClient.XReadGroup(&redis.XReadGroupArgs{
			Streams:  []string{node.streamKey, start},
			Group:    node.consumerGroup,
			Consumer: consumerName,
			Count:    consumerEvents,
			Block:    0,
		}).Result()
		if err != nil {
			log.Error("error on consume events: ", err.Error())

			return
		}

		for _, stream := range streams[0].Messages {
			node.redisWaitGrp.Add(1)
			go node.processStream(stream)
		}
		node.redisWaitGrp.Wait()
	}
}

func (node *ethAgentNode) processStream(message redis.XMessage) {
	defer node.redisWaitGrp.Done()
	ctx := context.Background()
	replica, objectType, err := ReplicaFromRedisMessage(message)
	objectReplica := replica.Data

	segment := &node.segment

	switch {
	case err != nil:
		log.Error("error on process event: ", err)
	case err == nil && objectReplica.Header.Number.Uint64()%uint64(node.AgentConfig.RedisConfig.BlockDivisor) == 0:
		// collect stream ids and block replicas
		segment.IdBatch = append(segment.IdBatch, message.ID)
		segment.BlockReplicaEvent = append(segment.BlockReplicaEvent, replica)
		segmentLength := node.AgentConfig.SegmentLength()
		if len(segment.BlockReplicaEvent) == 1 {
			segment.StartBlock = replica.Data.Header.Number.Uint64()
		}
		if len(segment.BlockReplicaEvent) == segmentLength {
			segment.EndBlock = replica.Data.Header.Number.Uint64()
			segment.Elements = uint64(segmentLength)
			segment.SegmentName = fmt.Sprint(replica.Data.NetworkId) + "-" + fmt.Sprint(segment.StartBlock) + objectType
			// avro encode, prove and upload

			_, err := node.EncodeProveAndUploadReplicaSegment(ctx, segment)
			if err != nil {
				log.Error("failed to avro encode, prove and upload block-result segment with err: ", err)
				panic(err)
			}
			// ack amd trim stream segment batch id
			xlen, err := utils.AckTrimStreamSegment(node.RedisClient, segmentLength, node.streamKey, node.consumerGroup, segment.IdBatch)
			if err != nil {
				log.Error("failed to match streamIDs length to segment length config: ", err)
			}
			log.Info("stream ids acked and trimmed: ", segment.IdBatch, ", for stream key: ", node.streamKey, ", with current length: ", xlen)
			// reset segment, name, id batch stores
			node.segment = event.ReplicaSegmentWrapped{}
			node.segment.SegmentName = ""
			node.segment.IdBatch = []string{}
		}
	default:
		// collect block replicas and stream ids to skip
		segment.SkipIDBatch = append(segment.SkipIDBatch, message.ID)
		log.Info("block-specimen not created for: ", objectReplica.Header.Number.Uint64(), ", base block number divisor is :", node.AgentConfig.RedisConfig.BlockDivisor)
		if len(segment.BlockReplicaEvent) != 0 {
			// we only proceed with ack'ing/trimming the skipped ids once a segment is flushed.
			return
		}
		// once segment is processed, trim the skipped ids too
		// ack amd trim stream skip batch ids
		xlen, err := utils.AckTrimStreamSegment(node.RedisClient, len(segment.SkipIDBatch), node.streamKey, node.consumerGroup, segment.SkipIDBatch)
		if err != nil {
			log.Error("failed to match streamIDs length to segment length config: ", err)
			panic(err)
		}
		log.Info("stream ids acked and trimmed: ", segment.SkipIDBatch, ", for stream key: ", node.streamKey, ", with current length: ", xlen)
		// reset skip id batch stores
		segment.SkipIDBatch = []string{}
	}
}

func (enode *ethAgentNode) setupRedis() {
	redisClient, streamKey, consumerGroup, err := utils.NewRedisClient(&enode.AgentConfig.RedisConfig)
	if err != nil {
		log.Fatalf("unable to get redis client from redis URL flag: %v", err)
	}

	// setup redis client
	enode.RedisClient = redisClient
	enode.streamKey = streamKey
	enode.consumerGroup = consumerGroup
}

func (enode *ethAgentNode) setupEthClient() {
	ethClient, err := utils.NewEthClient(enode.AgentConfig.ChainConfig.RPCURL)
	if err != nil {
		log.Fatalf("unable to get ethereum client from Eth client flag: %v", err)
	}

	enode.EthClient = ethClient
}

func (enode *ethAgentNode) setupReplicaCodec() {
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

func (enode *ethAgentNode) setupStorageManager() {
	storageManager, err := st.NewStorageManager(&enode.AgentConfig.StorageConfig)
	if err != nil {
		log.Fatalf("unable to setup storage manager: %v", err)
	}

	enode.StorageManager = storageManager
}

func createConsumerGroup(redisClient *redis.Client, streamKey, consumerGroup string) {
	if _, err := redisClient.XGroupCreateMkStream(streamKey, consumerGroup, "0").Result(); err != nil {
		if !strings.Contains(fmt.Sprint(err), "BUSYGROUP") {
			log.Printf("Error on create Consumer Group: %v ...\n", consumerGroup)
			panic(err)
		}
	}
}

// EncodeProveAndUploadReplicaSegment atomically encodes the event into an AVRO binary, proves the replica on proof-chain and upload and stores the binary file
func (node *ethAgentNode) EncodeProveAndUploadReplicaSegment(ctx context.Context, currentSegment *event.ReplicaSegmentWrapped) (string, error) {
	replicaSegmentAvro, err := handler.EncodeReplicaSegmentToAvro(node.ReplicaCodec, currentSegment.ReplicationSegment)
	if err != nil {
		return "", err
	}
	fmt.Printf("\n---> Processing %v <---\n", currentSegment.SegmentName)

	replicaURL, ccid := node.StorageManager.GenerateLocation(ctx, currentSegment, replicaSegmentAvro)
	log.Info("binary file should be available: ", replicaURL)

	log.Info("Submitting block-replica segment proof for: ", currentSegment.SegmentName)
	proofTxHash := make(chan string, 1)
	config := node.AgentConfig
	go proof.SendBlockReplicaProofTx(ctx, config, node.EthClient, currentSegment, replicaSegmentAvro, replicaURL, proofTxHash)
	pTxHash := <-proofTxHash

	if pTxHash == "" {
		return "", fmt.Errorf("failed to prove & upload block-replica segment event: %v", currentSegment.SegmentName)
	}

	log.Info("Proof-chain tx hash: ", pTxHash, " for block-replica segment: ", currentSegment.SegmentName)
	filename := objectFileName(currentSegment.SegmentName, pTxHash)
	err = node.StorageManager.Store(ctx, ccid, filename, replicaSegmentAvro)

	if err != nil {
		return "", fmt.Errorf("error in storing object: %w", err)
	}

	return pTxHash, nil
}

func objectFileName(objectName, txHash string) string {
	return objectName + "-" + txHash
}
