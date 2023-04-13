# Установка базового образа
FROM golang:latest

# Установка рабочей директории
WORKDIR /app

# Копирование файлов проекта в Docker-образ
COPY . .

# Установка зависимостей
RUN go mod download

# Установка порта
EXPOSE 8080

RUN go build -o app ./cmd/main/app.go

# Запуск приложения
CMD ["./app"]

