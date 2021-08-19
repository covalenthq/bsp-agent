OS := linux
APP_NAME := mq-store-agent

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=$(OS) go build -a -installsuffix cgo -o main ./...

.PHONY: docker-build
docker_build: build
	docker-compose build

.PHONY: run
run: docker_build
	docker-compose up -d

