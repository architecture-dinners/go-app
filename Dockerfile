FROM golang:latest

COPY . /go/src/github.com/go-app
WORKDIR /go/src/github.com/go-app

CMD go run server.go
