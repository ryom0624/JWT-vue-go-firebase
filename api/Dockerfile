FROM golang:1.12.9-alpine3.10

ENV GO111MODULE=on

WORKDIR /app
ADD . .

RUN apk update \
	&& apk add --no-cache git \
	&& go get gopkg.in/urfave/cli.v2@master \
	&& go get github.com/oxequa/realize

ENV CREDENTIAL=credentials/credential.json