FROM golang:alpine as builder

MAINTAINER eapmp

RUN mkdir /build
WORKDIR /build

COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,direct && go env -w CGO_ENABLED=0
RUN go mod download
RUN go build -o chatgpt .


FROM alpine:latest

# change timezone
RUN apk add --no-cache tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo "Asia/Shanghai" >  /etc/timezone

RUN adduser -S -D -H -h /app appuser
USER appuser

WORKDIR /app

COPY --from=builder /build/chatgpt ./
# 这里后续调整成只copy config中的配置文件
# COPY --from=builder /build/config ./config/

EXPOSE 51515

CMD ["./chatgpt"]