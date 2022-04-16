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

	pinner "github.com/covalenthq/ipfs-pinner"
	"github.com/ipfs/go-cid"
	log "github.com/sirupsen/logrus"
)

const (
	uploadTimeout int64 = 50
)

// HandleObjectUploadToBucket handles AVRO binary object upload to cloud bucket
func HandleObjectUploadToBucket(ctx context.Context, storageClient *storage.Client, binaryLocalPath, storageBucket, objectName, txHash string, object []byte) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(uploadTimeout))
	defer cancel()

	filename := objectFileName(objectName, txHash)

	switch {
	case binaryLocalPath == "":
		return writeToCloudStorage(ctx, storageClient, storageBucket, filename, object)
	case storageClient == nil:
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

		return writeToCloudStorage(ctx, storageClient, storageBucket, filename, object)
	}
}

// HandleObjectUploadToIPFS uploads the binary file to ipfs via the pinner client
func HandleObjectUploadToIPFS(ctx context.Context, client *pinner.Client, binaryLocalPath string, objectName string, txHash string) cid.Cid {
	// assuming that bin files are written (rather than cloud only storage)
	// need to explore strategy to directly upload in memory byte array via pinner
	if client == nil {
		return cid.Undef
	}
	filename := objectFileName(objectName, txHash)
	objectpath := filepath.Join(binaryLocalPath, filename)
	if _, err := os.Stat(objectpath); os.IsNotExist(err) {
		log.Infof("%s doesn't exist in local. Cannot upload to IFPS", objectpath)

		return cid.Undef
	}

	objf, err := os.Open(filepath.Clean(objectpath))
	if err != nil {
		log.Error("error opening specimen object file for upload: ", err)

		return cid.Undef
	}
	cid, err := client.UploadFile(ctx, objf)
	if err != nil {
		log.Error("failure in uploading specimen object to IPFS: ", err)
	}

	log.Infof("File %s successfully uploaded to IPFS with pin: %s", filename, cid.String())

	return cid
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
