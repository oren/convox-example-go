FROM golang:1.6-alpine

RUN apk add --update git && rm -rf /var/cache/apk/*
ADD . /go/src/app
RUN cd /go/src/app && go get && go install

EXPOSE 3000

ENTRYPOINT ["/go/bin/app"]
