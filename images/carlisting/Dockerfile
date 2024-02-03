FROM golang:1.21.6 as base

WORKDIR /go-car

COPY authentication authentication
COPY cmd/carlisting cmd/carlisting
COPY pkg/carlisting pkg/carlisting
COPY database database
COPY internal internal
COPY vendor vendor
COPY go.mod go.mod
COPY go.sum go.sum

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOFLAGS=-mod=vendor go build -o carlisting ./cmd/carlisting

FROM alpine:latest as final

WORKDIR /

COPY --from=base /go-car/carlisting .

EXPOSE 8082

ENTRYPOINT ["./carlisting"]