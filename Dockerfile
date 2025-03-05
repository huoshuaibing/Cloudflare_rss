FROM golang:1.19-alpine AS builder
RUN apk add --no-cache \
    autoconf \
    automake \
    libtool \
    curl \
    make \
    g++ \
    unzip

ARG CACHEBUST

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY . .

RUN go mod download && go mod verify

WORKDIR /app/cmd/

RUN ls -l /app && ls -l /app/cmd

RUN go build -v -o /app/CloudFlare_rss

FROM alpine:latest
ARG CACHEBUST
RUN apk update && apk --no-cache add ca-certificates
WORKDIR /work
COPY --from=builder /app/CloudFlare_rss .
RUN chmod +x /work/CloudFlare_rss
CMD [ "/work/CloudFlare_rss" ]
