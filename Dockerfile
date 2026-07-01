FROM golang:1.24-alpine AS builder
WORKDIR /app
ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o blog-server-stats ./cmd/server

ENTRYPOINT ["/app/blog-server-stats"]