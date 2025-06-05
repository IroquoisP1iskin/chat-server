# Chat Server 💬

Микросервис для обмена сообщениями с gRPC API

## 📌 Функционал
- Создание/удаление чат-комнат
- Отправка сообщений в реальном времени
- История сообщений
- gRPC + WebSocket интерфейсы

## 🛠 Технологии
- **Go** 1.24+
- **gRPC** + Protocol Buffers
- **PostgreSQL** (хранение сообщений)
- **Redis** (pub/sub для реального времени)
- **NATS** (опционально для масштабирования)

## 🚀 Быстрый старт

### Требования
- Go 1.24+
- PostgreSQL 15+
- Redis 7+

```bash
# 1. Клонировать репозиторий
git clone https://github.com/IroquoisP1iskin/chat-server.git
cd chat-server

# 2. Настроить окружение
cp .env.example .env
nano .env  # задать свои параметры

# 3. Запустить
make run
