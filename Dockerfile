FROM golang:1.16-buster AS builder

WORKDIR /src

COPY go.mod go.sum ./