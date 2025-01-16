docker:
	@docker-compose down
	@docker-compose -f "docker-compose.yml" up --build --remove-orphans --force-recreate --exit-code-from consumer

dockerdown:
	@docker-compose down

lint:
	@golangci-lint run

build:
	@echo "---- Building Agent from cmd/bspagent ----"
	@go build -o ./bin/bspagent ./cmd/bspagent/ 
	@echo "---- Done Building to ./bin/bspagent ----"

run-agent-eth:
	@echo "---- Running Agent from cmd/bspagent ----"
	@go run ./cmd/bspagent/*.go \
	--redis-url="redis://username:@localhost:6379/0?topic=replication-2#replicate-1"  \
	--avro-codec-path="./codec/block-ethereum.avsc"  \
	--binary-file-path="./bin/block-ethereum/" \
	--block-divisor=3  \
	--proof-chain-address="0x8243AF52B91649547DC80814670Dd1683F360E4c"  \
	--consumer-timeout=10000000  \
	--log-folder ./logs/ \
	--ipfs-pinner-server="http://127.0.0.1:3000/"

test:
	@echo "---- Testing Agent from cmd/bspagent ----"
	@go test ./... -coverprofile=coverage.out
	@echo "---- Done Testing for cmd/bspagent ----"

