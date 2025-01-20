FROM golang:1.23


WORKDIR /go/src

COPY . /go/src

RUN go mod tidy
RUN go build main.go