FROM golang:1.14-alpine AS builder

### add make gcc g++ python if you will use native dependencies
RUN mkdir -p $GOPATH/src/github.com/ar4mirez/say-what && \
apk add --no-cache git openssh-client

WORKDIR $GOPATH/src/github.com/ar4mirez/say-what

ENV GO111MODULE=on

COPY . .

RUN chmod 700 *

WORKDIR $GOPATH/src/github.com/ar4mirez/say-what

RUN go build ./... && go build

FROM alpine

RUN apk --no-cache add ca-certificates bash

WORKDIR /root/

COPY --from=builder /go/src/github.com/ar4mirez/say-what .

ENV PORT=3000

RUN chmod +x ./say-what

EXPOSE $PORT

ENTRYPOINT [ "./say-what" ]
