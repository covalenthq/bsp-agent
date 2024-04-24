// Package storage manages storage of network artifacts
package storage

import (
	"fmt"
	"os"

	"github.com/covalenthq/bsp-agent/internal/config"
	"github.com/covalenthq/bsp-agent/internal/metrics"
	"github.com/ipfs/go-cid"
	log "github.com/sirupsen/logrus"
)

// Manager composes of all the different storage types supported by the agent
type Manager struct {
	StorageConfig *config.StorageConfig

	LocalStore *LocalStoreClient
	IpfsStore  *ipfsStore

	ipfsSuccessCount metrics.Counter
	ipfsFailureCount metrics.Counter
}

// NewStorageManager creates and sets up a new storage manager with given config
func NewStorageManager(conf *config.StorageConfig) (*Manager, error) {
	manager := &Manager{}
	manager.StorageConfig = conf

	manager.setupIpfsStore()
	manager.setupLocalFs()
	manager.setupMetrics()

	if manager.IpfsStore == nil {
		log.Errorf("cannot setup IPFS store; no ipfs service flags provided")
	}

	return manager, nil
}

// GenerateLocation calculates the non-local location (ipfs) for the given segment and data, cid is returned in case of ipfs
func (manager *Manager) GenerateLocation(segmentName string, replicaSegmentAvro []byte) (string, cid.Cid) {
	config := manager.StorageConfig
	var replicaURL string
	var ccid = cid.Undef
	var err error
	switch {
	case manager.IpfsStore != nil:
		ccid, err = manager.IpfsStore.CalcCid(replicaSegmentAvro)
		if err != nil {
			log.Errorf("error generating cid for %s. Error: %s", config.BinaryFilePath+segmentName, err)
			replicaURL = "only local: " + config.BinaryFilePath + segmentName
		} else {
			replicaURL = "ipfs://" + ccid.String()
		}
	default:
		replicaURL = "only local: " + config.BinaryFilePath + segmentName
	}

	return replicaURL, ccid
}

// Store the given data in the stores supported by the agent.
// cid is needed for IFPS based stores, and cid.Undef can be passed in case IPFS is not needed.
func (manager *Manager) Store(ccid cid.Cid, filename string, data []byte) error {
	// write to local store
	var err error

	if manager.StorageConfig.BinaryFilePath != "" {
		err = validatePath(manager.StorageConfig.BinaryFilePath, filename)
		if err != nil {
			return err
		}
		err = manager.LocalStore.WriteToBinFile(manager.StorageConfig.BinaryFilePath, filename, data)
		if err != nil {
			return err
		}
	}

	if manager.IpfsStore != nil {
		if ccid == cid.Undef {
			return fmt.Errorf("cid is Undefined")
		}
		var ucid cid.Cid
		ucid, err = manager.IpfsStore.Upload(data)
		if err != nil {
			manager.ipfsFailureCount.Inc(1)
			log.Errorf("ipfs store reported error: %v", err)

			return err
		}
		manager.ipfsSuccessCount.Inc(1)
		log.Infof("client side cid is: %s, while uploaded is: %s", ccid.String(), ucid.String())
	}

	return err
}

// Close the stores which compose the manager
func (manager *Manager) Close() {
	// no op
}

func (manager *Manager) setupIpfsStore() {
	store, err := newIpfsStore(manager.StorageConfig)
	if err != nil {
		log.Fatalf("error creating ipfs store: %v", err)
	}

	manager.IpfsStore = store
}

func (manager *Manager) setupLocalFs() {
	if manager.StorageConfig.BinaryFilePath == "" {
		log.Warn("--binary-file-path flag not provided to write block-replica avro encoded binary files to local path")
	}
}

func (manager *Manager) setupMetrics() {
	manager.ipfsSuccessCount = metrics.GetOrRegisterCounter("agent/storage/ipfs/success", metrics.DefaultRegistry)
	manager.ipfsFailureCount = metrics.GetOrRegisterCounter("agent/storage/ipfs/failure", metrics.DefaultRegistry)
}

func validatePath(path, objectName string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("path does not exist: %w", err)
	}
	mode := fileInfo.Mode()
	if mode.IsDir() {
		log.Info("Writing block-replica binary file to local directory: ", path, objectName)
	}

	return nil
}
