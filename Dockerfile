FROM golang:1.8.3

ADD . /go/src/github.com/ags799/learn-go

RUN go install github.com/ags799/learn-go

RUN rm -rf /go/src

ENTRYPOINT /go/bin/learn-go

EXPOSE 8080
