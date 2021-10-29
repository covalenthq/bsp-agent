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
	@go run cmd/mqstoreagent/*.go