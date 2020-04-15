FROM golang:1.13-buster AS build
WORKDIR /root/http-dump
ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    && GO111MODULE="on" \
    && GOPRIVATE="git.inspii.com" \
    && GOSUMDB="off" \
    && GOPROXY="https://goproxy.cn,direct" \
    && go mod download \
    && go test --cover  $(go list ./... | grep -v /vendor/) \
    && cd cmd/server && go build -o main

FROM golang:1.13-buster
ENV TZ=Asia/Shanghai
WORKDIR /app
COPY --from=build /root/http-dump/cmd/server/main /app/http_dump
RUN chmod +x /app/http_dump

ENTRYPOINT ["/app/http_dump"]