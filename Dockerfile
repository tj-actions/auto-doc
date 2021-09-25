FROM golang:1.17-alpine as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY cmd ./cmd

RUN go build -o auto_doc *.go

FROM alpine:3.14.2 as main

RUN apk add bash

COPY --from=builder /app/auto_doc /app/bin/auto_doc

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENV PATH="/app/bin:${PATH}"

ENTRYPOINT ["/entrypoint.sh"]

