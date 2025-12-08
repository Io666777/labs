FROM golang:1.21-alpine

WORKDIR /app

# Устанавливаем переменные окружения для отключения модулей
ENV GO111MODULE=off

COPY main.go .

RUN go build -o main .

CMD ["./main"]