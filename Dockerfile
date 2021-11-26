# Build the project in the first phase. 
##
## Build
##
FROM golang:1.16-alpine as builder
WORKDIR /build

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Build the service
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-s -w" -o main ./cmd/mqstoreagent/

################################################################################
# OUTPUT IMAGE
# Copy artifacts from the builder and create an image with scratch

FROM scratch
LABEL author "Pranay Valson <pranay.valson@gmail.com>"
LABEL maintainer "Pranay Valson <pranay.valson@gmail.com>"
LABEL version "0.1"

# Copy the binary from build-phase
COPY --from=builder /build/main /srv/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /srv

CMD ["./main", "--redis-url=redis://username:@redis:6379/0?topic=replication#replicate", "--codec-path=./codec/block-replica.avsc", "--binary-file-path=./bin/block-replica/", "--gcp-svc-account=/Users/pranay/.config/gcloud/bsp-2.json","--replica-bucket=covalenthq-geth-block-specimen", "--segment-length=5", "--eth-client=0.0.0.1:8545" , "--proof-chain-address=0xb5B12cbe8bABAF96677F60f65317b81709062C47"]

EXPOSE 8008