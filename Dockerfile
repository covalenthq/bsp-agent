# Build the project in the first phase. 
FROM golang:1.16-alpine as builder

WORKDIR /build

COPY . .

RUN go mod download

RUN apk update && apk add --no-cache bash=5.1.8-r0

# Build the service
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-s -w" -o main ./cmd/mqstoreagent

RUN mkdir -p bin/block-replica

ENTRYPOINT [ "/bin/bash", "-l", "-c" ]

CMD [ "./entry.sh" ]

EXPOSE 8080
