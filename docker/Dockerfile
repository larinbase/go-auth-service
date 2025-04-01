FROM golang:1.23-alpine

WORKDIR /app

# Установка зависимостей для сборки
RUN apk add --no-cache git

# Копирование файлов проекта
COPY go.mod go.sum* ./
RUN go mod download

COPY . .

# Сборка приложения
RUN go build -o main ./app/main.go

# Запуск приложения
CMD ["./main"]