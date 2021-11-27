# Build the project in the first phase. 
##
FROM golang:1.16-alpine as builder
WORKDIR /build

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN apk update && apk add --no-cache bash

SHELL ["/bin/bash", "-c"]

# Build the service
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-s -w" -o main ./cmd/mqstoreagent && apk update && apk --no-cache add bash
################################################################################
# OUTPUT IMAGE
# Copy artifacts from the builder and create an image with scratch

FROM scratch
LABEL author "Pranay Valson <pranay.valson@gmail.com>"
LABEL maintainer "Pranay Valson <pranay.valson@gmail.com>"
LABEL version "0.1"

# Copy bash from build-phase to service dir
COPY --from=builder /build/bin/bash /srv/bin/bash
# Copy the binary from build-phase to service dir
COPY --from=builder /build/main /srv/
# Copy the AVRO codec from build-phase to service dir
COPY --from=builder /build/codec /srv/codec/
# Copy SSL certs
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /srv

ENTRYPOINT [ "/bin/bash", "-l", "-c" ]

CMD ["./main", "--redis-url=redis://username:@redis:6379/0?topic=replication#replicate", "--codec-path=./codec/block-replica.avsc", "--binary-file-path=./bin/block-replica/", "--gcp-svc-account=/Users/user/.config/gcloud/bsp.json","--replica-bucket=covalenthq-geth-block-specimen", "--segment-length=10", "--eth-client=http://127.0.0.1:7545" , "--proof-chain-address=0xca59d70517cc581E2277EdCA8587A0dd2BeC5eb9", "--consumer-timeout=80"]

EXPOSE 8008