# 构建环境
FROM golang:1.19-alpine AS builder
ARG APP_RELATIVE_PATH=./cmd/server
# 设置环境变量
ENV GO111MODULE=auto \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"
WORKDIR /data/app/
COPY ../../ .
RUN rm -rf ./bin/
RUN go build -ldflags="-s -w" -o ./bin/server ${APP_RELATIVE_PATH}
RUN mv config ./bin

## 运行环境
FROM alpine:latest
# 设置时区为上海
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

# 设置环境变量
ENV LANG=en_US.UTF-8

WORKDIR /data/app
COPY --from=builder /data/app/bin .

EXPOSE 8080
ENTRYPOINT [ "./server" ]