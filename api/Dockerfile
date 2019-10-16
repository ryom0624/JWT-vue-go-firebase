FROM golang:1.12.9-alpine3.10

ENV GO111MODULE=on

WORKDIR /app
ADD . .

RUN apk update \
	&& apk add --no-cache git

CMD ["go", "run", "main.go"]