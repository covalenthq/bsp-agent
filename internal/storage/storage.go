package storage

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

func HandleResultUploadToBucket(object event.ReplicationEvent, objectName string) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Error(err)
	}

	resultBucket := cfg.GcpConfig.ResultBucket

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile(cfg.GcpConfig.ServiceAccount))
	if err != nil {
		return err
	}
	writeToStorage(storageClient, resultBucket, objectName, object)

	return nil
}

func HandleSpecimenUploadToBucket(object event.ReplicationEvent, objectName string) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Error(err)
	}

	specimenBucket := cfg.GcpConfig.SpecimenBucket

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile(cfg.GcpConfig.ServiceAccount))
	if err != nil {
		return err
	}
	writeToStorage(storageClient, specimenBucket, objectName, object)

	return nil
}
func writeToStorage(client *storage.Client, bucket string, objectName string, object event.ReplicationEvent) error {
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

	return nil
}
