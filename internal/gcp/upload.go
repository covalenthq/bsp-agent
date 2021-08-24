package gcp

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"time"

	"cloud.google.com/go/storage"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"

	"github.com/covalenthq/mq-store-agent/internal/config"
	"github.com/covalenthq/mq-store-agent/internal/event"
)

var (
	storageClient *storage.Client
)

// HandleFileUploadToBucket uploads file to bucket
func HandleResultUploadToBucket(object event.ReplicationEvent, objectName string) error {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Error(err)
	}

	bucketResult := cfg.GcpConfig.ResultBucket

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile(cfg.GcpConfig.ServiceAccount))
	if err != nil {
		return err
	}
	writeToStorage(storageClient, bucketResult, objectName, object)

	return nil
}

func HandleSpecimenUploadToBucket(object event.ReplicationEvent, objectName string) error {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Error(err)
	}

	bucketSpecimen := cfg.GcpConfig.SpecimenBucket

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile(cfg.GcpConfig.ServiceAccount))
	if err != nil {
		return err
	}
	writeToStorage(storageClient, bucketSpecimen, objectName, object)

	return nil
}

func writeToStorage(client *storage.Client, bucket string, objectName string, object event.ReplicationEvent) error {
	// [START upload_file]
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
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
	// [END upload_file]
	return nil
}
