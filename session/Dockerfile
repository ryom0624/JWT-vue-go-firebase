FROM golang:1.12.9-alpine3.10

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app
ADD . .

RUN apk update \
	&& apk add --no-cache git \
	&& go get gopkg.in/urfave/cli.v2@master \
	&& go get github.com/oxequa/realize \
	&& go get github.com/go-delve/delve/cmd/dlv \
	&& go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv

ENV CREDENTIAL=credentials/credential.json