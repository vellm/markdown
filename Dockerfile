FROM golang:1.7.4-alpine

ADD . /go/src/github.com/vellm/vellm.io
RUN go install github.com/vellm/vellm.io

ENTRYPOINT /go/bin/vellm.io

EXPOSE 8080
