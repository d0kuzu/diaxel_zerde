# Этап 1: сборка Go-приложения
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Копируем go.mod и go.sum и качаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем все файлы проекта
COPY . .

# Сборка бинарника
RUN go build -o server .

# Этап 2: минимальный образ для запуска
FROM alpine:latest

WORKDIR /app

# Устанавливаем сертификаты (нужно для HTTPS запросов)
RUN apk --no-cache add ca-certificates

# Копируем бинарник из builder
COPY --from=builder /internal/app/server .

# Порт, который слушает приложение
EXPOSE 8080

# Запуск сервера
CMD ["./server"]
