// Package storage contains all function related to local and cloud storage
package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"cloud.google.com/go/storage"

	pinapi "github.com/covalenthq/ipfs-pinner"
	pincore "github.com/covalenthq/ipfs-pinner/core"
	"github.com/covalenthq/ipfs-pinner/pinclient"
	"github.com/ipfs/go-cid"
	log "github.com/sirupsen/logrus"
)

const (
	uploadTimeout int64 = 50
)

// HandleObjectUploadToBucket handles AVRO binary object upload to cloud bucket
func HandleObjectUploadToBucket(ctx context.Context, gcpStorageClient *storage.Client, binaryLocalPath, storageBucket, objectName, txHash string, object []byte) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(uploadTimeout))
	defer cancel()

	filename := objectFileName(objectName, txHash)

	switch {
	case binaryLocalPath == "":
		return writeToCloudStorage(ctx, gcpStorageClient, storageBucket, filename, object)
	case gcpStorageClient == nil:
		err := writeToBinFile(binaryLocalPath, filename, object)
		if err != nil {
			panic(err)
		}

		return err
	default:
		err := validatePath(binaryLocalPath, filename)
		if err != nil {
			panic(err)
		} else {
			err = writeToBinFile(binaryLocalPath, filename, object)
			if err != nil {
				panic(err)
			}
		}

		return writeToCloudStorage(ctx, gcpStorageClient, storageBucket, filename, object)
	}
}

// GenerateCidFor generates ipfs cid given some content
func GenerateCidFor(ctx context.Context, pinnode pinapi.PinnerNode, content []byte) (cid.Cid, error) {
	if pinnode == nil {
		return cid.Undef, fmt.Errorf("no pinner node")
	}

	rcid, err := pinnode.UnixfsService().GenerateDag(ctx, bytes.NewReader(content))
	if err != nil {
		return cid.Undef, fmt.Errorf("error generating dag: %w", err)
	}

	return rcid, nil
}

// HandleObjectUploadToIPFS uploads the binary file to ipfs via the pinner client
func HandleObjectUploadToIPFS(ctx context.Context, pinnode pinapi.PinnerNode, ccid cid.Cid, binaryLocalPath, segmentName, pTxHash string) cid.Cid {
	// assuming that bin files are written (rather than cloud only storage)
	// need to explore strategy to directly upload in memory byte array via pinner
	if pinnode == nil || ccid == cid.Undef {
		return cid.Undef
	}

	var file *os.File
	var err error
	if pinnode.PinService().ServiceType() == pincore.Web3Storage {
		file, err = generateCarFile(ctx, pinnode, ccid)
	} else {
		objPath := objectFilePath(segmentName, pTxHash, binaryLocalPath)
		file, err = os.Open(filepath.Clean(objPath))
	}

	if err != nil {
		log.Error("failure in opening/generating file for upload: ", err)

		return cid.Undef
	}

	fcid, err := pinnode.PinService().UploadFile(ctx, file)
	if err != nil {
		log.Error("failure in uploading specimen object to IPFS: ", err)

		return cid.Undef
	}

	log.Infof("File %s successfully uploaded to IPFS with pin: %s", file.Name(), fcid.String())

	return fcid
}

// GetPinnerNode get pinner node (web3.storage or pinata supported for now)
func GetPinnerNode(pst pincore.PinningService, token string) (pinapi.PinnerNode, error) {
	var pinnode pinapi.PinnerNode
	switch pst {
	case pincore.Pinata, pincore.Web3Storage:
		clientCreateReq := pinclient.NewClientRequest(pst).BearerToken(token)
		cidComputationOnly := (pst == pincore.Pinata)
		nodeCreateReq := pinapi.NewNodeRequest(clientCreateReq).CidVersion(0).CidComputeOnly(cidComputationOnly)
		pinnode = pinapi.NewPinnerNode(*nodeCreateReq)

		return pinnode, nil
	case "":
		return nil, nil
	case pincore.Other:
		fallthrough
	default:
		return nil, fmt.Errorf("unsupported pinning service: %s", pst)
	}
}

func generateCarFile(ctx context.Context, pinnode pinapi.PinnerNode, ccid cid.Cid) (*os.File, error) {
	carf, err := os.CreateTemp(os.TempDir(), "*.car")
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	log.Printf("car file location: %s\n", carf.Name())

	err = pinnode.CarExporter().Export(ctx, ccid, carf)
	if err != nil {
		_ = carf.Close()

		return nil, fmt.Errorf("%w", err)
	}

	noffset, err := carf.Seek(0, 0) // reset for Read
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	} else if noffset != 0 {
		return nil, fmt.Errorf("couldn't offset the car file to start")
	}

	return carf, nil
}

func writeToCloudStorage(ctx context.Context, client *storage.Client, bucket, objectName string, object []byte) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(uploadTimeout))
	defer cancel()

	wc := client.Bucket(bucket).Object(objectName).NewWriter(ctx)
	if _, err := io.Copy(wc, bytes.NewReader(object)); err != nil {
		return fmt.Errorf("error in copying data to file: %w", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("error in closing file: %w", err)
	}
	log.Info("File successfully uploaded to: https://storage.cloud.google.com/" + bucket + "/" + objectName)

	return nil
}

//nolint:gosec
func writeToBinFile(path, objectName string, object []byte) error {
	var _, err = os.Stat(filepath.Join(path, filepath.Base(objectName)))
	if os.IsNotExist(err) {
		fileSave, err := os.Create(filepath.Join(path, filepath.Base(objectName)))
		if err != nil {
			return fmt.Errorf("error in writing binary file: %w", err)
		}
		defer func() {
			if err := fileSave.Close(); err != nil {
				log.Error("Error closing file: ", err)
			}
		}()
		_, err = fileSave.Write(object)
		if err != nil {
			panic(err)
		}
	} else {
		log.Info("File already exists at: ", path, objectName)
	}
	log.Info("File written successfully to: ", path, objectName)

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

func objectFileName(objectName, txHash string) string {
	return objectName + "-" + txHash
}

func objectFilePath(objectName, txHash, binaryLocalPath string) string {
	filename := objectFileName(objectName, txHash)

	return filepath.Join(binaryLocalPath, filepath.Base(filename))
}
