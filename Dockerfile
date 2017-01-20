FROM golang:1.7.4-alpine

ADD . /go/src/github.com/vellm/vellm
RUN go install github.com/vellm/vellm

ENTRYPOINT /go/bin/vellm

EXPOSE 8080
