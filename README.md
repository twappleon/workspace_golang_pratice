# Go 语言实践项目

这是一个 Go 语言学习和实践的示例项目集合，包含了多个常见的 Go 编程概念和模式的演示。

## 📁 项目结构

```
workspace_golang_pratice/
├── README.md
├── workspace_golang_pratice.code-workspace
├── run_all_demos.sh
├── run_all_tests.sh
├── init_demo/
│   ├── .vscode/
│   │   ├── launch.json
│   │   ├── settings.json
│   │   └── tasks.json
│   ├── go.mod
│   ├── init_demo.go
│   └── init_demo_test.go
├── goroutine_leak_demo/
│   ├── go.mod
│   └── goroutine_leak_demo.go
├── pass_by_value_demo/
│   ├── go.mod
│   ├── pass_by_value_demo.go
│   └── pass_by_value_demo_test.go
├── anonymous_var_demo/
│   ├── go.mod
│   ├── anonymous_var_demo.go
│   └── anonymous_var_demo_test.go
├── wait_done_demo/
│   ├── go.mod
│   ├── wait_done_demo.go
│   └── wait_done_demo_test.go
├── channel_demo/
│   ├── go.mod
│   ├── channel_demo.go
│   └── channel_demo_test.go
├── sync_demo/
│   ├── go.mod
│   ├── sync_demo.go
│   └── sync_demo_test.go
├── benchmark_demo/
│   ├── go.mod
│   ├── benchmark_demo.go
│   └── benchmark_demo_test.go
├── benchmark_tools/
│   ├── go.mod
│   └── benchmark_tool.go
├── channel_no_close_demo/
│   ├── go.mod
│   ├── channel_no_close_demo.go
│   ├── channel_no_close_demo_test.go
│   └── README.md
└── concurrency_examples/
    ├── go.mod
    ├── README.md
    ├── run_all_examples.sh
    ├── channel_behavior/
    │   ├── go.mod
    │   ├── main.go
    │   └── channel_behavior_test.go
    ├── context_demo/
    │   ├── go.mod
    │   └── main.go
    ├── goroutine_leak/
    │   ├── go.mod
    │   └── main.go
    ├── memory_allocation/
    │   ├── go.mod
    │   └── main.go
    └── race_condition/
        ├── go.mod
        └── main.go
```

## 🚀 快速开始

### 环境要求

- Go 1.24.3 或更高版本
- Cursor 或 VS Code 编辑器
- Go 扩展（推荐安装）

### 安装和运行

1. **克隆项目**
   ```bash
   git clone https://github.com/twappleon/workspace_golang_pratice.git
   cd workspace_golang_pratice
   ```

2. **使用工作区文件打开项目**
   - 在 Cursor/VS Code 中打开 `workspace_golang_pratice.code-workspace`
   - 这会自动配置所有必要的设置

3. **运行任意示例**
   ```bash
   # 进入特定目录
   cd init_demo
   
   # 运行程序
   go run init_demo.go
   ```

4. **运行所有示例**
   ```bash
   # 使用提供的脚本运行所有演示
   ./run_all_demos.sh
   ```

## 📚 示例说明

### 1. init_demo - 初始化函数演示
**文件**: `init_demo/init_demo.go`

演示 Go 语言中 `init()` 函数的执行顺序和包级变量初始化。

**运行结果**:
```
变量初始化
init函数1执行
init函数2执行
main函数执行: Hello
```

**运行方法**:
```bash
cd init_demo
go run init_demo.go
```

### 2. goroutine_leak_demo - Goroutine 泄漏演示
**文件**: `goroutine_leak_demo/goroutine_leak_demo.go`

演示常见的 Goroutine 泄漏模式和如何避免。

**运行方法**:
```bash
cd goroutine_leak_demo
go run goroutine_leak_demo.go
```

### 3. pass_by_value_demo - 值传递演示
**文件**: `pass_by_value_demo/pass_by_value_demo.go`

演示 Go 语言中的值传递机制，包括基本类型和结构体的传递。

**运行方法**:
```bash
cd pass_by_value_demo
go run pass_by_value_demo.go
```

### 4. anonymous_var_demo - 匿名变量演示
**文件**: `anonymous_var_demo/anonymous_var_demo.go`

演示 Go 语言中匿名变量（`_`）的使用场景。

**运行方法**:
```bash
cd anonymous_var_demo
go run anonymous_var_demo.go
```

### 5. wait_done_demo - WaitGroup 演示
**文件**: `wait_done_demo/wait_done_demo.go`

演示如何使用 `sync.WaitGroup` 来等待多个 Goroutine 完成。

**运行方法**:
```bash
cd wait_done_demo
go run wait_done_demo.go
```

### 6. channel_demo - Channel 演示
**文件**: `channel_demo/channel_demo.go`

演示 Go 语言中 Channel 的使用，包括无缓冲和有缓冲 Channel。

**运行方法**:
```bash
cd channel_demo
go run channel_demo.go
```

### 7. sync_demo - 同步原语演示
**文件**: `sync_demo/sync_demo.go`

演示各种同步原语的使用，如 Mutex、RWMutex 等。

**运行方法**:
```bash
cd sync_demo
go run sync_demo.go
```

### 8. benchmark_demo - 压测演示
**文件**: `benchmark_demo/benchmark_demo.go`

演示 Go 语言中的性能压测，包括 QPS、并发、内存等测试。

**运行方法**:
```bash
cd benchmark_demo
go run benchmark_demo.go
```

### 9. benchmark_tools - 压测工具
**文件**: `benchmark_tools/benchmark_tool.go`

通用的 HTTP 压测工具，支持自定义参数进行性能测试。

**运行方法**:
```bash
cd benchmark_tools
go run benchmark_tool.go -url http://localhost:8080 -n 1000 -c 20
```

