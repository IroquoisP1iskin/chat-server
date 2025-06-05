# Auth Service 🔐

Микросервис аутентификации и управления пользователями с gRPC API

## 📌 Функционал
- Регистрация/авторизация пользователей
- Управление ролями (user/admin)
- JWT-токены
- gRPC + HTTP gateway (REST)

## 🛠 Технологии
- **Go** 1.24+
- **gRPC** + Protocol Buffers
- **PostgreSQL** (хранение пользователей)
- **Redis** (сессии)
- **Prometheus** (метрики)

## 🚀 Запуск

### Требования
- Go 1.24+
- PostgreSQL 15+
- Redis 7+

```bash
# 1. Клонировать репозиторий
git clone https://github.com/IroquoisP1iskin/auth.git

# 2. Настроить .env (скопировать из .env.example)
cp .env.example .env

# 3. Запустить сервис
make run
