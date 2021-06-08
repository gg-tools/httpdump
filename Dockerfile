FROM golang:1.16-alpine AS builder
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
ENV GOPRIVATE=""
ENV GOPROXY="https://goproxy.cn,direct"
ENV GOSUMDB="sum.golang.google.cn"

WORKDIR /root/httpdump/
ADD . .
RUN go mod download \
    && go test --cover $(go list ./... | grep -v /vendor/) \
    && go build -o main cmd/main.go

FROM alpine
WORKDIR /root/
ENV TZ Asia/Shanghai

RUN sed -e 's/dl-cdn[.]alpinelinux.org/mirrors.aliyun.com/g' -i~ /etc/apk/repositories
RUN apk --no-cache add curl

COPY --from=builder /root/httpdump/main httpdump
RUN chmod +x httpdump

ENTRYPOINT ["/root/httpdump"]