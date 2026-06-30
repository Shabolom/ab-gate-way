# AB Gateway

HTTP gateway для `ab-microservice` и auth-сервиса.

Сервис принимает HTTP-запросы от админки, валидирует входные данные, прокидывает авторизационные заголовки и ходит во внутренние gRPC-сервисы.

## Стек

* Go
* Echo
* gRPC
* OpenAPI
* Prometheus
* Docker Compose

## Возможности

* регистрация и авторизация пользователей;
* обновление token pair;
* получение текущего пользователя;
* получение списка пользователей;
* создание A/B-экспериментов;
* получение списка экспериментов;
* получение эксперимента по id;
* создание namespaces;
* получение списка namespaces;
* получение namespace по id;
* создание layers;
* получение списка layers;
* получение layer по id;
* создание custom params;
* получение списка custom params;
* получение custom param по id;
* создание feature toggles;
* получение списка feature toggles;
* получение feature toggle по id;
* изменение rollout percentage;
* изменение статуса feature toggle;
* проверка попадания пользователя в feature toggle.

## Структура проекта

```text
.
├── build/local              # локальное окружение
├── cmd/gate-way             # точка входа
├── docs/echo/openapi.yaml   # OpenAPI контракт
├── gen                      # сгенерированный gRPC-код
├── internal
│   ├── adapter              # gRPC adapters
│   ├── config               # конфигурация
│   ├── dto                  # HTTP DTO
│   ├── handler              # Echo handlers
│   ├── middleware           # auth middleware
│   ├── render               # HTTP responses/errors
│   └── service              # бизнес-логика gateway
├── pkg
├── docker-compose.yml
├── Makefile
├── go.mod
└── README.md
```

## Локальный запуск

```bash
git clone https://github.com/Shabolom/ab-gate-way.git
cd ab-gate-way
```

```bash
docker compose up -d
```

```bash
go mod download
```

```bash
go run ./cmd/gate-way
```

По умолчанию HTTP API доступен на:

```text
http://localhost:8080
```

## OpenAPI

Контракт лежит здесь:

```text
docs/echo/openapi.yaml
```

Основные группы ручек:

* `/v1/auth/...`
* `/v1/users/...`
* `/v1/experiments/...`
* `/v1/namespaces/...`
* `/v1/layers/...`
* `/v1/custom-params/...`
* `/v1/feature-toggles/...`

## Авторизация

Gateway ожидает токены в заголовках:

```http
Authorization: <access_token>
Refresh-Token: <refresh_token>
```

Эти заголовки прокидываются дальше во внутренние gRPC-сервисы через metadata.

## Основные ручки для админки

### Auth

```http
POST /v1/auth/register
POST /v1/auth/login
POST /v1/auth/logout
POST /v1/auth/refresh
```

### Users

```http
GET    /v1/users
GET    /v1/users/me
PATCH  /v1/users
DELETE /v1/users
```

### Experiments

```http
GET  /v1/experiments
GET  /v1/experiments/{id}
POST /v1/experiments
POST /v1/experiments/{experiment_id}/ready
POST /v1/experiments/{experiment_id}/stopped
POST /v1/experiments/user
```

### Namespaces

```http
GET  /v1/namespaces
GET  /v1/namespaces/{id}
POST /v1/namespaces
```

### Layers

```http
GET  /v1/layers
GET  /v1/layers/{id}
POST /v1/layers
```

### Custom Params

```http
GET  /v1/custom-params
GET  /v1/custom-params/{id}
POST /v1/custom-params
```

### Feature Toggles

```http
GET   /v1/feature-toggles
GET   /v1/feature-toggles/{id}
POST  /v1/feature-toggles
PATCH /v1/feature-toggles/{feature_toggle_id}/rollout
POST  /v1/feature-toggles/{feature_toggle_id}/status
GET   /v1/feature-toggles/{feature_toggle_id}/enabled
POST  /v1/feature-toggles/user
```

## Для админки

Админка должна общаться только с gateway.

Рекомендуемый frontend stack:

* Vue 3
* TypeScript
* Element Plus
* Pinia
* Vue Router
* Axios

Gateway является единой HTTP-точкой входа и скрывает внутренние gRPC-вызовы.
