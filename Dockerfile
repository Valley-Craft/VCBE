# Используем официальный образ Go в качестве базового образа
FROM golang:1.20

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы проекта в контейнер
COPY . .

# Собираем приложение
RUN go build -o main ./cmd/gohttp

ENV PORT=80
ENV SERVER_KEY=8i8XtPNtLb24S87TkG82Sdktx4m5a8AZ
ENV RCON_PASSWORD=3gd3X9P5a347Yk3AyLbVnCx432beTbPF
ENV WEB_HOOK_URL=https://discord.com/api/webhooks/1112272251660812298/gRBriIzRsxYi4O-fjbrk8NyhK3kDpzhMUQBXaEQ_Ju7raxqdb7E_jRG32OTaghICXTEu
ENV X_KEY_DONATE=e07494d6-6b16-48ee-bc0b-62fbed75f5f0

# Указываем порт, на котором будет работать приложение
EXPOSE 80

# Запускаем приложение при запуске контейнера
CMD ["./main"]
