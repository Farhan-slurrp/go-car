# Build the cipher service binary
FROM golang:1.12.7-stretch as base

# add the working directory for the project
WORKDIR /go/src/carlistingsvc

# Copy the service code
COPY authentication authentication
COPY cmd/carlisting cmd/carlisting
COPY database database
COPY internal internal
COPY pkg/carlisting pkg/carlisting
COPY vendor vendor
COPY go.mod go.mod
COPY go.sum go.sum

# building service binary at path discovergy/www
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOFLAGS=-mod=vendor -o carlisting go build ./cmd/carlisting

## Using the multi-stage image to just run the binary
FROM alpine:latest as final

# Init working directory to root /
WORKDIR /

# Copy just the binary from the base image
COPY --from=base /go/src/carlistingsvc/carlisting .

# just an indication that this port will be exposed by this container
EXPOSE 3334

# command to run at the immediate start of the container
ENTRYPOINT ["./carlisting"]