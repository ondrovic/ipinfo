FROM golang:1.22.5-alpine AS builder

WORKDIR /app
COPY go.mod ./
COPY main.go ./

RUN go build -o ipinfo-app

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/ipinfo-app .

EXPOSE 8080

CMD ["./ipinfo-app"]