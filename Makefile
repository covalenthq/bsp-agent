docker:
	@docker-compose down
	@docker-compose build
	@docker-compose up -d

dockerdown:
	@docker-compose down

build:
	@echo "---- Building Agent ----"
	@go build cmd/mqstoreagent/*.go -o ./../../bin/agent
	@echo "---- Done Building Agent ----"

run:
	@echo "---- Running Agent ----"
	@go run cmd/mqstoreagent/*.go --redis-url "redis://username:@localhost:6379/0?topic=replication#replicate"