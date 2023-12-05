FROM golang:alpine as builder

WORKDIR /server
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o main .

FROM alpine:latest

LABEL MAINTAINER="zzhang_xz@163.com"

WORKDIR /server

COPY --from=0 /server/main ./

EXPOSE 8080
ENTRYPOINT ./main
