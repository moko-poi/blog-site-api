FROM golang:1.16.3-alpine AS build

ENV GO111MODULE=on

WORKDIR /

COPY . /go/src/engineer-jobhunting-api

RUN apk update && apk add --no-cache git
RUN cd /go/src/engineer-jobhunting-api/api && go build -o bin/build main.go

FROM alpine:3.8

COPY --from=build /go/src/engineer-jobhunting-api/api/bin/build /usr/local/bin/

CMD ["build"]
