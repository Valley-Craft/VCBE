# Используем официальный образ Go в качестве базового образа
FROM golang:1.20

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы проекта в контейнер
COPY . .

# Собираем приложение
RUN go build -o main ./cmd/gohttp

ENV PORT=8080
ENV SERVER_KEY=8i8XtPNtLb24S87TkG82Sdktx4m5a8AZ
ENV RCON_PASSWORD=3gd3X9P5a347Yk3AyLbVnCx432beTbPF

# Указываем порт, на котором будет работать приложение
EXPOSE 8080

# Запускаем приложение при запуске контейнера
CMD ["./main"]
