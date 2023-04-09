FROM golang:1.17 as builder

WORKDIR /app
ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /build
COPY server .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app go-kratos

FROM alpine:3.11
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app

COPY --from=builder /build/app ./
COPY --from=builder /build/config.yaml ./

ENV TZ=Asia/Shanghai
EXPOSE 80

ENTRYPOINT ["/app/app"]
