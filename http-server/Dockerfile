FROM golang:1.26.1 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o http-server cmd/server/main.go

#####################
FROM scratch

WORKDIR /app

COPY --from=builder /app/http-server .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENTRYPOINT ["./http-server"]
