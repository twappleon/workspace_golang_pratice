# HTTP 压测工具

这是一个通用的 HTTP 压测工具，用于测试 Web 服务的性能表现。

## 📝 功能特性

- **多并发支持**: 支持自定义并发数进行压测
- **灵活参数**: 可配置请求数、超时时间、HTTP 方法等
- **详细统计**: 提供 QPS、延迟百分位数、成功率等指标
- **简单易用**: 命令行工具，使用方便

## 🚀 使用方法

### 基本用法

```bash
cd benchmark_tools
go run benchmark_tool.go -url http://localhost:8080
```

### 高级用法

```bash
# 自定义请求数和并发数
go run benchmark_tool.go -url http://localhost:8080 -n 5000 -c 50

# 指定 HTTP 方法和超时时间
go run benchmark_tool.go -url http://localhost:8080 -method POST -timeout 10s

# 完整参数示例
go run benchmark_tool.go \
  -url http://api.example.com/users \
  -n 10000 \
  -c 100 \
  -method GET \
  -timeout 30s
```

## 📊 参数说明

| 参数 | 默认值 | 说明 |
|------|--------|------|
| `-url` | `http://localhost:8080` | 目标 HTTP 地址 |
| `-n` | `1000` | 总请求数 |
| `-c` | `10` | 并发数 |
| `-method` | `GET` | HTTP 方法 |
| `-timeout` | `30s` | 请求超时时间 |

## 📈 输出结果

工具会输出详细的压测统计信息：

```
🚀 HTTP 压测工具
================
目标URL: http://localhost:8080
总请求数: 1000
并发数: 20
HTTP方法: GET
超时时间: 30s

开始压测...
=== 压测结果 ===
总请求数: 1000
成功请求数: 1000
失败请求数: 0
总耗时: 2.5s
平均延迟: 50ms
最小延迟: 10ms
最大延迟: 200ms
QPS: 400.00
成功率: 100.00%

延迟百分位数:
  50%: 45ms
  90%: 80ms
  95%: 120ms
  99%: 180ms
```

## 💡 使用建议

### 1. 压测前准备
- 确保目标服务正常运行
- 选择合适的并发数和请求数
- 监控目标服务的资源使用情况

### 2. 参数调优
- **并发数**: 从较小的值开始，逐步增加
- **请求数**: 根据测试需求设置，建议至少 1000 次
- **超时时间**: 根据服务响应时间设置

### 3. 结果分析
- **QPS**: 每秒请求数，越高越好
- **延迟**: 响应时间，越低越好
- **成功率**: 应该接近 100%
- **百分位数**: 了解延迟分布情况

## 🔧 构建和安装

### 构建可执行文件
```bash
cd benchmark_tools
go build -o benchmark-tool benchmark_tool.go
```

### 安装到系统
```bash
go install .
```

## 📖 示例场景

### 1. 测试 API 性能
```bash
go run benchmark_tool.go -url http://api.example.com/users -n 5000 -c 50
```

### 2. 测试 POST 接口
```bash
go run benchmark_tool.go -url http://api.example.com/users -method POST -n 1000 -c 20
```

### 3. 高并发测试
```bash
go run benchmark_tool.go -url http://api.example.com/users -n 10000 -c 200
```

## 🔗 相关工具

- [Apache Bench (ab)](https://httpd.apache.org/docs/2.4/programs/ab.html)
- [wrk](https://github.com/wg/wrk)
- [hey](https://github.com/rakyll/hey)

## 📄 许可证

本项目采用 MIT 许可证。 