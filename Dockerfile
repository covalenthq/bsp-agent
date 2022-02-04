# Build - first phase. 
FROM golang:1.16-alpine as builder
RUN mkdir /build
WORKDIR /build
COPY . .
RUN go mod download
# Build the service
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-s -w" -o main ./cmd/mqstoreagent
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-s -w" -o extractor ./scripts/extractor.go

# Runtime/test -  second phase.
FROM alpine:3.15.0
RUN mkdir /app
WORKDIR /app
RUN apk update && apk add --no-cache bash=5.1.8-r0
RUN mkdir -p bin/block-ethereum bin/block-elrond
COPY --from=builder /build/main /app
COPY --from=builder /build/entry.sh /app
COPY --from=builder /build/data /app/data
COPY --from=builder /build/codec /app/codec
COPY --from=builder /build/extractor /app/scripts

ENTRYPOINT [ "/bin/bash", "-l", "-c" ]
CMD [ "./entry.sh" ]
EXPOSE 8080
