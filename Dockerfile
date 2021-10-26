# Build the project in the first phase. 
FROM golang:1.16-alpine as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the service
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-s -w" -o main ./cmd/mqstoreagent/

################################################################################
# OUTPUT IMAGE
# Copy artifacts from the builder and create an image with scratch

FROM scratch
LABEL Author "Pranay Valson <pranay.valson@gmail.com>"
LABEL Maintainer "Pranay Valson <pranay.valson@gmail.com>"
LABEL Version "0.1"

# Copy the binary from build-phase
COPY --from=builder /build/main /srv/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /srv

# EXPOSE 8080

CMD [ "./main" ]
