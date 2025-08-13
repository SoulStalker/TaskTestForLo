# TaskTestForLo
# Task API (Go + Gin, In‑Memory, Async Logging, Graceful Shutdown)


## Эндпоинты
- `GET /tasks?status=todo|doing|done` — список задач (фильтрация по статусу).
- `GET /tasks/{id}` — получить задачу по ID.
- `POST /tasks` — создать задачу. Тело:
  ```json
  { "title": "string", "status": "todo|doing|done" }
  ```

## Сборка и запуск
```bash
git clone https://github.com/SoulStalker/TaskTestForLo.git
cd TestTaskFoLo
go mod tidy
go run cmd/server/main.go
# сервер слушает :8080
```

## Примеры
```bash
# Создать задачу
curl -X POST http://localhost:8080/tasks   -H 'Content-Type: application/json'   -d '{"title":"Купить хлеб","status":"todo"}' # можно без статуса

# Получить по ID
curl http://localhost:8080/tasks/1

# Список всех
curl http://localhost:8080/tasks

# Список по статусу
curl http://localhost:8080/tasks?status=todo
```

## Архитектура
```
cmd/server/main.go                # композиция зависимостей, graceful shutdown
internal/domain                   # доменные типы (Task, Status), интерфейс репозитория
internal/usecase                  # бизнес‑логика (TaskService)
internal/repo/memory              # in‑memory репозиторий (map + RWMutex)
internal/delivery/http            # HTTP хендлеры и роутер (Gin)
internal/logger                   # асинхронный логгер с каналом и отдельной горутиной
```

### Асинхронное логирование
- `logger.AsyncLogger` пишет JSON‑события в `stdout` через буферизированный канал.
- При переполнении канала события **не блокируют** обработчики — запись дропается с фиксированным сообщением.
- На shutdown вызывается `AsyncLogger.Shutdown(ctx)` для корректного завершения воркера.

### Graceful shutdown
- Использует `signal.NotifyContext` для перехвата `SIGINT/SIGTERM`.
- `http.Server.Shutdown(ctx)` завершает приём новых соединений и ждёт активные.
- Логгер корректно закрывается после HTTP‑сервера.

## Тестирование
Фокус сделан на минимальном, но чистом каркасе. Для unit‑тестов можно мокать `TaskRepository` и инжектить через `TaskService`.

## Замечания
- Хранилище — in‑memory, данные пропадают при перезапуске.
- Все времена — UTC.
