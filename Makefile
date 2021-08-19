# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

.PHONY: test
test: 
	${GOTEST} -v  ./...

.PHONY: run
run:
	${MAKE} -C cmd/mqstoreagent run