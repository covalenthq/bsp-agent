package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/storage"

	log "github.com/sirupsen/logrus"
)

var (
	uploadTimeout int64 = 50
)

func HandleObjectUploadToBucket(ctx context.Context, storageClient *storage.Client, objectType, objectStorage, objectName string, object interface{}) error {

	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(uploadTimeout))
	defer cancel()

	switch objectType {
	case "block-specimen":
		bucket := objectStorage
		return writeToStorage(ctx, storageClient, bucket, objectName, object)
	case "block-result":
		bucket := objectStorage
		return writeToStorage(ctx, storageClient, bucket, objectName, object)
	}

	return fmt.Errorf("type: %v not supported", objectType)
}

func writeToStorage(ctx context.Context, client *storage.Client, bucket string, objectName string, object interface{}) error {

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
