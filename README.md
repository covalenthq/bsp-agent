# mq-store-agent

Reads messages from a redis mq/stream and persists them in object storage in gcp

## example env

```env
GCP_SERVICE_ACCOUNT=/Users/pranay/.config/gcloud/path/to/service/account.json
GCP_PROJECT_ID=covalent-project
GCP_RESULT_BUCKET=covalenthq-geth-block-result
GCP_SPECIMEN_BUCKET=covalenthq-geth-block-specimen
REDIS_ADDRESS=localhost:6379
REDIS_STREAM_KEY=replication
REDIS_CONSUMER_GROUP=replicate
```
