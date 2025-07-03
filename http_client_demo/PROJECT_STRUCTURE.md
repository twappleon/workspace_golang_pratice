# HTTP客户端项目结构说明

## 项目拆分说明

本项目将原来的单一文件拆分为多个独立的模块，便于维护和扩展。

## 文件结构

```
http_client_demo/
├── main.go                    # 客户端主程序入口
├── http_client.go             # HTTP客户端核心功能
├── http_client_util.go        # HTTP客户端工具方法
├── go.mod                     # 客户端模块文件
├── run_server.sh              # 启动服务器脚本
├── run_client.sh              # 启动客户端脚本
├── README.md                  # 项目说明文档
├── PROJECT_STRUCTURE.md       # 项目结构说明（本文件）
└── server/                    # 独立服务器模块
    ├── main.go                # 服务器主程序入口
    ├── simple_server.go       # 服务器实现
    └── go.mod                 # 服务器模块文件
```

## 模块说明

### 客户端模块 (http_client_demo/)

#### main.go
- **功能**: 客户端主程序入口
- **职责**: 
  - 启动服务器进程
  - 创建HTTP客户端实例
  - 执行各种HTTP请求演示
  - 展示结果

#### http_client.go
- **功能**: HTTP客户端核心功能
- **包含**:
  - `User` 结构体定义
  - `APIResponse` 结构体定义
  - `HTTPClient` 结构体定义
  - `NewHTTPClient()` 构造函数
  - 基础HTTP方法：`GetUser()`, `CreateUser()`, `LoginWithForm()`, `UploadFile()`

#### http_client_util.go
- **功能**: HTTP客户端工具方法
- **包含**:
  - 加密相关：`CreateUserWithEncryption()`, `encryptData()`
  - 认证相关：`GetUserWithCustomToken()`, `generateHash()`
  - 高级功能：`GetUserWithRetry()`, `GetUsersBatch()`

### 服务器模块 (server/)

#### main.go
- **功能**: 服务器主程序入口
- **职责**: 启动HTTP模拟服务器

#### simple_server.go
- **功能**: HTTP服务器实现
- **包含**:
  - `SimpleServer` 结构体定义
  - 路由处理：用户管理、登录、文件上传、加密用户创建
  - API Key验证
  - 响应格式化

## 运行方式

### 1. 分别启动（推荐）

```bash
# 终端1：启动服务器
./run_server.sh

# 终端2：启动客户端
./run_client.sh
```

### 2. 一键启动

```bash
# 程序会自动启动服务器并运行客户端演示
go run .
```

### 3. 独立运行

```bash
# 只运行服务器
cd server && go run .

# 只运行客户端（需要服务器已启动）
go run .
```

## 优势

1. **模块化**: 客户端和服务器完全分离，可以独立开发和部署
2. **可维护性**: 代码按功能拆分，便于维护和扩展
3. **可重用性**: 客户端模块可以在其他项目中重用
4. **测试友好**: 可以独立测试客户端和服务器
5. **部署灵活**: 可以分别部署到不同的服务器

## 扩展建议

1. **添加配置文件**: 支持从配置文件读取服务器地址、API Key等
2. **添加日志系统**: 使用结构化日志记录请求和响应
3. **添加监控**: 集成Prometheus等监控系统
4. **添加测试**: 为每个模块添加单元测试和集成测试
5. **添加文档**: 使用Swagger生成API文档 