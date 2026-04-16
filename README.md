# Gonio - Golang Gin 后端脚手架（Redis 限流、JWT、Clean Architecture）

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?logo=go)](https://go.dev/)
[![Gin](https://img.shields.io/badge/Gin-Web%20Framework-00A86B)](https://github.com/gin-gonic/gin)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](./LICENSE)

## 语言切换

- 中文版（
- [English](./README.en.md)

Gonio 是一个面向生产环境的 **Golang 后端脚手架**，基于 **Gin + GORM + Redis + MySQL**，内置 **IP + 路由限流（Redis Lua）**、JWT 认证、结构化日志、消息队列、多语言校验、优雅停机。

如果你在找 **Go API 起步模板**、**Gin + Redis 限流示例**，或者 **Clean Architecture 风格的后端脚手架**，Gonio 可以直接作为工程基础使用。

## 为什么选择 Gonio

- 面向高并发 API 服务：连接池、超时、重试、限流、日志链路完整
- 清晰分层架构：`Handler -> Service -> Repository`
- 中间件完善：RequestID、Logger、Recovery、CORS、I18n、Auth、Rate Limit
- 可扩展性强：通过 `ServiceContext` 统一依赖注入
- 可观测性友好：Zap 结构化日志 + DB/Redis 日志钩子

## 核心能力

- **Gin Web API 框架**
- **GORM + MySQL** 数据访问层
- **Redis 缓存与 Redis 分布式限流**
- 基于 Redis Lua 的 **IP + 路由限流**（原子执行）
- 面向 app/admin 的 **JWT 认证**
- 基于 Watermill 的消息队列（Redis Streams / MySQL）
- **I18n 校验与错误信息**
- **优雅停机**

## 目录结构

```text
.
├── cmd/
│   └── server/         # 服务启动入口
├── config/             # 配置文件
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
│   └── svc/            # ServiceContext 依赖注入
├── migration/
├── go.mod
└── Makefile
```

## 快速开始

### 环境要求

- Go 1.25+
- MySQL 8.0+
- Redis 6.0+

### 添加配置

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
  max_lifetime: 3600        # 连接最大存活时间(秒)
  conn_max_idle_time: 600   # 空闲连接最大存活时间(秒)
  dial_timeout: 5           # 连接超时(秒)
  read_timeout: 10          # 读超时(秒)
  write_timeout: 10         # 写超时(秒)
  ping_timeout: 3           # 探活超时(秒)
  prepare_stmt: true        # 缓存预编译语句，减少 SQL 解析开销
  skip_default_transaction: true  # 非事务查询不包裹事务，提升约 30% 性能

mq:
  driver: redis          # redis 或 mysql
  consumer_group: Gonio-group
  topic_concurrency:     # 按 topic 单独配置，未配置默认 1
    email: 3
    sms: 1
    stats: 2
  default_max_len: 20   # stream 全局默认最大长度
  topic_max_len:        # 按 topic 单独设置，优先级高于全局
    email: 10
    sms: 2000
    stats: 10000
  trim_interval: 3600   # 定期 XTRIM 间隔（秒），0 为不启用

redis:
  addr: 127.0.0.1:6379
  password: ""
  db: 2
  pool_size: 100
  min_idle_conns: 10        # 最小空闲连接数
  max_idle_conns: 50        # 最大空闲连接数
  pool_timeout: 5           # 获取连接池超时(秒)
  dial_timeout: 5           # 连接超时(秒)
  read_timeout: 3           # 读超时(秒)
  write_timeout: 3          # 写超时(秒)
  conn_max_idle_time: 300   # 空闲连接最大存活时间(秒)
  conn_max_lifetime: 3600   # 连接最大存活时间(秒)
  ping_timeout: 3           # 探活超时(秒)
  max_retries: 3            # 命令失败最大重试次数
  min_retry_backoff: 8      # 最小重试退避时间(毫秒)
  max_retry_backoff: 512    # 最大重试退避时间(毫秒)

jwt:
  secret: "Gonio-secret-key-change-me"
  expire: 604800 # seconds

log:
  mode: dev # dev / prod
  level: info
  sql_level: warn
  file_path: logs/app.log  # 日志文件路径，空则只输出到控制台
  max_size: 100            # 单文件最大 MB
  max_backups: 7           # 保留旧文件数
  max_age: 30              # 保留天数
  compress: true           # 压缩归档

smtp:
  host: smtp.example.com
  port: 465
  username: no-reply@example.com
  password: ""
  from: no-reply@example.com
  tls: true
```

### 运行项目

```bash
git clone https://github.com/your-username/Gonio.git
cd Gonio

go mod tidy
make run
```

健康检查：

```bash
curl http://localhost:8080/health
```

## 限流示例

Gonio 支持基于 Redis 的 API 限流，维度为 **IP + 路由 + 方法**。

- 商品列表接口：`1 request / 1 second`
- 商品创建接口：`1 request / 3 seconds`

## 典型场景

- 电商 API / 用户中心 / 管理后台
- 需要登录认证 + 限流 + 日志审计的业务系统
- 需要快速落地的 Go 微服务或单体 API 项目

## API 概览

- App API：`/app/v1/*`
- Admin API：`/admin/v1/*`
- 健康检查：`/health`

## 规划路线

- 滑动窗口 / 令牌桶限流策略
- OpenAPI / Swagger 文档
- Prometheus 指标与链路追踪集成

## 关键词

`Golang 限流器`、`Gin 限流中间件`、`Redis Lua 限流`、`IP 路由限流`、`Go API Rate Limiter`

## 参与贡献

欢迎提交 Issue 和 PR。

如果这个项目对你有帮助，欢迎给个 ⭐ 支持。

## 许可证

MIT License，详见 [LICENSE](./LICENSE)。
