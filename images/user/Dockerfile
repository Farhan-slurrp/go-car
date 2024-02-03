# Build the cipher service binary
FROM golang:1.12.7-stretch as base

# add the working directory for the project
WORKDIR /go/src/usersvc

# Copy the service code
COPY authentication authentication
COPY cmd/user cmd/user
COPY database database
COPY internal internal
COPY pkg/user pkg/user
COPY vendor vendor
COPY go.mod go.mod
COPY go.sum go.sum

# args to put env variables
ARG google_client_id
ARG google_client_secret
ENV GOOGLE_CLIENT_ID $google_client_id
ENV GOOGLE_CLIENT_SECRET $google_client_secret

# building service binary at path discovergy/www
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOFLAGS=-mod=vendor -o user go build ./cmd/user

## Using the multi-stage image to just run the binary
FROM alpine:latest as final

# Init working directory to root /
WORKDIR /

# Copy just the binary from the base image
COPY --from=base /go/src/usersvc/user .

# just an indication that this port will be exposed by this container
EXPOSE 3334

# command to run at the immediate start of the container
ENTRYPOINT ["./user"]