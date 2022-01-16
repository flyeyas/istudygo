FROM golang:1.17.5 as builder

ENV GO111MODULE=off \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /bbtgo

COPY . .

RUN go build -o bbtgo .

FROM scratch

COPY --from=builder /bbtgo/bbtgo /

EXPOSE 8080

ENTRYPOINT ["/bbtgo"]