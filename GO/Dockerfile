FROM golang:1.25.1

WORKDIR /app

COPY text.go .

RUN go build -o app text.go

CMD ["./app"]