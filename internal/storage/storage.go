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

	log "github.com/sirupsen/logrus"
)

var (
	uploadTimeout int64 = 50
)

// HandleObjectUploadToBucket handles AVRO binary object upload to cloud bucket
func HandleObjectUploadToBucket(ctx context.Context, storageClient *storage.Client, binaryLocalPath, storageBucket, objectName, txHash string, object []byte) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(uploadTimeout))
	defer cancel()

	switch {
	case binaryLocalPath == "":
		return writeToCloudStorage(ctx, storageClient, storageBucket, objectName+"-"+txHash, object)
	case storageClient == nil:
		err := writeToBinFile(binaryLocalPath, objectName+"-"+txHash, object)
		if err != nil {
			panic(err)
		}

		return err
	default:
		err := validatePath(binaryLocalPath, objectName+"-"+txHash)
		if err != nil {
			panic(err)
		} else {
			err = writeToBinFile(binaryLocalPath, objectName+"-"+txHash, object)
			if err != nil {
				panic(err)
			}
		}

		return writeToCloudStorage(ctx, storageClient, storageBucket, objectName+"-"+txHash, object)
	}
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
