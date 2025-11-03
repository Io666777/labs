FROM golang:1.21-alpine

WORKDIR /app

# Копируем исходный код
COPY . .

# Скачиваем зависимости и собираем приложение
RUN go mod init app && \
    go get -u gorm.io/gorm gorm.io/driver/postgres && \
    go build -o app .

# Ждем пока PostgreSQL запустится и запускаем приложение
CMD sleep 10 && ./app