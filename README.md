# Микросервис для управления пользователями

---
## Описание
Разрабатываемый сервис управления пользователями предназначен для аутентификации, авторизации и управления профилями пользователей в веб-приложениях. Он обеспечивает безопасный доступ к системе, контроль прав пользователей и централизованное управление учетными записями. Данный сервис позволит разработчикам интегрировать его в свои проекты, обеспечивая удобную и гибкую систему аутентификации.

---

## Технологический стек
- Go 1.21+
- Gin Web Framework
- PostgreSQL
- JWT для аутентификации
- bcrypt для хеширования паролей
- Docker & Docker Compose

---
## База данных
- Таблица users с полями:
  - id (UUID)
  - email (unique)
  - password_hash
  - created_at
  - updated_at
- Таблица roles
  - id (uuid)
  - name
- Таблица refresh_sessions
  - id
  - created_at
  - expired_at
---

## API Endpoints
Все endpoints реализованы с использованием Gin Framework:
```
POST /api/v1/auth/register
Request:
{
    "email": "user@example.com",
    "password": "securepassword"
}
Response:
{   
    "data": {
        "access_token": "jwt.token.here",
        "refresh_token": "uuid"
    },
    "success": true
}

POST /api/v1/auth/login
Request:
{
    "email": "user@example.com",
    "password": "securepassword"
}
Response:
{   
    "data": {
        "access_token": "jwt.token.here",
        "refresh_token": "uuid"
    },
    "success": true
}

POST /api/v1/auth/refresh-tokens
Request:
{
  "access_token": "jwt.token.here",
  "refresh_token": "uuid"
}
Response:
{   
    "data": {
        "access_token": "jwt.token.here",
        "refresh_token": "uuid"
    },
    "success": true
}
```