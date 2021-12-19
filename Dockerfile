FROM golang:latest

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./cmd/dplatform .

EXPOSE 8080

RUN go build -o main .

CMD ./main

