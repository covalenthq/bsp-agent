package storage

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

// LocalStoreClient stores file on the local filesystem
type LocalStoreClient struct{}

// WriteToBinFile writes to binary file in given path
func (client *LocalStoreClient) WriteToBinFile(path, objectName string, object []byte) error {
	var _, err = os.Stat(filepath.Join(path, filepath.Base(objectName)))
	if os.IsNotExist(err) {
		fullPath := filepath.Join(path, filepath.Base(objectName))
		fileSave, err := os.Create(filepath.Clean(fullPath))
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
