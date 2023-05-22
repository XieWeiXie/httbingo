FROM golang:1.20.0 as builder

ENV TZ=Asia/Shanghai
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
ARG APP=httpbingo

WORKDIR /app
RUN GOOS=linux go build -ldflags="-s -w" -o /app/${APP} main.go

FROM alpine

ENV TZ Asia/Shanghai
RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
RUN apt-get update && apt-get install -y ca-certificates

COPY --from=builder /app/${APP} /bin/httpbingo
EXPOSE 9091
CMD ["${APP}"]
