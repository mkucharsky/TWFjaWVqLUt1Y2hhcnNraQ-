FROM golang:latest

LABEL maintainer="mkucharsky <maciek.kucharski.93@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/bin

RUN go build ./../cmd/web/
CMD ./web
