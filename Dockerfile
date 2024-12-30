FROM golang:1.23

RUN mkdir /code

WORKDIR /code

COPY . /code/

RUN go install
RUN go mod tidy
RUN go run main.go migrate