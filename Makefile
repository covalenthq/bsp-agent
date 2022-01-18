docker:
	@docker-compose down
	@docker-compose build
	@docker-compose up -d

dockerdown:
	@docker-compose down

run-build:
	@echo "---- Building Agent from cmd/mqstoreagent ----"
	@go build -o ./bin/mqstoreagent/agent ./cmd/mqstoreagent/*.go 
	@echo "---- Done Building to ./bin/agent ----"

run-agent-eth:
	@echo "---- Running Agent from cmd/mqstoreagent ----"
	@go run ./cmd/mqstoreagent/*.go \
	--redis-url="redis://username:@localhost:6379/0?topic=replication#replicate" \
	--codec-path="./codec/block-replica.avsc" \
	--binary-file-path="./bin/block-ethereum/" \
	--gcp-svc-account="/Users/pranay/.config/gcloud/bsp-2.json" \
	--replica-bucket="covalenthq-geth-block-specimen" \
	--segment-length=5 \
 	--eth-client=http://127.0.0.1:7545  \
 	--proof-chain-address=0x3D25EBCeFC7F1E5a5664C8D6AA63Dc3D513761D6 \
	--consumer-timeout=80

run-agent-elrond:
	@echo "---- Running Agent from cmd/mqstoreagent ----"
	@go run ./cmd/mqstoreagent/*.go \
	--redis-url="redis://username:@localhost:6379/0?topic=replication#replicate" \
   	--codec-path="./codec/block-elrond.avsc" \
    --binary-file-path="./bin/block-elrond/" \
    --gcp-svc-account="/Users/pranay/.config/gcloud/bsp.json" \
    --replica-bucket="covalenthq-geth-block-specimen" \
    --segment-length=5 \
 	--eth-client=http://127.0.0.1:7545  \
 	--proof-chain-address=0x3D25EBCeFC7F1E5a5664C8D6AA63Dc3D513761D6 \
	--consumer-timeout=80 \
	--websocket-urls="34.66.210.112:20000"