# mq-store-agent

Reads block-result, block-specimen produced by geth block specimen/result producers streamed into a redis stream, decodes them from RLP encoding and persists them in object storage in gcp

Please refer to this [document](https://www.notion.so/covalenthq/Covalent-Network-Data-Objects-Spec-426cf047602f49a181444e5eeba5c9e4) for more details

## block-result

    ```go
    type BlockResult struct {
        Hash            common.Hash
        TotalDifficulty *big.Int
        Header          *Header
        Transactions    []*Transaction
        Uncles          []*Header
        Receipts        []*Receipt
	    Senders      []common.Address
        }
    ```

## block-specimen

    ```go
        type BlockSpecimen struct {
            AccountRead []*accountRead
            StorageRead []*storageRead
            CodeRead    []*codeRead
        }

        type accountRead struct {
            Address  common.Address
            Nonce    uint64
            Balance  *big.Int
            CodeHash common.Hash
        }

        type storageRead struct {
            Account common.Address
            SlotKey common.Hash
            Value   common.Hash
        }

        type codeRead struct {
            Hash common.Hash
            Code []byte
        }
    ```

## example env

    ```env
    GCP_SERVICE_ACCOUNT=/Users/pranay/.config/gcloud/path/to/service/account.json
    GCP_PROJECT_ID=covalent-project
    GCP_RESULT_BUCKET=covalenthq-geth-block-result
    GCP_SPECIMEN_BUCKET=covalenthq-geth-block-specimen
    REDIS_ADDRESS=localhost:6379
    REDIS_STREAM_KEY=replication
    REDIS_CONSUMER_GROUP=replicate
    CONSUME_EVENTS=10
    ```
