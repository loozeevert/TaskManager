FROM golang:1.23-alpine

WORKDIR /app

# Копируем зависимости ПЕРВЫМИ (для кэширования)
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем весь исходный код
COPY . .

# Устанавливаем gin для hot reload
RUN go install github.com/codegangsta/gin@latest

# Права на entrypoint.sh
RUN chmod +x entrypoint.sh

# Запуск
ENTRYPOINT ["./entrypoint.sh"]
