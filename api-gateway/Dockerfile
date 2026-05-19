# Используем максимально легкий образ alpine
FROM alpine:latest

WORKDIR /app

# Устанавливаем сертификаты безопасности (необходимы для HTTPS запросов)
RUN apk --no-cache add ca-certificates

# Копируем предварительно скомпилированный бинарный файл
COPY serverfile .

# Делаем файл исполняемым
RUN chmod +x serverfile

EXPOSE 8080

CMD ["./serverfile"]