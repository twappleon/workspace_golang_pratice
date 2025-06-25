# Go语言学习与实践工作空间

[![Go Version](https://img.shields.io/badge/Go-1.24.3+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

这是一个Go语言学习和实践的工作空间，包含多个独立的示例项目，涵盖Go语言的核心概念、并发编程、内存管理等各个方面。

## 📁 项目结构

```
workspace_golang_pratice/
├── README.md                    # 工作空间总览
├── LICENSE                      # 许可证文件
├── .gitignore                   # Git忽略文件
├── workspace_golang_pratice.code-workspace  # VS Code工作空间配置
├── concurrency_examples/        # 并发编程示例集
│   ├── README.md                # 并发示例详细说明
│   ├── run_all_examples.sh      # 一键运行脚本
│   ├── race_condition/          # 竞态条件示例
│   ├── channel_behavior/        # Channel行为示例
│   ├── memory_allocation/       # 内存分配示例
│   ├── goroutine_leak/          # Goroutine泄漏示例
│   └── context_demo/            # Context使用示例
├── init_demo/                   # 包初始化示例
├── goroutine_leak_demo/         # Goroutine泄漏演示
├── pass_by_value_demo/          # 值传递示例
├── anonymous_var_demo/          # 匿名变量示例
├── wait_done_demo/              # Wait/Done模式示例
├── channel_demo/                # Channel基础示例
└── sync_demo/                   # 同步原语示例
```

## 🎯 学习内容

### 1. 并发编程实践 (`concurrency_examples/`)
- **竞态条件处理** - 并发安全问题及解决方案
- **Channel通信** - Go的CSP并发模型
- **内存管理** - 并发环境下的内存分配和GC
- **Goroutine管理** - 避免goroutine泄漏
- **Context模式** - 上下文控制和取消机制

### 2. 基础概念示例
- **包初始化** (`init_demo/`) - 包级变量和init函数
- **值传递** (`pass_by_value_demo/`) - Go的值传递机制
- **匿名变量** (`anonymous_var_demo/`) - 匿名变量的使用
- **Wait/Done模式** (`wait_done_demo/`) - 同步等待模式
- **Channel基础** (`channel_demo/`) - Channel的基本使用
- **同步原语** (`sync_demo/`) - 各种同步机制
- **Goroutine泄漏** (`goroutine_leak_demo/`) - 泄漏问题演示

## 🚀 快速开始

### 环境要求
- **Go版本**: 1.24.3+
- **操作系统**: Linux, macOS, Windows
- **IDE推荐**: VS Code with Go extension

### 运行示例

#### 1. 运行并发编程示例
```bash
cd concurrency_examples
./run_all_examples.sh
```

#### 2. 运行单个并发示例
```bash
cd concurrency_examples/race_condition
go run main.go
```

#### 3. 运行基础概念示例
```bash
# 包初始化示例
cd init_demo
go run init_demo.go

# 值传递示例
cd ../pass_by_value_demo
go run pass_by_value_demo.go

# Channel基础示例
cd ../channel_demo
go run channel_demo.go
```

#### 4. 检测竞态条件
```bash
cd concurrency_examples/race_condition
go run -race main.go
```

## 📚 学习路径

### 初学者路径
1. **基础概念** → 从 `init_demo` 开始理解包机制
2. **值传递** → 学习 `pass_by_value_demo` 理解传值机制
3. **Channel基础** → 通过 `channel_demo` 了解通信
4. **并发编程** → 深入学习 `concurrency_examples`
5. **同步机制** → 掌握 `sync_demo` 和 `wait_done_demo`

### 进阶路径
1. **并发安全** → 深入理解竞态条件和同步
2. **内存管理** → 学习内存分配和GC优化
3. **资源管理** → 掌握goroutine和context管理
4. **性能调优** → 通过实际示例优化性能

## 🔧 开发环境配置

### VS Code工作空间
本项目包含预配置的VS Code工作空间文件：
- `workspace_golang_pratice.code-workspace` - 包含所有项目的配置
- 自动配置Go语言服务器
- 代码格式化、导入整理等设置

### 推荐扩展
- [Go](https://marketplace.visualstudio.com/items?itemName=golang.Go) - 官方Go扩展
- [Go Test Explorer](https://marketplace.visualstudio.com/items?itemName=premparihar.gotestexplorer) - 测试探索器

## 📖 详细文档

### 并发编程示例
详细说明请查看 [concurrency_examples/README.md](concurrency_examples/README.md)

### 各示例说明
- 每个示例目录都包含独立的 `go.mod` 文件
- 示例代码包含详细的中文注释
- 运行结果展示实际的行为

## 🛠️ 工具和脚本

### 一键运行脚本
```bash
# 运行所有并发示例
./concurrency_examples/run_all_examples.sh
```

### 常用命令
```bash
# 格式化代码
go fmt ./...

# 检查代码
go vet ./...

# 运行测试
go test ./...

# 构建项目
go build ./...
```

## 🤝 贡献指南

欢迎贡献代码和文档！

### 贡献方式
1. Fork 项目
2. 创建特性分支
3. 提交更改
4. 推送到分支
5. 创建 Pull Request

### 代码规范
- 遵循Go官方代码规范
- 添加适当的中文注释
- 包含测试用例
- 更新相关文档

## 📖 相关资源

### 官方文档
- [Go官方文档](https://golang.org/doc/)
- [Go并发编程指南](https://golang.org/doc/effective_go.html#concurrency)
- [Go内存模型](https://golang.org/ref/mem)

### 推荐书籍
- 《Go程序设计语言》
- 《Go并发编程实战》
- 《Go语言实战》

### 在线资源
- [Go by Example](https://gobyexample.com/)
- [Go语言中文网](https://studygolang.com/)
- [Go Web Examples](https://gowebexamples.com/)

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 🙏 致谢

感谢Go语言社区提供的优秀编程模型和工具。

---

**Happy Learning! 🚀**

如有问题或建议，请提交 [Issue](https://github.com/your-repo/issues) 或联系维护者。

# Init 函数演示

这个示例演示了 Go 语言中 `init()` 函数的执行顺序和包级变量初始化机制。

## 📝 代码说明

### 包级变量初始化
```go
var msg = func() string {
    fmt.Println("变量初始化")
    return "Hello"
}()
```

这个变量在包被导入时就会执行初始化函数，输出 "变量初始化"。

### Init 函数
```go
func init() {
    fmt.Println("init函数1执行")
}

func init() {
    fmt.Println("init函数2执行")
}
```

Go 语言允许在一个包中定义多个 `init()` 函数，它们会按照定义的顺序执行。

### Main 函数
```go
func main() {
    fmt.Println("main函数执行:", msg)
}
```

`main()` 函数在所有 `init()` 函数执行完毕后才会执行。

## 🚀 运行方法

```bash
# 进入目录
cd init_demo

# 运行程序
go run init_demo.go
```

## 📊 运行结果

```
变量初始化
init函数1执行
init函数2执行
main函数执行: Hello
```

## 🔍 执行顺序

1. **包级变量初始化** - 首先执行包级变量的初始化
2. **Init 函数执行** - 按照定义顺序执行所有 `init()` 函数
3. **Main 函数执行** - 最后执行 `main()` 函数

## 💡 学习要点

- `init()` 函数在包被导入时自动执行
- 一个包可以有多个 `init()` 函数
- `init()` 函数不能手动调用
- 包级变量在 `init()` 函数之前初始化
- `main()` 函数在所有 `init()` 函数之后执行

## 🔗 相关概念

- [Go 包初始化](https://golang.org/ref/spec#Package_initialization)
- [Init 函数](https://golang.org/doc/effective_go.html#init)
- [包级变量](https://golang.org/ref/spec#Variable_declarations) 