# Channel基础示例

这个示例演示了Go语言中Channel的基本使用方法，包括创建、发送、接收、关闭等操作。

## 📁 文件结构

```
channel_demo/
├── README.md        # 项目说明
├── go.mod           # 模块定义
└── channel_demo.go  # 主程序文件
```

## 🎯 学习目标

- 理解Channel的基本概念
- 掌握Channel的创建和使用
- 了解无缓冲和有缓冲Channel的区别
- 学会Channel的关闭和检测

## 📖 示例内容

### 创建Channel
```go
// 无缓冲channel
ch1 := make(chan int)

// 有缓冲channel
ch2 := make(chan int, 3)
```

### 发送和接收
```go
// 发送数据
ch <- 42

// 接收数据
value := <-ch
```

### Channel关闭
```go
// 关闭channel
close(ch)

// 检测channel是否关闭
value, ok := <-ch
if !ok {
    fmt.Println("Channel已关闭")
}
```

## 🚀 运行示例

```bash
go run channel_demo.go
```

## 📋 预期输出

```
=== 无缓冲Channel演示 ===
发送方准备发送...
接收方准备接收...
接收到数据: 42

=== 有缓冲Channel演示 ===
发送数据: 1, 2, 3
接收数据: 1
接收数据: 2
接收数据: 3

=== Channel关闭演示 ===
发送数据: 100, 200
Channel已关闭
```

## 💡 关键概念

1. **无缓冲Channel**：发送和接收必须同时准备好
2. **有缓冲Channel**：可以缓存指定数量的数据
3. **Channel关闭**：只能关闭一次，关闭后不能再发送
4. **Channel检测**：使用ok值检测channel是否关闭

## 🔍 学习要点

- Channel是Go语言并发编程的核心
- 无缓冲Channel提供同步通信
- 有缓冲Channel提供异步通信
- 发送方负责关闭Channel
- 关闭的Channel仍然可以接收剩余数据 