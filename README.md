# Go 语言实践项目

这是一个 Go 语言学习和实践的示例项目集合，包含了多个常见的 Go 编程概念和模式的演示。

## 📋 目录

- [项目结构](#-项目结构)
- [快速开始](#-快速开始)
- [示例说明](#-示例说明)
- [并发编程](#-并发编程)
- [性能测试](#-性能测试)
- [设计模式](#-设计模式)
- [最佳实践](#-最佳实践)
- [贡献指南](#-贡献指南)

## 📁 项目结构

```
workspace_golang_pratice/
├── README.md                           # 项目说明文档
├── run_all_demos.sh                    # 运行所有演示的脚本
├── run_all_tests.sh                    # 运行所有测试的脚本
├── workspace_golang_pratice.code-workspace  # VS Code 工作区配置
├── LICENSE                             # 许可证文件
│
├── 基础概念演示/
│   ├── init_demo/                      # init函数演示
│   ├── pass_by_value_demo/             # 值传递演示
│   ├── anonymous_var_demo/             # 匿名变量演示
│   └── test_defer/                     # defer语句演示
│
├── 并发编程演示/
│   ├── channel_demo/                   # 基础channel演示
│   ├── channel_no_close_demo/          # channel不关闭问题演示
│   ├── channel_demo_examples/          # 完整channel示例集合
│   ├── sync_demo/                      # sync包使用演示
│   ├── wait_done_demo/                 # WaitGroup演示
│   ├── goroutine_leak_demo/            # goroutine泄漏演示
│   └── concurrency_examples/           # 并发编程综合示例
│
├── 性能测试演示/
│   ├── benchmark_demo/                 # 基准测试演示
│   └── benchmark_tools/                # 性能测试工具
│
└── 设计模式演示/
    └── oop_demo/                       # 面向对象和设计模式
```

## 🚀 快速开始

### 环境要求

- **Go 版本**: 1.21 或更高版本
- **编辑器**: Cursor 或 VS Code（推荐）
- **扩展**: Go 扩展（必需）

### 安装和运行

1. **克隆项目**
   ```bash
   git clone https://github.com/twappleon/workspace_golang_pratice.git
   cd workspace_golang_pratice
   ```

2. **使用工作区文件**
   - 在 Cursor/VS Code 中打开 `workspace_golang_pratice.code-workspace`
   - 自动配置所有必要的设置和扩展

3. **运行单个示例**
   ```bash
   # 进入特定目录
   cd init_demo
   
   # 运行程序
   go run .
   
   # 运行测试
   go test -v
   ```

4. **运行所有示例**
   ```bash
   # 运行所有演示
   ./run_all_demos.sh
   
   # 运行所有测试
   ./run_all_tests.sh
   ```

## 📚 示例详解

### 🔧 基础概念演示

#### 1. init_demo - 初始化函数
**位置**: `init_demo/`
**内容**: 演示 `init()` 函数的执行顺序和包级变量初始化
**运行**: `cd init_demo && go run .`

#### 2. pass_by_value_demo - 值传递
**位置**: `pass_by_value_demo/`
**内容**: 演示 Go 语言中的值传递机制
**运行**: `cd pass_by_value_demo && go run .`

#### 3. anonymous_var_demo - 匿名变量
**位置**: `anonymous_var_demo/`
**内容**: 演示匿名变量的使用场景
**运行**: `cd anonymous_var_demo && go run .`

### 🔄 并发编程演示

#### 1. channel_demo - 基础Channel
**位置**: `channel_demo/`
**内容**: Channel 的基本使用和操作
**运行**: `cd channel_demo && go run .`

#### 2. channel_demo_examples - 完整Channel示例
**位置**: `channel_demo_examples/`
**内容**: 包含以下完整示例：
- 基础Channel操作
- Select语句使用
- Pipeline模式
- Worker Pool模式
**运行**: `cd channel_demo_examples && go run .`

#### 3. channel_no_close_demo - Channel关闭问题
**位置**: `channel_no_close_demo/`
**内容**: 演示Channel不关闭导致的问题和解决方案
**运行**: `cd channel_no_close_demo && go run .`

#### 4. sync_demo - 同步原语
**位置**: `sync_demo/`
**内容**: sync包的使用（Mutex、RWMutex、Once等）
**运行**: `cd sync_demo && go run .`

#### 5. wait_done_demo - WaitGroup
**位置**: `wait_done_demo/`
**内容**: WaitGroup的使用和最佳实践
**运行**: `cd wait_done_demo && go run .`

#### 6. goroutine_leak_demo - Goroutine泄漏
**位置**: `goroutine_leak_demo/`
**内容**: 演示Goroutine泄漏的常见原因和预防
**运行**: `cd goroutine_leak_demo && go run .`

#### 7. concurrency_examples - 并发综合示例
**位置**: `concurrency_examples/`
**内容**: 包含以下示例：
- Channel行为分析
- Context使用
- Goroutine泄漏检测
- 内存分配分析
- 竞态条件检测
**运行**: `cd concurrency_examples && ./run_all_examples.sh`

### ⚡ 性能测试演示

#### 1. benchmark_demo - 基准测试
**位置**: `benchmark_demo/`
**内容**: Go基准测试的使用方法
**运行**: `cd benchmark_demo && go test -bench=.`

#### 2. benchmark_tools - 性能工具
**位置**: `benchmark_tools/`
**内容**: 性能测试和分析工具
**运行**: `cd benchmark_tools && go run .`

### 🎨 设计模式演示

#### 1. oop_demo - 面向对象和设计模式
**位置**: `oop_demo/`
**内容**: Go语言中的面向对象编程和设计模式实现
**运行**: `cd oop_demo && go run .`

## 🛠️ 开发工具

### 推荐的VS Code扩展

1. **Go** - 官方Go扩展
2. **Go Test Explorer** - 测试资源管理器
3. **Go Outliner** - 代码大纲
4. **Go Doc** - 文档查看器
5. **Go Lint** - 代码检查

### 调试配置

项目包含完整的VS Code调试配置：
- 支持断点调试
- 支持测试调试
- 支持基准测试调试

## 📖 学习路径

### 初学者路径
1. `init_demo` - 了解Go程序初始化
2. `pass_by_value_demo` - 理解值传递
3. `channel_demo` - 学习基础并发
4. `sync_demo` - 掌握同步机制

### 进阶路径
1. `channel_demo_examples` - 深入Channel使用
2. `concurrency_examples` - 并发编程实践
3. `benchmark_demo` - 性能测试
4. `oop_demo` - 设计模式

### 专家路径
1. `goroutine_leak_demo` - 内存泄漏分析
2. `channel_no_close_demo` - 并发问题诊断
3. `benchmark_tools` - 性能分析工具
4. 自定义示例开发

## 🚨 注意事项

### 并发安全
- 所有并发示例都经过测试
- 注意goroutine泄漏问题
- 使用适当的同步机制

### 性能考虑
- 基准测试结果仅供参考
- 实际性能可能因环境而异
- 建议在生产环境中进行充分测试

### 最佳实践
1. **Channel使用**: 发送者负责关闭channel
2. **Goroutine管理**: 确保所有goroutine都能正常退出
3. **错误处理**: 正确处理并发错误
4. **资源管理**: 及时释放资源

## 🤝 贡献指南

1. **Fork 项目**
2. **创建特性分支**: `git checkout -b feature/AmazingFeature`
3. **提交更改**: `git commit -m 'Add some AmazingFeature'`
4. **推送到分支**: `git push origin feature/AmazingFeature`
5. **打开 Pull Request**

### 贡献规范
- 添加适当的测试
- 更新相关文档
- 遵循Go代码规范
- 添加有意义的提交信息

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 🔗 相关资源

### 官方文档
- [Go 官方文档](https://golang.org/doc/)
- [Go 并发编程指南](https://golang.org/doc/effective_go.html#concurrency)
- [Go 测试文档](https://golang.org/pkg/testing/)

### 学习资源
- [Go by Example](https://gobyexample.com/)
- [Go Web Examples](https://gowebexamples.com/)
- [Go Concurrency Patterns](https://blog.golang.org/pipelines)

### 工具推荐
- [GoLand](https://www.jetbrains.com/go/) - 专业Go IDE
- [Delve](https://github.com/go-delve/delve) - Go调试器
- [pprof](https://golang.org/pkg/runtime/pprof/) - 性能分析

---

**注意**: 这些示例主要用于学习和演示目的。在生产环境中使用时，请根据具体需求进行调整和充分测试。

**Star ⭐**: 如果这个项目对你有帮助，请给个Star支持一下！ 