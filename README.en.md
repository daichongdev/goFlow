# Gonio - Golang Gin Backend Scaffold with Redis Rate Limiter, JWT Auth, and Clean Architecture

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?logo=go)](https://go.dev/)
[![Gin](https://img.shields.io/badge/Gin-Web%20Framework-00A86B)](https://github.com/gin-gonic/gin)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](./LICENSE)

## Language

- [Chinese Version](./README.md)
- English (current)

Gonio is a production-ready **Golang backend scaffold** built with **Gin + GORM + Redis + MySQL**. It includes **IP + route rate limiting (Redis Lua)**, JWT auth, structured logging, message queue integration, i18n validation, and graceful shutdown.

If you are looking for a **Go API starter template**, **Gin + Redis rate limiter example**, or a **clean architecture backend scaffold**, Gonio is built for that.

## Why Gonio

- Designed for high-concurrency APIs: pooling, timeout control, retries, rate limiting, and complete logging chain
- Clear layered architecture: `Handler -> Service -> Repository`
- Practical middleware stack: RequestID, Logger, Recovery, CORS, I18n, Auth, Rate Limit
- Strong extensibility via centralized dependency injection in `ServiceContext`
- Observability-friendly: Zap structured logging with DB/Redis log hooks

## Core Features

- **Gin Web API Framework**
- **GORM + MySQL** data access layer
- **Redis cache and Redis-based distributed rate limiter**
- **IP + Route Rate Limit** via Redis Lua script (atomic)
- **JWT Authentication** for app/admin
- **Watermill MQ** with Redis Streams / MySQL backend
- **I18n Validation and Error Messages**
- **Graceful Shutdown**

## Architecture

```text
.
├── cmd/
│   └── server/         # service entrypoint
├── config/             # config files
├── internal/
│   ├── config/
│   ├── database/
│   ├── handler/
│   ├── middleware/
│   ├── model/
│   ├── mq/
│   ├── pkg/            # errcode / response / validator / ratelimit
│   ├── repository/
│   ├── router/
│   ├── service/
│   └── svc/            # ServiceContext dependency injection
├── migration/
├── go.mod
└── Makefile
```

## Quick Start

### Requirements

- Go 1.25+
- MySQL 8.0+
- Redis 6.0+

### Add Config

```yaml
/config/config-dev.yaml
server:
  port: 8080
  mode: release # debug / release
  read_timeout: 10
  write_timeout: 10
  auto_migrate: true

mysql:
  host: 127.0.0.1
  port: 3306
  username: root
  password: "123456"
  database: silk_route
  max_idle_conns: 10
  max_open_conns: 100
  max_lifetime: 3600        # max connection lifetime (seconds)
  conn_max_idle_time: 600   # max idle connection lifetime (seconds)
  dial_timeout: 5           # dial timeout (seconds)
  read_timeout: 10          # read timeout (seconds)
  write_timeout: 10         # write timeout (seconds)
  ping_timeout: 3           # ping timeout (seconds)
  prepare_stmt: true        # cache prepared statements
  skip_default_transaction: true  # improve non-transaction query performance

mq:
  driver: redis          # redis or mysql
  consumer_group: Gonio-group
  topic_concurrency:     # per-topic concurrency, default 1
    email: 3
    sms: 1
    stats: 2
  default_max_len: 20    # global default stream max length
  topic_max_len:         # per-topic max length, higher priority than global
    email: 10
    sms: 2000
    stats: 10000
  trim_interval: 3600    # XTRIM interval (seconds), 0 disables trimming

redis:
  addr: 127.0.0.1:6379
  password: ""
  db: 2
  pool_size: 100
  min_idle_conns: 10
  max_idle_conns: 50
  pool_timeout: 5
  dial_timeout: 5
  read_timeout: 3
  write_timeout: 3
  conn_max_idle_time: 300
  conn_max_lifetime: 3600
  ping_timeout: 3
  max_retries: 3
  min_retry_backoff: 8
  max_retry_backoff: 512

jwt:
  secret: "Gonio-secret-key-change-me"
  expire: 604800 # seconds

log:
  mode: dev # dev / prod
  level: info
  sql_level: warn
  file_path: logs/app.log
  max_size: 100
  max_backups: 7
  max_age: 30
  compress: true

smtp:
  host: smtp.example.com
  port: 465
  username: no-reply@example.com
  password: ""
  from: no-reply@example.com
  tls: true
```

### Run

```bash
git clone https://github.com/your-username/Gonio.git
cd Gonio

go mod tidy
make run
```

Health check:

```bash
curl http://localhost:8080/health
```

## Rate Limiter Example

Gonio supports Redis-based API rate limiting by **IP + route + method**.

- Product list API: `1 request / 1 second`
- Product create API: `1 request / 3 seconds`

## Typical Use Cases

- E-commerce APIs / user center / admin backend
- Business systems requiring auth, rate limiting, and audit logs
- Go microservice or monolith API projects that need fast delivery

## API Overview

- App APIs: `/app/v1/*`
- Admin APIs: `/admin/v1/*`
- Health Check: `/health`

## Roadmap

- Sliding window / token bucket rate limit strategy
- OpenAPI / Swagger docs
- Prometheus metrics and tracing integration

## SEO Keywords

Golang backend scaffold, Gin boilerplate, Go web api template, Redis rate limiter, Gin rate limit middleware, JWT auth in Go, GORM MySQL starter, Clean Architecture Go.

## Contributing

Issues and PRs are welcome.

If this project helps you, please consider giving it a star on GitHub.

## License

MIT License. See [LICENSE](./LICENSE).
