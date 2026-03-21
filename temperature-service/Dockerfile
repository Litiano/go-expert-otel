FROM golang:1.26.1 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o temperature cmd/server/main.go

#####################
FROM scratch

WORKDIR /app

COPY --from=builder /app/temperature .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENTRYPOINT ["./temperature"]

#EXPOSE 8090
