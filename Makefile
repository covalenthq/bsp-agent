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
	@go run ./cmd/mqstoreagent/*.go --redis-url "redis://username:@localhost:6379/0?topic=replication#replicate" --codec-path "./codec/" --result-target="covalenthq-geth-block-result" --specimen-target="covalenthq-geth-block-specimen" --segment-length=5 --gcp-svc-account="/Users/pranay/.config/gcloud/block-specimen-producer.json" --eth-client="http://127.0.0.1:7545" --proof-chain-address="0x54a878Ef5ecf8D7C114937e04b9984dE5f83344c"

run-agent-binary:
	@echo "---- Running Agent binary from bin/agent ----"
	@./bin/mqstoreagent/agent --redis-url "redis://username:@localhost:6379/0?topic=replication#replicate"  --codec-path "./codec/" --result-target="covalenthq-geth-block-result" --specimen-target="covalenthq-geth-block-specimen" --segment-length=5 --gcp-svc-account="/Users/pranay/.config/gcloud/block-specimen-producer.json" --eth-client="http://127.0.0.1:7545" --proof-chain-address="0x54a878Ef5ecf8D7C114937e04b9984dE5f83344c"