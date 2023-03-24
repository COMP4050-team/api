FROM --platform=linux/amd64 golang:1.20-alpine as builder

RUN apk add build-base

WORKDIR /app

COPY . /app

RUN go build -o main ./cmd/server.go

FROM --platform=linux/amd64 alpine:latest

WORKDIR /app

COPY --from=builder /app/main /app

CMD ["/app/main", "-test-executor-endpoint=https://test-executor-rs.fly.dev/"]