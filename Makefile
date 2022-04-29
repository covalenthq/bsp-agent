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
	--log-folder ./logs/  \

run-agent-elrond:
	@echo "---- Running Agent from cmd/bspagent ----"
	@go run ./cmd/bspagent/*.go \
	--redis-url="redis://username:@localhost:6379/0?topic=replication#replicate" \
	--avro-codec-path="./codec/block-elrond.avsc" \
	--binary-file-path="./bin/block-elrond/" \
	--gcp-svc-account="/Users/pranay/.config/gcloud/bsp-2.json" \
	--replica-bucket="covalenthq-geth-block-specimen" \
	--segment-length=1 \
	--proof-chain-address=0xbFCa723A2661350f86f397CEdF807D6e596d7874 \
	--consumer-timeout=80 \
	--websocket-urls="34.66.210.112:20000 34.66.210.112:20001 34.66.210.112:20002 34.66.210.112:20003" 
