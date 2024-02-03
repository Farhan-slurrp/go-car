FROM golang:1.21.6 as base

WORKDIR /go-car

COPY authentication authentication
COPY cmd/user cmd/user
COPY database database
COPY internal internal
COPY pkg/user pkg/user
COPY vendor vendor
COPY go.mod go.mod
COPY go.sum go-sum

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOFLAGS=-mod=vendor go build -o user ./cmd/user

FROM alpine:latest as final

WORKDIR /

COPY --from=base /go-car/user .

EXPOSE 8080
COPY .env .env

ENTRYPOINT ["./user"]
