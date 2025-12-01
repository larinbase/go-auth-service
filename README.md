# Сервис авторизации и управления пользователями на Go

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15+-336791?style=flat&logo=postgresql)
![Docker](https://img.shields.io/badge/Docker-ready-2496ED?style=flat&logo=docker)

---

## Описание

Go Auth Service — микросервис для аутентификации, авторизации и управления профилями пользователей в современных веб-приложениях. Поддерживает вход по паролю, одноразовому коду, рефреш токены, интеграцию с Keycloak, смену пароля пользователя и безопасное хранение учетных данных. Идеален для разработки защищённых SaaS-платформ и b2b/b2c решений.

---

##  Ключевые возможности

- Регистрация/логин через email и пароль
- Вход по email + one-time code (`/auth/v2/sendCode`, `/auth/v2/login`)
- JWT для аутентификации и refresh токены
- Интеграция с Keycloak (`/auth/v3/keycloak/register`)
- Управление паролями (смена пароля)
- Гибкая архитектура (легко расширяется)
- Защищённые эндпоинты, работа с ролями

---

## Технологический стек

- **Go** 1.21+
- **Gin Web Framework**
- **PostgreSQL** (рекомендовано 15+)
- **JWT** (golang-jwt/jwt)
- **bcrypt** для паролей
- **Docker & Docker Compose**
- **Makefile** для автоматизации
- **Postman collection** для теста API

---

## ⚡ Быстрый старт
```
# 1. Клонируйте репозиторий
git clone https://github.com/larinbase/go-auth-service.git
cd go-auth-service

# 2. Сконфигурируйте .env
cp .env-example .env
# (отредактируйте переменные)

# 3. Запуск через Docker Compose
make up     # или docker-compose up --build

# API будет доступен по адресу http://localhost:8080
```
---

## 📚 API Reference

#### 1. **Регистрация**
```
POST /api/v1/auth/register
```
_Body:_
```json
{
  "email": "user@example.com",
  "password": "securepassword"
}
```
_Response:_
```json
{
  "data": {
    "access_token": "jwt.token.here",
    "refresh_token": "uuid"
  },
  "success": true
}
```

#### 2. **Логин по email/паролю**
```
POST /api/auth/login
```
_Body:_
```json
{
  "email": "user@example.com",
  "password": "securepassword"
}
```
_Response аналогичен регистрации._

#### 3. **Обновление токенов**
```
POST /api/auth/refresh-tokens
```

_Body:_
```json
{
  "access_token": "...",
  "refresh_token": "..."
}
```
_Response аналогичен регистрации._

#### 4. **Отправка кода на email (2FA или magic link)**
```
POST /api/auth/v2/sendCode
```

_Body:_
```json
{
  "email": "user@example.com"
}
```
#### 5. **Логин по коду**
```
POST /api/auth/v2/login
```

_Body:_
```json
{
  "email": "user@example.com",
  "code": 100000
}
```
#### 6. **Регистрация через Keycloak**
```
 POST /api/auth/v3/keycloak/register
 ```
 
_Body:_
```json
{
  "email": "user@example.com",
  "password": "securepassword"
}
```
---

### Пользователь
#### 1. **Смена пароля**
```
PATCH /api/user/change-password

Headers: Authorization: Bearer <token>
```

_Body:_
```json
{
  "old_password": "oldpassword",
  "new_password": "newpassword"
}
```
---
