FROM golang:1.20.0 as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLE=0 \
    TZ=Asia/Shanghai

ARG APP=httpbingo

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY . .

RUN CGO_ENABLE=0 GOOS=linux go build -ldflags="-s -w" -o /app/target/${APP} main.go

FROM alpine:latest

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
RUN apt-get update && apt-get install -y ca-certificates

WORKDIR /app
COPY --from=builder /app/target/${APP} /app/${APP}
EXPOSE 9091
CMD ["/app/${APP}"]

