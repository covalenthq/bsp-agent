package storage

import (
	"bytes"
	"context"
	"encoding/json"
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

func HandleObjectUploadToBucket(ctx context.Context, storageClient *storage.Client, path, storageBucket, objectName string, object interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(uploadTimeout))
	defer cancel()

	if path == "" {
		return writeToCloudStorage(ctx, storageClient, storageBucket, objectName, object)
	} else {
		err := validateDirPath(path, objectName)
		if err != nil {
			panic(err)
		} else {
			err = writeToBinFile(ctx, path, objectName, object)
			if err != nil {
				panic(err)
			}
		}
		return nil
		//return writeToCloudStorage(ctx, storageClient, storageBucket, objectName, object)
	}
}

func writeToCloudStorage(ctx context.Context, client *storage.Client, bucket, objectName string, object interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(uploadTimeout))
	defer cancel()

	wc := client.Bucket(bucket).Object(objectName).NewWriter(ctx)
	content, err := json.Marshal(object)
	if err != nil {
		return err
	}
	if _, err := io.Copy(wc, bytes.NewReader(content)); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}
	log.Info("Object uploaded to: https://storage.cloud.google.com/" + bucket + "/" + objectName)

	return nil
}

//todo: reading from the AVRO binary file aka cat (need AVRO tools)
func writeToBinFile(ctx context.Context, path, objectName string, object interface{}) error {
	// ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(uploadTimeout))
	// defer cancel()
	var _, err = os.Stat(filepath.Join(path, filepath.Base(objectName)))
	if os.IsNotExist(err) {
		fileSave, err := os.Create(filepath.Join(path, filepath.Base(objectName)))
		if err != nil {
			return err
		}
		defer fileSave.Close()
		_, err = fileSave.Write([]byte(fmt.Sprintf("%v", object)))
		if err != nil {
			panic(err)
		}
	} else {
		log.Info("File already exists at: ", path, "/", objectName)
	}
	log.Info("File written successfully to: ", path, "/", objectName)

	return nil
}

func validateDirPath(path, objectName string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("path does not exist: %v", err)
	}
	mode := fileInfo.Mode()
	if mode.IsDir() {
		log.Info("Writing block-replica binary file to local directory: ", path, "/", objectName)
	}

	return nil
}
