# Go Todo API
![Go Version](https://img.shields.io/badge/go-1.25.5-blue.svg)
![Gin Framework](https://img.shields.io/badge/framework-Gin-green.svg)
![Docker Supported](https://img.shields.io/badge/docker-supported-blue.svg)
![License](https://img.shields.io/badge/license-MIT-red.svg)
一个基于 Gin + GORM 的任务管理后端应用，支持用户认证、任务创建、查询、更新和删除功能。

## 🚀 功能特性

- ✅ **用户认证** - 注册、登录、JWT 身份验证
- ✅ **任务管理** - 创建、查询、更新、删除任务
- ✅ **权限控制** - 基于 JWT 的请求认证中间件
- ✅ **数据持久化** - 支持 MySQL 和 SQLite 数据库
- ✅ **API 文档** - Swagger/OpenAPI 自动化文档
- ✅ **容器化部署** - Docker 和 Docker Compose 支持
- ✅ **日志和监控** - 请求日志记录和 CORS 支持

## 📋 技术栈

| 组件 | 技术 | 版本 |
|------|------|------|
| 框架 | Gin Web Framework | v1.11.0 |
| ORM | GORM | v1.31.1 |
| 数据库驱动 | MySQL | v1.6.0 |
| 数据库驱动（测试） | SQLite | v1.11.0 |
| 认证 | JWT | v5.3.0 |
| 加密 | golang.org/x/crypto | v0.47.0 |
| 配置管理 | Viper | v1.21.0 |
| API 文档 | Swagger/Swag | v1.6.1 |
| Go 版本 | 1.25.5 | |

## 📦 项目结构

```
go-todo/
├── main.go                 # 应用入口
├── go.mod                  # Go 模块文件
├── config.yaml             # 配置文件
├── Dockerfile              # Docker 镜像构建配置
├── docker-compose.yml      # Docker Compose 服务编排
├── common/                 # 公共工具函数
│   ├── jwt.go              # JWT 令牌处理
│   └── response.go         # 统一响应格式
├── config/                 # 配置管理
│   └── database.go         # 数据库连接配置
├── controllers/            # 控制器层（业务逻辑）
│   ├── user_controller.go  # 用户相关接口
│   └── todo.go             # 任务相关接口
├── middleware/             # 中间件
│   ├── auth.go             # JWT 认证中间件
│   ├── cors.go             # CORS 跨域中间件
│   └── logger.go           # 日志中间件
├── models/                 # 数据模型
│   ├── user.go             # 用户模型
│   └── todo.go             # 任务模型
├── routes/                 # 路由定义
│   └── routes.go           # 路由配置
├── service/                # 业务服务层
│   ├── user_service.go     # 用户服务
│   ├── todo_service.go     # 任务服务
│   └── todo_service_test.go# 任务服务测试
└── docs/                   # API 文档
    ├── docs.go             # Swagger 文档生成文件
    ├── swagger.json        # Swagger JSON 文档
    └── swagger.yaml        # Swagger YAML 文档
```

## 🔌 API 端点

### 认证接口

| 方法 | 端点 | 描述 |
|------|------|------|
| POST | `/api/v1/auth/register` | 用户注册 |
| POST | `/api/v1/auth/login` | 用户登录 |

### 任务接口（需要认证）

| 方法 | 端点 | 描述 |
|------|------|------|
| POST | `/api/v1/todos` | 创建新任务 |
| GET | `/api/v1/todos` | 获取所有任务 |
| GET | `/api/v1/todos/:id` | 获取单个任务 |
| PUT | `/api/v1/todos/:id` | 更新任务 |
| DELETE | `/api/v1/todos/:id` | 删除任务 |

### 文档接口

| 方法 | 端点 | 描述 |
|------|------|------|
| GET | `/swagger/*any` | Swagger UI 文档 |

## 🛠️ 安装与运行

### 前置要求

- Go 1.25.5 或更高版本
- MySQL 8.0（可选，可使用 SQLite）
- Docker & Docker Compose（可选，用于容器化部署）

### 本地运行

1. **克隆项目**

```bash
git clone <repository-url>
cd go-todo
```

2. **安装依赖**

```bash
go mod download
go mod tidy
```

3. **配置数据库**

编辑 `config.yaml`，配置你的数据库连接信息：

```yaml
server:
  port: 8080

database:
  username: "root"
  password: "your_password"
  host: "127.0.0.1"
  port: 3306
  dbname: "todo_db"
```

4. **运行应用**

```bash
go run main.go
```

应用将在 `http://localhost:8080` 启动，API 文档可访问 `http://localhost:8080/swagger/index.html`

### Docker 运行

#### 使用 Docker Compose（推荐）

```bash
docker-compose up --build
```

这会同时启动应用服务和 MySQL 数据库。应用将在 `http://localhost:8080` 可访问。

#### 使用 Docker 单独构建

```bash
docker build -t go-todo:latest .
docker run -p 8080:8080 go-todo:latest
```

## 📝 使用示例

### 用户注册

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123"
  }'
```

### 用户登录

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123"
  }'
```

响应中会获得 JWT 令牌，用于后续的认证请求。

### 创建任务

```bash
curl -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_jwt_token>" \
  -d '{
    "title": "完成项目文档",
    "description": "编写详细的 README 和 API 文档",
    "status": "pending"
  }'
```

### 获取所有任务

```bash
curl -X GET http://localhost:8080/api/v1/todos \
  -H "Authorization: Bearer <your_jwt_token>"
```

### 更新任务

```bash
curl -X PUT http://localhost:8080/api/v1/todos/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_jwt_token>" \
  -d '{
    "status": "completed"
  }'
```

### 删除任务

```bash
curl -X DELETE http://localhost:8080/api/v1/todos/1 \
  -H "Authorization: Bearer <your_jwt_token>"
```

## 🔐 安全特性

- **密码加密** - 使用 `golang.org/x/crypto` 进行密码散列和验证
- **JWT 认证** - 使用 JWT 进行无状态身份验证
- **CORS 保护** - 配置了跨域请求处理
- **中间件保护** - 所有受保护的路由都需要有效的 JWT 令牌

## 📊 数据模型

### User（用户）

```go
type User struct {
    ID        uint
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt
    Username  string  // 唯一
    Password  string  // 加密存储
    Todos     []Todo  // 一对多关系
}
```

### Todo（任务）

```go
type Todo struct {
    ID        uint
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt
    Title     string
    Description string
    Status    string
    UserID    uint    // 外键
}
```

## 🧪 测试

运行单元测试：

```bash
go test ./...
```

运行特定包的测试：

```bash
go test ./service -v
```

## 📚 配置说明

### config.yaml

- `server.port` - 服务端口（默认：8080）
- `database.username` - 数据库用户名
- `database.password` - 数据库密码
- `database.host` - 数据库主机
- `database.port` - 数据库端口
- `database.dbname` - 数据库名称

Viper 支持环境变量覆盖，可通过设置 `DATABASE_HOST`、`DATABASE_PASSWORD` 等环境变量来覆盖配置文件中的值。

## 🐛 常见问题

### 无法连接到数据库

- 确认 MySQL 服务正在运行
- 检查 `config.yaml` 中的数据库凭证是否正确
- 确保数据库已创建：`CREATE DATABASE todo_db;`

### JWT 令牌过期

- 使用 `/api/v1/auth/login` 重新登录获取新令牌
- 令牌过期时间在应用代码中配置

### Docker Compose 启动失败

- 检查端口 3306 和 8080 是否被占用
- 查看日志：`docker-compose logs`
- 确保 Docker daemon 正在运行


