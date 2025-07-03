# HTTP客户端调用示例

这个项目演示了如何使用Go语言调用各种类型的HTTP API接口，包括不同的HTTP方法、数据格式和加密方式。

## 功能特性

### 1. HTTP方法支持
- **GET请求**: 获取用户信息
- **POST JSON**: 创建用户
- **POST Form**: 登录认证
- **POST Raw**: 文件上传

### 2. 数据格式支持
- JSON格式数据
- Form表单数据
- Raw二进制数据
- 加密数据

### 3. 认证方式
- API Key认证
- 自定义Token认证
- 表单登录认证

### 4. 高级功能
- 数据加密传输
- 重试机制
- 批量并发请求
- 错误处理

## 项目结构

```
http_client_demo/
├── main.go                    # 主程序，包含HTTP客户端示例
├── http_client.go             # HTTP客户端核心功能
├── http_client_util.go        # HTTP客户端工具方法（加密、重试等）
├── go.mod                     # Go模块文件
├── run_demo.sh                # 一键运行脚本（已废弃）
├── run_server.sh              # 启动服务器脚本
├── run_client.sh              # 启动客户端脚本
├── README.md                  # 项目说明文档
└── server/                    # 独立服务器模块
    ├── main.go                # 服务器主程序
    ├── simple_server.go       # 服务器实现
    └── go.mod                 # 服务器模块文件
```

## 快速开始

### 1. 安装依赖
```bash
# 安装客户端依赖
go mod tidy

# 安装服务器依赖
cd server && go mod tidy && cd ..
```

### 2. 运行示例

#### 方式一：分别启动（推荐）
```bash
# 终端1：启动服务器
./run_server.sh

# 终端2：启动客户端
./run_client.sh
```

#### 方式二：一键启动
```bash
# 程序会自动启动服务器并运行客户端演示
go run .
```

**注意**：推荐使用方式一，可以更好地观察服务器和客户端的运行状态。

## 示例说明

### 1. GET请求示例
```go
user, err := client.GetUser(1)
```
- 发送GET请求获取用户信息
- 包含Authorization头
- 处理JSON响应

### 2. POST JSON请求示例
```go
newUser := &User{
    Name:     "张三",
    Email:    "zhangsan@example.com",
    Password: "password123",
}
createdUser, err := client.CreateUser(newUser)
```
- 发送JSON格式的用户数据
- 创建新用户
- 返回创建的用户信息

### 3. POST Form请求示例
```go
token, err := client.LoginWithForm("zhangsan@example.com", "password123")
```
- 发送表单数据
- 用户登录认证
- 返回认证token

### 4. POST Raw数据请求示例
```go
fileData := []byte("这是文件内容")
err = client.UploadFile("test.txt", fileData)
```
- 发送原始二进制数据
- 文件上传功能
- 自定义请求头

### 5. 带加密的POST请求示例
```go
encryptionKey := []byte("your-32-byte-encryption-key-here")
encryptedUser, err := client.CreateUserWithEncryption(newUser, encryptionKey)
```
- 使用AES加密数据
- 安全传输敏感信息
- Base64编码传输

### 6. 带自定义Token认证的请求示例
```go
customTokenUser, err := client.GetUserWithCustomToken(1, "your-secret-key")
```
- 生成自定义认证token
- SHA256哈希加密
- 时间戳验证

### 7. 带重试机制的请求示例
```go
retryUser, err := client.GetUserWithRetry(1, 3)
```
- 自动重试失败请求
- 指数退避策略
- 最大重试次数限制

### 8. 批量请求示例
```go
userIDs := []int{1, 2, 3, 4, 5}
users, err := client.GetUsersBatch(userIDs)
```
- 并发处理多个请求
- 使用goroutine和channel
- 错误聚合处理

## 数据结构

### User结构体
```go
type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password,omitempty"`
}
```

### APIResponse结构体
```go
type APIResponse struct {
    Success bool        `json:"success"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}
```

## 加密功能

### AES加密
项目使用AES-256-GCM模式进行数据加密：
- 随机生成nonce
- 加密JSON数据
- Base64编码传输

### SHA256哈希
用于生成认证token和密码哈希：
- 时间戳+用户ID+密钥
- 生成唯一token
- 防止重放攻击

## 错误处理

所有HTTP请求都包含完整的错误处理：
- 网络连接错误
- HTTP状态码错误
- JSON解析错误
- 业务逻辑错误

## 配置说明

### API配置
```go
client := NewHTTPClient("http://localhost:8080", "your-api-key-here")
```
- 基础URL配置
- API Key认证
- 超时设置

### 加密配置
```go
encryptionKey := []byte("your-32-byte-encryption-key-here")
```
- 32字节AES密钥
- 用于数据加密
- 需要与服务器端保持一致

## 测试说明

项目包含一个简化的HTTP服务器用于测试：
- 支持所有示例接口
- 模拟真实API响应
- 包含认证验证
- 自动启动和关闭

## 扩展建议

1. **添加更多HTTP方法**: PUT、DELETE、PATCH等
2. **支持更多数据格式**: XML、Protocol Buffers等
3. **增强加密功能**: RSA、ECC等非对称加密
4. **添加连接池**: 提高并发性能
5. **支持代理**: HTTP/HTTPS代理配置
6. **添加监控**: 请求统计和性能监控
7. **支持压缩**: Gzip、Brotli等压缩格式

## 注意事项

1. 生产环境中请使用HTTPS
2. 妥善保管加密密钥
3. 实现适当的重试策略
4. 添加请求超时设置
5. 处理并发限制
6. 实现日志记录
7. 添加单元测试

## 许可证

MIT License 