FROM alpine as certs
RUN apk update && apk add ca-certificates

FROM golang:1.16.6-alpine3.14 AS builder

WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build -o initdb -mod=vendor -ldflags='-s -w'  -installsuffix cgo ./cmd/main.go

FROM alpine
COPY --from=certs /etc/ssl/certs /etc/ssl/certs

WORKDIR /initdb
RUN mkdir migrations
COPY --from=builder ./build/initdb ./
COPY --from=builder ./build/migrations ./migrations/
ENTRYPOINT ["./initdb"]