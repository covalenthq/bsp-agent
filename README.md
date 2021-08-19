# mq-store-agent

reads messages from a redis mq/stream and persists them in object storage in gcp

## example env

```env
GCP_SERVICE_ACCOUNT=/Users/name/.config/gcloud/path/to/creds.json
GCP_PROJECT_ID=name-of-project
GCP_BUCKET_NAME=name-of-gcp-storage-bucket
```
