# Go并发编程实践示例集

[![Go Version](https://img.shields.io/badge/Go-1.24.3+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

这是一个全面的Go语言并发编程实践示例集合，旨在帮助开发者深入理解Go并发编程的核心概念、常见问题和最佳实践。

## 📋 目录

- [项目概述](#项目概述)
- [示例结构](#示例结构)
- [快速开始](#快速开始)
- [详细示例](#详细示例)
- [学习路径](#学习路径)
- [最佳实践](#最佳实践)
- [故障排除](#故障排除)
- [贡献指南](#贡献指南)

## 🎯 项目概述

本项目包含5个核心示例模块，涵盖了Go并发编程的各个方面：

- **竞态条件处理** - 理解并发安全问题及解决方案
- **Channel通信** - 掌握Go的CSP并发模型
- **内存管理** - 学习并发环境下的内存分配和GC
- **Goroutine管理** - 避免常见的goroutine泄漏问题
- **Context模式** - 掌握上下文控制和取消机制

## 📁 示例结构

```
concurrency_examples/
├── README.md                 # 项目说明文档
├── run_all_examples.sh       # 一键运行所有示例的脚本
├── go.mod                    # 根模块定义
├── race_condition/           # 竞态条件示例
│   ├── go.mod
│   └── main.go
├── channel_behavior/         # Channel行为示例
│   ├── go.mod
│   └── main.go
├── memory_allocation/        # 内存分配示例
│   ├── go.mod
│   └── main.go
├── goroutine_leak/           # Goroutine泄漏示例
│   ├── go.mod
│   └── main.go
└── context_demo/             # Context使用示例
    ├── go.mod
    └── main.go
```

## 🚀 快速开始

### 环境要求

- Go 1.24.3 或更高版本
- 支持的操作系统：Linux, macOS, Windows

### 安装和运行

1. **克隆或下载项目**
   ```bash
   git clone <repository-url>
   cd concurrency_examples
   ```

2. **运行单个示例**
   ```bash
   # 运行竞态条件示例
   cd race_condition
   go run main.go
   
   # 运行Channel行为示例
   cd ../channel_behavior
   go run main.go
   ```

3. **一键运行所有示例**
   ```bash
   # 在项目根目录执行
   ./run_all_examples.sh
   ```

4. **检测竞态条件**
   ```bash
   cd race_condition
   go run -race main.go
   ```

## 📚 详细示例

### 1. 竞态条件示例 (`race_condition/`)

**学习目标**：理解并发安全问题及解决方案

**核心概念**：
- 竞态条件（Race Condition）
- 互斥锁（Mutex）
- 原子操作（Atomic Operations）
- 读写锁（RWMutex）

**示例内容**：
```go
// 演示竞态条件问题
func raceConditionDemo() {
    var counter int
    // 多个goroutine同时修改counter
    // 结果通常小于预期值
}

// 使用互斥锁解决
func mutexSolution() {
    var mu sync.Mutex
    // 保证原子性操作
}

// 使用原子操作解决
func atomicSolution() {
    var counter int64
    atomic.AddInt64(&counter, 1)
}
```

**运行结果示例**：
```
=== 竞态条件演示 ===
预期结果: 10000, 实际结果: 8110
注意：由于竞态条件，结果通常小于10000

=== 使用互斥锁解决竞态条件 ===
预期结果: 10000, 实际结果: 10000
```

### 2. Channel行为示例 (`channel_behavior/`)

**学习目标**：掌握Go的CSP并发模型

**核心概念**：
- 无缓冲Channel
- 有缓冲Channel
- Channel阻塞行为
- Select语句
- Channel关闭和遍历

**示例内容**：
```go
// 无缓冲channel - 同步通信
ch := make(chan int)
go func() { ch <- 42 }()
value := <-ch

// 有缓冲channel - 异步通信
ch := make(chan int, 2)
ch <- 1
ch <- 2

// Select语句 - 非阻塞操作
select {
case msg1 := <-ch1:
    // 处理消息1
case msg2 := <-ch2:
    // 处理消息2
case <-time.After(timeout):
    // 超时处理
}
```

### 3. 内存分配示例 (`memory_allocation/`)

**学习目标**：理解并发环境下的内存管理

**核心概念**：
- 内存分配统计
- 对象池模式
- 内存泄漏检测
- GC调优
- 并发内存分配

**示例内容**：
```go
// 内存使用统计
var m runtime.MemStats
runtime.ReadMemStats(&m)
fmt.Printf("内存使用: %d MB\n", m.Alloc/1024/1024)

// 对象池实现
pool := make(chan *ExpensiveObject, 10)
// 从池中获取对象，减少分配开销
```

### 4. Goroutine泄漏示例 (`goroutine_leak/`)

**学习目标**：避免常见的goroutine管理问题

**核心概念**：
- Goroutine泄漏检测
- Context取消机制
- Channel泄漏
- 资源清理

**示例内容**：
```go
// 检测goroutine数量
initialGoroutines := runtime.NumGoroutine()
// 启动goroutine
// 检查泄漏
currentGoroutines := runtime.NumGoroutine()
fmt.Printf("泄漏数量: %d\n", currentGoroutines-initialGoroutines)

// 使用context控制
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
```

### 5. Context示例 (`context_demo/`)

**学习目标**：掌握上下文控制和取消机制

**核心概念**：
- Context基础使用
- Context传递
- Context值传递
- 超时和死线控制

**示例内容**：
```go
// 超时控制
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

// 值传递
ctx = context.WithValue(ctx, "user_id", "12345")

// 取消控制
ctx, cancel := context.WithCancel(context.Background())
cancel() // 发送取消信号
```

## 🛤️ 学习路径

### 初学者路径
1. **Channel行为** → 理解Go的并发通信模型
2. **Context示例** → 学习上下文控制
3. **竞态条件** → 认识并发安全问题
4. **Goroutine泄漏** → 避免资源泄漏
5. **内存分配** → 优化性能

### 进阶路径
1. **竞态条件** → 深入理解并发安全
2. **内存分配** → 性能调优
3. **Goroutine泄漏** → 系统稳定性
4. **Channel行为** → 复杂并发模式
5. **Context示例** → 大规模系统设计

## 💡 最佳实践

### 并发安全
- ✅ 使用互斥锁保护共享资源
- ✅ 优先使用原子操作
- ✅ 使用读写锁提高读性能
- ❌ 避免直接访问共享变量

### Channel使用
- ✅ 发送方负责关闭channel
- ✅ 使用select避免阻塞
- ✅ 检查channel是否已关闭
- ❌ 避免向已关闭的channel发送数据

### 资源管理
- ✅ 使用context控制goroutine生命周期
- ✅ 及时清理资源
- ✅ 监控goroutine数量
- ❌ 避免goroutine泄漏

### 性能优化
- ✅ 使用对象池减少分配
- ✅ 合理设置channel缓冲区
- ✅ 监控内存使用
- ❌ 避免频繁的内存分配

## 🔧 故障排除

### 常见问题

1. **竞态条件检测**
   ```bash
   go run -race main.go
   ```

2. **内存泄漏检测**
   ```bash
   go tool pprof -http=:8080 mem.prof
   ```

3. **Goroutine泄漏检测**
   ```go
   fmt.Printf("Goroutine数量: %d\n", runtime.NumGoroutine())
   ```

4. **性能分析**
   ```bash
   go test -bench=. -benchmem
   ```

### 调试技巧

- 使用`fmt.Printf`添加调试信息
- 使用`runtime/debug`包获取堆栈信息
- 使用`pprof`进行性能分析
- 使用`go vet`检查代码问题

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
- 添加适当的注释
- 包含测试用例
- 更新相关文档

## 📖 相关资源

### 官方文档
- [Go并发编程指南](https://golang.org/doc/effective_go.html#concurrency)
- [Go内存模型](https://golang.org/ref/mem)
- [Go并发模式](https://blog.golang.org/pipelines)

### 推荐书籍
- 《Go并发编程实战》
- 《Go语言实战》
- 《Go程序设计语言》

### 在线资源
- [Go by Example](https://gobyexample.com/)
- [Go Web Examples](https://gowebexamples.com/)
- [Go语言中文网](https://studygolang.com/)

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 🙏 致谢

感谢Go语言社区提供的优秀并发编程模型和工具。

---

**Happy Coding! 🚀**

如有问题或建议，请提交 [Issue](https://github.com/your-repo/issues) 或联系维护者。 