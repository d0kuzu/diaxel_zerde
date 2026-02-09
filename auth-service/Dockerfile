# Используем максимально легкий образ
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Устанавливаем сертификаты (нужны для внешних API)
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o auth-service .

FROM alpine:latest

WORKDIR /app

# Устанавливаем сертификаты (нужны для внешних API)
RUN apk --no-cache add ca-certificates

# Просто копируем ваш уже готовый файл server
COPY --from=builder /app/auth-service .
COPY .env .env

# Даем права на запуск (на всякий случай)
RUN chmod +x auth-service

EXPOSE 8080

CMD ["./auth-service"]