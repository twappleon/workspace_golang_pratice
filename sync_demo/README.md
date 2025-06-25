# 同步原语演示

这个示例演示了 Go 语言中各种同步原语的使用，包括 Mutex、RWMutex、WaitGroup 等。

## 📝 代码说明

### Mutex (互斥锁)
```go
var mu sync.Mutex
mu.Lock()
// 临界区代码
mu.Unlock()
```

Mutex 提供互斥访问，确保同一时间只有一个 Goroutine 能访问共享资源。

### RWMutex (读写锁)
```go
var rwmu sync.RWMutex
rwmu.RLock() // 读锁
// 读取操作
rwmu.RUnlock()

rwmu.Lock() // 写锁
// 写入操作
rwmu.Unlock()
```

RWMutex 允许多个读操作同时进行，但写操作是互斥的。

### WaitGroup
```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // 工作代码
}()
wg.Wait()
```

WaitGroup 用于等待多个 Goroutine 完成。

## 🚀 运行方法

```bash
# 进入目录
cd sync_demo

# 运行程序
go run sync_demo.go
```

## 📊 运行结果

程序会演示各种同步原语的使用场景，包括：
- 使用 Mutex 保护共享数据
- 使用 RWMutex 优化读写性能
- 使用 WaitGroup 等待 Goroutine 完成
- 避免竞态条件

## 💡 学习要点

- 同步原语用于协调多个 Goroutine 的访问
- Mutex 提供互斥访问保护
- RWMutex 在读多写少的场景下性能更好
- WaitGroup 用于等待多个 Goroutine 完成
- 正确使用同步原语可以避免竞态条件

## 🔗 相关概念

- [sync 包](https://golang.org/pkg/sync/)
- [Mutex](https://golang.org/pkg/sync/#Mutex)
- [RWMutex](https://golang.org/pkg/sync/#RWMutex)
- [WaitGroup](https://golang.org/pkg/sync/#WaitGroup) 