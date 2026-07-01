FROM crpi-k33rnujh6dmvkp4x.cn-shenzhen.personal.cr.aliyuncs.com/zhangdx-cn/golang:1.24.13-alpine3.23 AS builder
WORKDIR /app
ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o blog-server-stats .

EXPOSE 8081

ENTRYPOINT ["/app/blog-server-stats"]