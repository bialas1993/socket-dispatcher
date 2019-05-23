FROM golang:1.12

LABEL maintainer="Dawid Biały <dawid_bialy@tvn.pl>"

WORKDIR /go/src/github.com/bialas1993/socket-dispatcher

ENV GO111MODULE=on

RUN go mod download

RUN ["go", "get", "github.com/go-playground/justdoit"]

ENTRYPOINT justdoit -watch="./" -include="(.+\\.go|.+\\.c)$" -build="go build ./cmd/daemon/" -run="SOCKET_DISPATCHER_PORTS=\"8000-8002\" ./daemon"