FROM golang:1.24-alpine
# Устанавливаем прокси для ускорения загрузки и обхода проблем с сетью
ENV GOPROXY=https://proxy.golang.org,direct

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# 3. Только потом копируем код
COPY . .

# 4. Собираем (флаг -v покажет, какие пакеты компилируются сейчас)
RUN CGO_ENABLED=0 go build -v -o main .

CMD ["./main"]