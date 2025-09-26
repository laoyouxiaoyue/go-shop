# Shop - 微服务电商用户管理系统

一个基于 Go 语言开发的微服务架构电商用户管理系统，采用 gRPC 和 RESTful API 的混合架构设计。

## 🏗️ 项目架构

### 微服务架构
- **用户服务 (User Service)**: gRPC 服务，负责用户数据管理和业务逻辑
- **Web 网关 (Web Gateway)**: HTTP RESTful API 服务，负责对外提供接口

### 技术栈
- **后端框架**: Gin (Web API) + gRPC
- **数据库**: MySQL + GORM
- **认证**: JWT
- **配置管理**: Viper
- **日志**: Zap
- **服务发现**: Consul
- **数据验证**: go-playground/validator

## 📁 项目结构

```
shop/
├── user/                    # 用户微服务 (gRPC)
│   ├── global/             # 全局配置
│   ├── handler/            # gRPC 处理器
│   ├── main.go            # 用户服务入口
│   ├── model/             # 数据模型
│   ├── proto/             # Protocol Buffers 定义
│   ├── service/           # 业务逻辑层
│   └── utils/             # 工具函数
├── web/                    # Web 网关服务
│   ├── api/               # HTTP API 处理器
│   ├── config/            # 配置结构
│   ├── forms/             # 表单验证结构
│   ├── global/            # 全局配置
│   ├── initialize/        # 初始化模块
│   ├── main.go           # Web 服务入口
│   ├── middlewares/       # 中间件
│   ├── models/           # 数据模型
│   ├── router/           # 路由配置
│   ├── utils/            # 工具函数
│   └── validator/        # 自定义验证器
├── go.mod                 # Go 模块依赖
└── README.md             # 项目说明文档
```

## 🚀 快速开始

### 环境要求
- Go 1.25.0+
- MySQL 5.7+
- Protocol Buffers 编译器

### 1. 克隆项目
```bash
git clone <repository-url>
cd shop
```

### 2. 安装依赖
```bash
go mod download
```

### 3. 配置数据库
创建 MySQL 数据库：
```sql
CREATE DATABASE shop CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

修改数据库连接配置：
- 用户服务: `user/global/global.go`
- 数据库迁移: `user/model/main/main.go`

### 4. 生成 Protocol Buffers 代码
```bash
# 进入 proto 目录
cd old_user/proto

# 生成 Go 代码
protoc -I . old_user.proto --go_out=.
protoc -I . old_user.proto --go-grpc_out=.
```

### 5. 运行数据库迁移
```bash
cd old_user/model/main
go run main.go
```

### 6. 启动服务

#### 启动用户服务 (gRPC)
```bash
cd old_user
go run main.go
# 默认端口: 5005
```

#### 启动 Web 网关服务
```bash
cd web
go run main.go
# 默认端口: 8021
```

## 📋 API 接口

### 用户相关接口

#### 1. 用户登录
```http
POST /v1/user/login/password
Content-Type: application/json

{
    "mobile": "13800138000",
    "password": "123456"
}
```

**响应:**
```json
{
    "msg": "登录成功"
}
```

#### 2. 获取用户列表
```http
GET /v1/user/list?pn=1&psize=10
Authorization: x-token <jwt_token>
```

**响应:**
```json
[
    {
        "id": 1,
        "nickname": "用户昵称",
        "birthday": "1990-01-01T00:00:00Z",
        "gender": "male",
        "mobile": "13800138000"
    }
]
```

## 🔧 配置说明

### Web 服务配置 (`web/config-debug.yaml`)
```yaml
name: 'old_user-web'
port: 8021
user_server:
  host: '127.0.0.1'
  port: 5005

jwt:
  key: '996948441'
```

### 数据库配置
默认配置：
- 主机: 127.0.0.1:3306
- 数据库: shop
- 用户名: root
- 密码: root

## 🛡️ 安全特性

- **JWT 认证**: 基于 JWT 的用户身份验证
- **密码加密**: 使用 MD5 加密存储用户密码
- **输入验证**: 自定义验证器验证手机号格式
- **CORS 支持**: 跨域请求处理

## 📊 数据模型

### User 模型
```go
type User struct {
    ID        int32     `json:"id"`
    Mobile    string    `json:"mobile"`      // 手机号 (唯一)
    Password  string    `json:"password"`    // 密码
    NickName  string    `json:"nick_name"`   // 昵称
    Birthday  *time.Time `json:"birthday"`   // 生日
    Gender    string    `json:"gender"`      // 性别
    Role      int       `json:"role"`        // 角色 (1:普通用户 2:管理员)
    CreatedAt time.Time `json:"created_at"`  // 创建时间
    UpdatedAt time.Time `json:"updated_at"`  // 更新时间
}
```

## 🔍 gRPC 服务接口

### User 服务方法
- `GetUserList(PageInfo) returns (UserListResponse)` - 获取用户列表
- `GetUserByMobile(MobileRequest) returns (UserInfoResponse)` - 根据手机号获取用户
- `GetUserById(IdRequest) returns (UserInfoResponse)` - 根据ID获取用户
- `CreateUser(CreateUserInfo) returns (UserInfoResponse)` - 创建用户
- `UpdateUser(UpdateUserInfo) returns (Empty)` - 更新用户信息
- `CheckPassword(PassWordCheckInfo) returns (CheckResponse)` - 验证密码

## 🧪 测试

### 运行测试
```bash
# 用户服务测试
cd old_user/service
go test -v

# Web 服务测试
cd web
go test -v
```

## 📝 开发指南

### 添加新的 API 接口
1. 在 `web/forms/` 中定义请求表单结构
2. 在 `web/api/` 中实现处理逻辑
3. 在 `web/router/` 中注册路由
4. 更新 gRPC 服务接口（如需要）

### 添加新的 gRPC 方法
1. 在 `user/proto/user.proto` 中定义服务方法
2. 重新生成 Protocol Buffers 代码
3. 在 `user/service/` 中实现业务逻辑
4. 在 `user/handler/` 中实现 gRPC 处理器

## 🐛 常见问题

### 1. 数据库连接失败
- 检查 MySQL 服务是否启动
- 验证数据库连接配置是否正确
- 确认数据库用户权限

### 2. gRPC 连接失败
- 检查用户服务是否正常启动
- 验证端口配置是否正确
- 确认防火墙设置

### 3. JWT 认证失败
- 检查 JWT 密钥配置
- 验证 token 格式是否正确
- 确认 token 是否过期

## 🤝 贡献指南

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建 Pull Request

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 👥 作者

- **Shuai** - 项目维护者

## 🙏 致谢

感谢所有为这个项目做出贡献的开发者和开源社区。
