FROM golang:1.24-alpine

WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Копируем proto (если он используется в replace)
COPY proto ./proto

RUN go mod download

# Копируем остальной код
COPY . .

RUN go build -o main .

CMD ["./main"]
