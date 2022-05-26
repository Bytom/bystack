# Build Bystack in a stock Go builder container
FROM golang:1.18-alpine as builder

RUN apk add --no-cache make git

ENV GOPROXY=https://goproxy.cn,direct

ADD . /go/src/github.com/bytom/bystack
RUN cd /go/src/github.com/bytom/bystack && make bystackd && make bystackcli

# Pull Bystack into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/bytom/bystack/cmd/bystackd/bystackd /usr/local/bin/
COPY --from=builder /go/src/github.com/bytom/bystack/cmd/bystackcli/bystackcli /usr/local/bin/

EXPOSE 1999 46656 46657 9888
