FROM golang:1.22-alpine

RUN apk update && apk add git

RUN mkdir /go/src/api

WORKDIR /go/src/api

COPY . /go/src/api

RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]
