package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"

	"github.com/covalenthq/mq-store-agent/internal/config"
)

var (
	//storageClient *storage.Client
	uploadTimeout int64 = 50
)

func HandleObjectUploadToBucket(config *config.Config, objectType string, objectName string, object interface{}) error {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(uploadTimeout))
	defer cancel()

	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile(config.GcpConfig.ServiceAccount))
	if err != nil {
		return err
	}

	switch objectType {
	case "block-specimen":
		bucket := config.GcpConfig.SpecimenBucket
		return writeToStorage(storageClient, bucket, objectName, object)
	case "block-result":
		bucket := config.GcpConfig.ResultBucket
		return writeToStorage(storageClient, bucket, objectName, object)
	}

	return nil
}

func writeToStorage(client *storage.Client, bucket string, objectName string, object interface{}) error {
	ctx := context.Background()
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
