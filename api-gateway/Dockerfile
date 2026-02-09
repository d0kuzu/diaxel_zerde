FROM golang:1.24-alpine

# Папка внутри контейнера, где будет лежать код
WORKDIR /app

# Копируем файлы зависимостей и скачиваем их
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь остальной код
COPY . .

# Собираем приложение
RUN go build -o main .

# Запускаем
CMD ["./main"]