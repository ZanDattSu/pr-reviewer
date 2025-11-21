# Этап 1: Сборка приложения

FROM golang:1.24.1 AS builder

WORKDIR /app

# Копируем файлы go.mod и go.sum отдельно для кеширования зависимостей
COPY go.mod go.sum ./

# Скачиваем зависимости заранее
RUN go mod download

COPY . .

# Отключаем CGO — не используем C-библиотеки, всё на чистом Go.
# Собираем бинарник для Alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd

# Этап 2: Минимальный образ для запуска приложения

FROM alpine:3.21.3

WORKDIR /app

# Создаём непривилегированного пользователя и группу с фиксированными UID/GID
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=builder /app/server .

COPY --from=builder /app/.env .

COPY --from=builder /app/migrations ./migrations

# Меняем владельца файлов на созданного пользователя
RUN chown -R appuser:appgroup /app

# Переключаемся на непривилегированного пользователя
USER appuser

EXPOSE 8080

CMD ["./server"]