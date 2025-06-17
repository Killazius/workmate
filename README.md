# Тестовое задание WORKMATE


## Функции

- Создание, удаление и проверка состояния
- Хранение в памяти
- Простая и расширяемая архитектура

## API ENDPOINTS

- `POST /tasks` - Создание новой задачи
- `GET /tasks/{id}` - Получить статус задачи
- `DELETE /tasks/{id}` - Удалить задачу

### Настройка окружения

1. Создайте файл `.env` на основе примера `.env.example`:
```env
CONFIG_PATH = "config/<name>.yaml"
```
Где `<name>` — название вашего конфигурационного файла. (по умолчанию config/config.yaml)

2. Создайте конфигурационный файл в папке config. Пример содержимого конфигурационного файла:
##### config/config.yaml
```yaml
http_server:
  port: "8080"
  host: "0.0.0.0"
  timeout: 5s
  idle_timeout: 60s
logger_path: "config/logger.json"
```
3. Создайте конфигурационный файл для логгера в папке config:
##### config/logger.json
```json
{
  "level": "debug",
  "encoding": "json",
  "outputPaths": ["stdout"],
  "errorOutputPaths": ["stderr"],
  "encoderConfig": {
    "timeKey": "timestamp",
    "timeEncoder": "rfc3339",
    "messageKey": "message",
    "levelKey": "level",
    "levelEncoder": "lowercase",
    "callerKey": "caller",
    "callerEncoder": "short"
  }
}
```

4. Запустите приложение при помощи Docker:
```bash
docker compose up -d --build
```