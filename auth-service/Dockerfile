# Используем максимально легкий образ
FROM alpine:latest

WORKDIR /app

# Устанавливаем сертификаты (нужны для внешних API)
RUN apk --no-cache add ca-certificates

# Просто копируем ваш уже готовый файл server
COPY server .
COPY .env .env

# Даем права на запуск (на всякий случай)
RUN chmod +x server

EXPOSE 8080

CMD ["./server"]