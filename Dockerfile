FROM golang:1.17.5-alpine3.15 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/dplatform

FROM alpine:3.15
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /app/main .
CMD ["./main"]


