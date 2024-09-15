# Шаг 1: Используем официальный образ Golang для сборки приложения
FROM golang:1.23 AS builder

# Устанавливаем рабочую директорию в контейнере
WORKDIR /app

# Копируем go.mod и go.sum и загружаем зависимости
COPY go.mod go.sum ./
RUN go mod tidy

# Копируем все файлы проекта в рабочую директорию
COPY . .

# Собираем приложение
RUN go build -o kode_test ./cmd/api

# Шаг 2: Создаем минимальный финальный контейнер
FROM gcr.io/distroless/base

# Копируем скомпилированное приложение из контейнера builder
COPY --from=builder /app/kode_test /kode_test

# Открываем порт
EXPOSE 8080

# Указываем команду для запуска приложения
CMD ["/kode_test"]