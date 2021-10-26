docker:
	@docker-compose down
	@docker-compose build
	@docker-compose up -d

dockerdown:
	@docker-compose down

build:
	@echo "---- Building Agent ----"
	@go build -o agent cmd/mqstoreagent/*.go

run:
	@echo "---- Running Agent ----"
	@export REDIS_HOST=localhost
	@export STREAM=replication
	@export GROUP=replicate-1
	@go run cmd/mqstoreagent/*.go