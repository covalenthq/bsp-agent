package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/storage"

	"github.com/covalenthq/mq-store-agent/internal/config"
)

var (
	uploadTimeout int64 = 50
)

func HandleObjectUploadToBucket(ctx context.Context, config *config.GcpConfig, storageClient *storage.Client, objectType string, objectName string, object interface{}) error {

	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(uploadTimeout))
	defer cancel()

	switch objectType {
	case "block-specimen":
		bucket := config.SpecimenBucket
		return writeToStorage(ctx, storageClient, bucket, objectName, object)
	case "block-result":
		bucket := config.ResultBucket
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

	return nil
}
