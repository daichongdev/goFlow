# GoFlow

GoFlow 是一个基于 Go 语言的高性能、现代化 Web 服务脚手架。它采用了清晰的 **Handler-Service-Repository** 分层架构，并集成了多种企业级特性，旨在帮助开发者快速构建稳定、可扩展的后端应用。

## 🚀 特性

- **现代化技术栈**: 基于 [Gin](https://github.com/gin-gonic/gin) (Web 框架), [GORM](https://gorm.io/) (ORM), [Redis](https://github.com/redis/go-redis) (缓存)、限流器(Rate-limit) 和 [Zap](https://github.com/uber-go/zap) (日志)。
- **优雅的架构设计**: 严格遵循分层架构，通过 `ServiceContext` 实现统一的依赖注入 (DI)。
- **基础设施优化**:
    - **并行初始化**: MySQL 与 Redis 并行启动，提升服务响应速度。
    - **连接池管理**: 预配置高性能数据库连接池。
- **完善的中间件**:
    - 结构化日志 (Zap) & 请求追踪 (RequestID)
    - 异常恢复 (Recovery)
    - 跨域支持 (CORS)
    - 多语言支持 (I18n)
- **业务特性**:
    - **JWT 认证**: 集成用户/管理员双端认证。
    - **异步任务**: 集成 [Watermill](https://github.com/ThreeDotsLabs/watermill) 消息队列，支持 Redis Streams/MySQL 驱动。
    - **多语言验证**: 支持请求参数的多语言校验与翻译。
    - **平滑停机**: 支持 OS 信号捕获与优雅停机 (Graceful Shutdown)。
- **开发工具**: 包含 Makefile 脚本，支持环境配置切换 (dev/prod)。

## 📁 项目结构

```text
.
├── cmd/                # 入口目录
│   └── server/         # API 服务启动入口
├── configs/            # 配置文件 (YAML)
├── internal/           # 私有业务逻辑
│   ├── config/         # 配置解析
│   ├── database/       # 数据库连接 (MySQL, Redis)
│   ├── handler/        # 接口处理器 (HTTP 层)
│   ├── middleware/     # 中间件 (Auth, Logger, I18n 等)
│   ├── model/          # 数据模型 (DB 实体)
│   ├── mq/             # 消息队列 (Publisher, Router)
│   ├── pkg/            # 内部通用工具 (Response, ErrCode, Validator)
│   ├── repository/     # 数据仓库层 (DB 交互)
│   ├── router/         # 路由定义
│   ├── service/        # 业务逻辑层
│   └── svc/            # 依赖注入容器 (ServiceContext)
├── migration/          # 数据库迁移脚本
├── go.mod              # Go 模块管理
└── Makefile            # 构建与运行指令
```

## 🛠️ 快速开始

### 1. 环境要求
- Go 1.25+
- MySQL 8.0+
- Redis 6.0+

### 2. 安装与运行
```bash
# 克隆项目
git clone https://github.com/your-username/goflow.git
cd goflow

# 安装依赖
go mod tidy

# 准备配置文件 (修改 configs/config.yaml 中的数据库信息)
cp configs/config.yaml.example configs/config.yaml

# 运行服务
make run
```

### 3. API 测试
服务启动后，可以通过 `http://localhost:8080/health` 检查运行状态。

## 🔧 开发建议

- **新增业务**: 遵循 `Model -> Repository -> Service -> Handler -> Router` 的开发流程。
- **依赖管理**: 所有外部依赖均在 `internal/svc/service_context.go` 中初始化并注入。
- **错误处理**: 使用 `internal/pkg/errcode` 定义统一的业务错误码。

## 📄 开源协议
本项目采用 [MIT](LICENSE) 协议。
