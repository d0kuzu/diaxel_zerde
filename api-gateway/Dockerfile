# Используем максимально легкий образ
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o serverfile .

FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/serverfile .
COPY .env .env

RUN chmod +x serverfile

EXPOSE 8080

CMD ["./serverfile"]