## 🛠️ 开发环境配置

### Cursor/VS Code 配置

项目已预配置了以下设置：

- **Go 语言服务器**: 启用 gopls
- **代码格式化**: 保存时自动格式化
- **导入整理**: 保存时自动整理导入
- **调试配置**: 支持调试和运行
- **任务配置**: 提供运行和构建任务

### 运行按钮配置

如果运行按钮没有出现，请按以下步骤操作：

1. **重新加载窗口**
   - 按 `Cmd+Shift+P` (Mac) 或 `Ctrl+Shift+P` (Windows/Linux)
   - 输入 "Developer: Reload Window"

2. **重启语言服务器**
   - 按 `Cmd+Shift+P`
   - 输入 "Go: Restart Language Server"

3. **手动运行**
   - 使用快捷键 `F5` 开始调试
   - 或使用 `Cmd+Shift+P` → "Tasks: Run Task" → "Run Go Program"

## 🔧 常用命令

### 运行程序
```bash
# 运行当前目录的程序
go run .

# 运行特定文件
go run filename.go

# 构建程序
go build

# 运行测试
go test
```

### 模块管理
```bash
# 初始化模块
go mod init module_name

# 整理依赖
go mod tidy

# 验证模块
go mod verify
```

### 工具安装
```bash
# 安装 gopls 语言服务器
go install golang.org/x/tools/gopls@latest

# 安装 goimports
go install golang.org/x/tools/cmd/goimports@latest

# 安装 golint
go install golang.org/x/lint/golint@latest
```

## 🧪 测试

项目包含了完整的测试套件，涵盖了所有主要功能模块。

### 运行测试

#### 运行所有测试
```bash
# 使用提供的脚本运行所有测试
./run_all_tests.sh
```

#### 运行特定模块的测试
```bash
# 进入特定目录
cd channel_demo

# 运行测试
go test -v

# 运行基准测试
go test -bench=. -benchmem

# 运行测试并显示覆盖率
go test -v -cover
```

#### 运行基准测试
```bash
# 运行所有基准测试
go test -bench=. -benchmem ./...

# 运行特定基准测试
go test -bench=BenchmarkChannelSendReceive -benchmem
```

### 测试覆盖

每个模块都包含以下类型的测试：

- **单元测试**: 测试各个函数和方法的正确性
- **集成测试**: 测试模块间的交互
- **基准测试**: 测试性能表现
- **并发测试**: 测试并发安全性

### 测试文件说明

| 模块 | 测试文件 | 测试内容 |
|------|----------|----------|
| `init_demo` | `init_demo_test.go` | init 函数执行顺序、包级变量初始化 |
| `channel_demo` | `channel_demo_test.go` | 无缓冲/有缓冲 channel、select 语句、channel 关闭 |
| `sync_demo` | `sync_demo_test.go` | Mutex、RWMutex、WaitGroup、Once、竞态条件 |
| `pass_by_value_demo` | `pass_by_value_demo_test.go` | 值传递、指针传递、切片传递 |
| `wait_done_demo` | `wait_done_demo_test.go` | WaitGroup 基本使用、并发、超时、panic 处理 |
| `anonymous_var_demo` | `anonymous_var_demo_test.go` | 匿名变量使用、类型断言、错误处理 |
| `benchmark_demo` | `benchmark_demo_test.go` | 数据库操作、HTTP 服务器、压测器 |
| `channel_no_close_demo` | `channel_no_close_demo_test.go` | channel 关闭、goroutine 泄漏检测、context 取消 |

### 测试最佳实践

1. **测试命名**: 使用 `Test` 前缀的函数名，描述测试的具体场景
2. **测试覆盖**: 确保测试覆盖正常路径、边界条件和错误情况
3. **基准测试**: 为性能关键的函数提供基准测试
4. **并发测试**: 使用 `go test -race` 检测竞态条件
5. **测试隔离**: 每个测试应该是独立的，不依赖其他测试的状态

### 测试命令示例

```bash
# 运行测试并生成覆盖率报告
go test -v -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# 运行竞态检测
go test -race ./...

# 运行测试并显示详细输出
go test -v -count=1 ./...

# 运行基准测试并保存结果
go test -bench=. -benchmem -benchtime=5s ./... > benchmark_results.txt
```

## 📖 学习资源

- [Go 官方文档](https://golang.org/doc/)
- [Go 语言规范](https://golang.org/ref/spec)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go 并发模式](https://golang.org/doc/codewalk/sharemem/)

## 🤝 贡献

欢迎提交 Issue 和 Pull Request 来改进这个项目。

### 贡献流程

1. Fork 本仓库
2. 创建你的特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交你的更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开一个 Pull Request

### 问题反馈

如果你发现了问题或有改进建议，请通过以下方式反馈：

- [提交 Issue](https://github.com/twappleon/workspace_golang_pratice/issues)
- [创建 Pull Request](https://github.com/twappleon/workspace_golang_pratice/pulls)

## 📄 许可证

本项目采用 MIT 许可证。详见 [LICENSE](LICENSE) 文件。

## 🔗 相关链接

- [项目地址](https://github.com/twappleon/workspace_golang_pratice)
- [Go 官网](https://golang.org/)
- [Go 包管理](https://golang.org/cmd/go/)
- [Go 工具链](https://golang.org/cmd/)

## 📊 项目状态

![GitHub last commit](https://img.shields.io/github/last-commit/twappleon/workspace_golang_pratice)
![GitHub issues](https://img.shields.io/github/issues/twappleon/workspace_golang_pratice)
![GitHub pull requests](https://img.shields.io/github/issues-pr/twappleon/workspace_golang_pratice)
![GitHub license](https://img.shields.io/github/license/twappleon/workspace_golang_pratice) 