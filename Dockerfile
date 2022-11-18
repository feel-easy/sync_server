FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64\
    GOPROXY="https://goproxy.cn,direct"
WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o sync_server .

FROM scratch
COPY ./asset /asset
COPY --from=builder /build/sync_server /

ENTRYPOINT ["/sync_server"]