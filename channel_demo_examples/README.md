# Channel Demo Examples

这个项目包含了Go语言中channel的各种使用示例。

## 文件说明

- `main.go` - 主程序，运行所有demo
- `basic_channel.go` - 基础channel操作演示
- `select_demo.go` - select语句使用演示
- `pipeline_demo.go` - channel pipeline模式演示
- `worker_pool.go` - worker pool模式演示

## 运行方式

```bash
cd channel_demo_examples
go run .
```

## Demo内容

### 1. 基础Channel Demo
- 无缓冲channel的创建和使用
- 有缓冲channel的创建和使用
- channel的关闭和检查

### 2. Select Demo
- 多channel的select操作
- 带timeout的select
- 非阻塞的select (带default)

### 3. Pipeline Demo
- 使用channel构建数据处理管道
- 数据流: generate -> square -> filterEven -> sqrt
- 展示如何链式处理数据

### 4. Worker Pool Demo
- 使用channel实现worker pool模式
- 并发处理多个任务
- 使用WaitGroup同步

## 学习要点

1. **Channel基础**: 创建、发送、接收、关闭
2. **缓冲vs无缓冲**: 理解阻塞行为
3. **Select语句**: 多channel处理
4. **Pipeline模式**: 数据流处理
5. **Worker Pool**: 并发任务处理
6. **同步机制**: WaitGroup的使用 