FROM golang:1.19-alpine as builder
RUN apt-get update && apt-get install -y \
                                 autoconf \
                                 automake \
                                 libtool \
                                 curl \
                                 make\
                                 g++ \
                                 unzip
ARG CACHEBUST
RUN mkdir -p $GOPATH/src/github.com/iblockin/CloudFlare_rss
RUN mkdir -p /work/build
WORKDIR $GOPATH/src/github.com/iblockin/CloudFlare_rss
COPY . $GOPATH/src/github.com/iblockin/CloudFlare_rss/
RUN go get github.com/lunny/html2md
RUN go get github.com/mmcdole/gofeed
WORKDIR $GOPATH/src/github.com/iblockin/CloudFlare_rss/cmd/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /work/build/CloudFlare_rss

FROM alpine:latest
ARG CACHEBUST
RUN apk update && apk --no-cache add ca-certificates
WORKDIR /work
COPY --from=builder /work/build .
RUN chmod +x /work/CloudFlare_rss
CMD [ "/work/CloudFlare_rss" ]
