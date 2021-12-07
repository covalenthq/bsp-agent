# BSP-agent (mq-store-agent)

Reads block-replicas (block-results, block-specimens, or any other types) produced by geth block specimen/result producers streamed into a redis stream, decodes them from RLP encoding, creates a proof transaction on a virtual chain with a sha-256 checksum of the data in the object, and persists them into object storage for a specified google bucket and saves it to a file - atomically.

Please refer to these internal [instructions](https://docs.google.com/document/d/1BMC9-VXZfpB6mGczSu8ylUXJZ_CIx4ephepDtlruv_Q/edit?usp=sharing) for running the BSP with the mq-store-agent (BSP-agent).

Please refer to this internal [whitepaper](https://docs.google.com/document/d/1J6RalVVfMSh2kSKNHM3Agb4GngzWVw9e1PqLSVb3-PU/edit#) to understand more about its function.

Externally facing validator [documentation.](https://www.notion.so/covalenthq/Validator-Documentation-e9fdba94c9e149aeba798ece303dc5d4)

Externally facing workshop [deck.](https://docs.google.com/presentation/d/1qInReJcMxvVywJ8onoFPoKCwuorJ8LpOn3hwLJIl7bg/edit?usp=sharing)

## Architecture

![diagram](arch.png)

## Block-replica

Block Replicas are created by [BSP](https://docs.google.com/document/d/1BMC9-VXZfpB6mGczSu8ylUXJZ_CIx4ephepDtlruv_Q/edit#heading=h.5owqpz3w99gp) here and fed into [Redis streams](https://redis.io/topics/streams-intro).

These objects are extracted and read into the following struct by the agent. There are three types of objects currently -

1. block-replica - have all the data shown below.
1. block-result - have all the data shown below except "State".
1. block-specimen - have all the data shown below except "TotalDifficulty", "Receipts"(&"Logs"), "Senders"

```go
    type BlockReplica struct {
        Type            string
        NetworkId       uint64
        Hash            common.Hash
        TotalDifficulty *big.Int
        Header          *Header
        Transactions    []*Transaction
        Uncles          []*Header
        Receipts        []*Receipt
        Senders         []common.Address
        State           *StateSpecimen
    }
```

### State-specimen

The "State" is comprised of all state information related to accounts ever touched for a given block.

```go
    type StateSpecimen struct {
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

## Environment

The Eth private key allows the validator to make proof transactions to the proof-chain contract and is required. Other env vars are optional depending on your redis, eth account configuration. Add the following to your `.envrc` at the root dir with final relative path `~/mq-store-agent/envrc`

```bash
    export ETH_PRIVATE_KEY=private/key/senders #required
    export REDIS_PWD=your-redis-password #optional
    export ETH_KEYSTORE_PATH=path/to/keystore/file.json #optional
    export ETH_KEYSTORE_PWD=password/to/access/keystore/file.json #optional
```

Please `brew install direnv` add the following to you bash -

```bash
    eval "$(direnv hook bash)" # bash users - add the following line to your ~/.bashrc
    eval "$(direnv hook zsh)" # zsh users - add the following line to your ~/.zshrc
```

And enable the vars with `direnv allow .`
For which you should see something like -

```bash
    direnv: loading ~/Documents/covalent/mq-store-agent/.envrc
    direnv: export +ETH_PRIVATE_KEY
```

## Build & Run

Run the agent directly with the following -

```bash
go run ./cmd/mqstoreagent/*.go \ 
    --redis-url="redis://username:@localhost:6379/0?topic=replication#replicate" \ 
    --codec-path="./codec/block-replica.avsc" \ 
    --binary-file-path="./bin/block-replica/" \ 
    --gcp-svc-account="/Users/<user>/.config/gcloud/<gcp-service-account.json>" \ 
    --replica-bucket="<covalenthq-geth-block-replica-bucket>" \ 
    --segment-length=5 \ 
    --eth-client="http://127.0.0.1:7545" \ 
    --proof-chain-address="0xb5B12cbe8bABAF96677F60f65317b81709062C47"
```

Or update the Makefile with the correct --gcp-svc-account, --replica-bucket & --proof-chain-address and run with the following.

```bash
    make run-build
    make run-agent
```

### Docker

The docker image for this service can be found [here](https://github.com/covalenthq/mq-store-agent/pkgs/container/mq-store-agent)

Run only the mq-store-agent with the following.

```bash
    docker pull ghcr.io/covalenthq/mq-store-agent:latest
    docker run ghcr.io/covalenthq/mq-store-agent:latest
```

Use `docker-compose` to get all the necessary services along with to also get running along with the mq-store-agent with the following from root. The other services are -

1. redis-srv (Open source (BSD licensed), in-memory data structure store)
1. redis-commander-web (Redis web management tool written in node.js)
1. ganache-cli (Ethereum blockchain & client)
1. proof-chain (Validation (proofing) smart-contracts)

```bash
    cd mq-store-agent
    docker-compose -f "docker-compose.yml" up --build --force-recreate --remove-orphans
```

## Scripts

![diagram](wow.jpeg)

There are two lua scripts in `/scripts` for usage with the `redis-cli`.

1. redis-count.lua - This allows for counting of total stream messages within bounds.

    -- call with params [stream-key] , [first-stream-id] [last-stream-id] 
    -- get to the ids with XINFO STREAM [stream-key]

```bash
> redis-cli --eval redis-count.lua replication , "1637349819851-0" "1637349831400-35"
> (integer) 6280
```

2. redis-trim.lua - This allows for removing messages from the stream within bounds.

    -- call with params [stream-key] , [number-of-elements-to-trim-from-start]
    -- get to the ids with XINFO STREAM [stream-key]

```bash
> redis-cli --eval redis-trim.lua replication , 5 
> (integer) 5
```

### Inspect

To view pretty print the results from the creation of avro encoded block-replica files

```bash
go run extractor.go \ 
    --binary-file-path="../bin/block-replica/" \ 
    --codec-path="../codec/block-replica.avsc" \ 
    --indent-json=0
```

Please make sure that the --binary-file-path and --codec-path matches the ones given while running the agent above. --indent-json (0,1,2) can be used to pretty print and inspect the AVRO json objects.
