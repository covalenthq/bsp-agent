# Build the project in the first phase. 
##
FROM golang:1.16-alpine as builder

WORKDIR /build

# COPY go.mod ./
# COPY go.sum ./
COPY . .

RUN go mod download

RUN apk update && apk add --no-cache bash
#RUN /bin/sh -c "apk add --no-cache bash"
# SHELL ["/bin/bash", "-c"]
# RUN chmod 777 /bin/bash

# Build the service
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-s -w" -o main ./cmd/mqstoreagent

RUN chmod 777 main
# ################################################################################
# # OUTPUT IMAGE
# # Copy artifacts from the builder and create an image with scratch

# FROM scratch
# LABEL author "Pranay Valson <pranay.valson@gmail.com>"
# LABEL maintainer "Pranay Valson <pranay.valson@gmail.com>"
# LABEL version "0.1"

# #RUN echo $(which bash)
# # Copy the binary from build-phase to service dir
# COPY --from=builder /build/main /srv/
# # Copy the AVRO codec from build-phase to service dir
# COPY --from=builder /build/codec /srv/codec/
# # Copy bash from build-phase to service dir
# COPY --from=builder /bin/ /usr/bin/
# # Copy SSL certs
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

#RUN export PATH=$PATH:/srv/bin/

RUN mkdir -p /bin/block-replica

ENTRYPOINT [ "/bin/bash", "-l", "-c" ]

CMD [ "./entry.sh" ]

EXPOSE 8080