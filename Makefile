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

run-agent:
	@echo "---- Running Agent from cmd/mqstoreagent ----"
	@go run ./cmd/mqstoreagent/*.go \
	--redis-url="redis://username:@localhost:6379/0?topic=replication#replicate" \
	--codec-path="./codec/block-replica.avsc" \
	--binary-file-path="./bin/block-replica/" \
	--gcp-svc-account="/Users/pranay/.config/gcloud/bsp-2.json" \
	--replica-bucket="covalenthq-geth-block-specimen" \
	--segment-length=5 \
	--eth-client="http://127.0.0.1:7545" \
	--proof-chain-address="0xb5B12cbe8bABAF96677F60f65317b81709062C47"