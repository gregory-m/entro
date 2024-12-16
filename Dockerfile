FROM golang:1.23 AS builder
RUN go version
COPY . /go/src/github.com/gregory-m/entro
WORKDIR /go/src/github.com/gregory-m/entro
RUN set -Eeux && \
    go mod download && \
    go mod verify
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0\
    go build



FROM alpine:3.17.1
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/gregory-m/entro/entro .

EXPOSE 3333
ENTRYPOINT ["./entro"]