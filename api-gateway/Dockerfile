# Используем максимально легкий образ
FROM alpine:latest

WORKDIR /app

# Устанавливаем сертификаты (нужны для внешних API, например, Telegram)
RUN apk --no-cache add ca-certificates

# Копируем переименованный файл
COPY serverfile .
COPY .env .env

# Даем права на запуск именно этому файлу
RUN chmod +x serverfile

EXPOSE 8080

# Запускаем правильный файл
CMD ["./serverfile"]