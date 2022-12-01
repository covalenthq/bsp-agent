// Package storage manages storage of network artifacts
package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"time"

	gcp "cloud.google.com/go/storage"
	"github.com/covalenthq/bsp-agent/internal/config"
	"github.com/covalenthq/bsp-agent/internal/metrics"
	"github.com/covalenthq/bsp-agent/internal/utils"
	"github.com/ipfs/go-cid"
	log "github.com/sirupsen/logrus"
)

const (
	uploadTimeout int64 = 50
)

// Manager composes of all the different storage types supported by the agent
type Manager struct {
	StorageConfig *config.StorageConfig

	GcpStore   *gcp.Client
	LocalStore *LocalStoreClient
	IpfsStore  *ipfsStore

	ipfsSuccessCount metrics.Counter
	ipfsFailureCount metrics.Counter
}

// NewStorageManager creates and sets up a new storage manager with given config
func NewStorageManager(conf *config.StorageConfig) (*Manager, error) {
	manager := &Manager{}
	manager.StorageConfig = conf

	manager.setupGcpStore()
	manager.setupIpfsStore()
	manager.setupLocalFs()
	manager.setupMetrics()

	if manager.IpfsStore == nil {
		log.Errorf("cannot setup IPFS store; no ipfs service flags provided")
	}
	if manager.GcpStore == nil {
		log.Errorf("cannot setup GCP store; no gcp service flags provided")
	}

	return manager, nil
}

// GenerateLocation calculates the non-local location (gcp or ipfs) for the given segment and data, cid is returned in case of ipfs
func (manager *Manager) GenerateLocation(ctx context.Context, segmentName string, replicaSegmentAvro []byte) (string, cid.Cid) {
	config := manager.StorageConfig
	var replicaURL string
	var ccid = cid.Undef
	var err error
	switch {
	case manager.GcpStore != nil:
		replicaURL = "https://storage.cloud.google.com/" + config.ReplicaBucketLoc + "/" + segmentName
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
func (manager *Manager) Store(ctx context.Context, ccid cid.Cid, filename string, data []byte) error {
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

	// ipfs store has priority over gcp
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
	} else if manager.GcpStore != nil {
		err = manager.writeToCloudStorage(ctx, filename, data)
	}

	return err
}

// Close the stores which compose the manager
func (manager *Manager) Close() {
	if manager.GcpStore != nil {
		err := manager.GcpStore.Close()
		if err != nil {
			log.Error("error in closing storage client: ", err)
		}
	}
}

func (manager *Manager) setupGcpStore() {
	// setup gcp storage
	storageConfig := manager.StorageConfig
	gcpStorageClient, err := utils.NewGCPStorageClient(storageConfig.GcpSvcAccountAuthFile)
	if err != nil {
		log.Printf("unable to get gcp storage client; --gcp-svc-account flag not set or set incorrectly: %v", err)

		return
	}

	manager.GcpStore = gcpStorageClient
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

func (manager *Manager) writeToCloudStorage(ctx context.Context, filename string, object []byte) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(uploadTimeout))
	defer cancel()

	bucket := manager.StorageConfig.ReplicaBucketLoc
	wc := manager.GcpStore.Bucket(bucket).Object(filename).NewWriter(ctx)
	if _, err := io.Copy(wc, bytes.NewReader(object)); err != nil {
		return fmt.Errorf("error in copying data to file: %w", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("error in closing file: %w", err)
	}
	log.Info("File successfully uploaded to: https://storage.cloud.google.com/" + bucket + "/" + filename)

	return nil
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
