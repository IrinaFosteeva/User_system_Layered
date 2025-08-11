# Layered User CRUD (Go + PostgreSQL)

Простой проект с **слоистой архитектурой**: repository -> service -> handler.
CRUD для сущности `User` (create, read, update, delete) с хранением в PostgreSQL.

## Структура
```
/ (root)
  go.mod
  main.go
  .env.example
  docker-compose.yml
  /db
    migrations.sql
  /models
    user.go
  /repository
    postgres_repo.go
  /service
    user_service.go
  /handler
    user_handler.go
```

## Быстрый запуск (локально с Docker)
1. Скопировать `.env.example` в `.env` и при необходимости изменить переменные.
2. `docker-compose up -d` — запустит postgres на `localhost:5432`.
3. Применить миграцию: `psql "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -f db/migrations.sql` (или используйте любой клиент)
4. `go run ./...` или `go run main.go`
5. API:
    - GET  `/users` — список
    - GET  `/users/{id}` — получить пользователя
    - POST `/users` — создать, body JSON `{ "name": "A", "email": "a@b.com" }`
    - PUT  `/users/{id}` — обновить
    - DELETE `/users/{id}` — удалить

## Docker compose
`docker-compose.yml` содержит сервис Postgres (пользователь: postgres, пароль: postgres, БД: postgres)

---