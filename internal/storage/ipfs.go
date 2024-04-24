package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/covalenthq/bsp-agent/internal/config"
	"github.com/ipfs/go-cid"
	log "github.com/sirupsen/logrus"
	"github.com/ybbus/httpretry"
)

type ipfsStore struct {
	baseURL    string // url of ipfs-pinner node
	httpClient *http.Client
}

func newIpfsStore(conf *config.StorageConfig) (*ipfsStore, error) {
	store := ipfsStore{}
	log.Infof("ipfs store pointint at: %s", conf.IpfsPinnerServer)
	if _, err := url.Parse(conf.IpfsPinnerServer); err != nil {
		return nil, fmt.Errorf("url parse error: %w", err)
	}
	store.baseURL = conf.IpfsPinnerServer
	store.httpClient = newHTTPClient(nil)

	return &store, nil
}

// Upload uploads given content to ipfs
func (store *ipfsStore) Upload(contents []byte) (cid.Cid, error) {
	return store.fetchCid(contents, true)
}

// CalcCid finds the cid of given content
func (store *ipfsStore) CalcCid(contents []byte) (cid.Cid, error) {
	return store.fetchCid(contents, false)
}

func (store *ipfsStore) fetchCid(contents []byte, withUpload bool) (cid.Cid, error) {
	targetURL, _ := url.Parse(store.baseURL)
	ctx := context.Background()
	if withUpload {
		targetURL.Path = path.Join(targetURL.Path, "upload")
	} else {
		targetURL.Path = path.Join(targetURL.Path, "cid")
	}

	buffer, contentType, err := store.createMultiformWriter(contents)
	if err != nil {
		return cid.Undef, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, targetURL.String(), buffer)
	if err != nil {
		return cid.Undef, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", contentType)
	resp, err := store.httpClient.Do(req)
	if err != nil {
		return cid.Undef, fmt.Errorf("error in fetching data: %w", err)
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Error("error closing response body", err)
		}
	}()

	rbody, err := io.ReadAll(resp.Body)
	if err != nil {
		return cid.Undef, fmt.Errorf("error in reading response body: %w", err)
	}

	responseMap := make(map[string]string)
	err = json.Unmarshal(rbody, &responseMap)

	if err != nil {
		return cid.Undef, fmt.Errorf("unmarshalling error: %w", err)
	}

	if msg, ok := responseMap["error"]; ok {
		return cid.Undef, fmt.Errorf("server error: %s", msg)
	}

	if rcid, ok := responseMap["cid"]; ok {
		//nolint:wrapcheck
		return cid.Parse(rcid)
	}

	return cid.Undef, fmt.Errorf("unexpected response: %s", string(rbody))
}

func (store *ipfsStore) createMultiformWriter(contents []byte) (*bytes.Buffer, string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if err := writer.WriteField("filedata", string(contents)); err != nil {
		return body, "", fmt.Errorf("error writing form field: %w", err)
	}

	if err := writer.Close(); err != nil {
		log.Error("error closing multipart writer", err)
	}

	return body, writer.FormDataContentType(), nil
}

func newHTTPClient(client *http.Client) *http.Client {
	if client == nil {
		client = &http.Client{}
	}

	return httpretry.NewCustomClient(
		client,
		// retry x times
		httpretry.WithMaxRetryCount(3),
		// retry on status == 429, if status >= 500, if err != nil, or if response was nil (status == 0)
		httpretry.WithRetryPolicy(func(statusCode int, err error) bool {
			return err != nil || statusCode == http.StatusTooManyRequests || statusCode >= http.StatusInternalServerError || statusCode == 0
		}),
		httpretry.WithBackoffPolicy(httpretry.ExponentialBackoff(3*time.Second, time.Minute, 2*time.Second)),
	)
}
