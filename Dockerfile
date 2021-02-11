FROM golang:1.13.6-alpine

ENV GO111MODULE=on

WORKDIR /redis-initiation

COPY go.mod .

RUN go mod download
