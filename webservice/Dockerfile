FROM golang:1.6-alpine

RUN apk add --update git && rm -rf /var/cache/apk/*
ADD . /go/src/app
ADD ./config.json /go/bin/config.json
WORKDIR /go/src/app
RUN go get && go install
WORKDIR /go/bin

EXPOSE 3000
ENTRYPOINT ["/go/bin/app"]
