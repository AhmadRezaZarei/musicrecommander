# syntax=docker/dockerfile:1
FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build ./cmd/musicrecommander/main.go

EXPOSE 8569

CMD [ "./main" ]