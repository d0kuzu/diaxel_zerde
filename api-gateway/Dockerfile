# Используем максимально легкий образ
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN env GOPROXY=https://goproxy.io,direct GOSUMDB=off go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -p 1 -o serverfile .

FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/serverfile .

RUN chmod +x serverfile

EXPOSE 8080

CMD ["./serverfile"]