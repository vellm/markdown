FROM golang:1.7.4-alpine

ADD . /go/src/github.com/HenrikFricke/markdown
RUN go install github.com/HenrikFricke/markdown

ENTRYPOINT /go/bin/markdown

EXPOSE 8080
