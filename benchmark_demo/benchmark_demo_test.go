package main

import (
	"runtime"
	"testing"
	"time"
)

// TestDatabaseOperations 测试数据库操作
func TestDatabaseOperations(t *testing.T) {
	db := NewDatabase()

	// 测试写入
	db.Write("key1", "value1")
	db.Write("key2", "value2")

	// 测试读取
	if value := db.Read("key1"); value != "value1" {
		t.Errorf("期望值为 'value1'，实际为 '%s'", value)
	}

	if value := db.Read("key2"); value != "value2" {
		t.Errorf("期望值为 'value2'，实际为 '%s'", value)
	}

	// 测试不存在的键
	if value := db.Read("nonexistent"); value != "" {
		t.Errorf("期望空字符串，实际为 '%s'", value)
	}
}

// TestHTTPServer 测试 HTTP 服务器
func TestHTTPServer(t *testing.T) {
	server := NewHTTPServer()

	// 测试 POST 请求
	response := server.HandleRequest("POST", "test_key", "test_value")
	if response != "OK" {
		t.Errorf("期望响应为 'OK'，实际为 '%s'", response)
	}

	// 测试 GET 请求
	response = server.HandleRequest("GET", "test_key", "")
	if response != "test_value" {
		t.Errorf("期望响应为 'test_value'，实际为 '%s'", response)
	}

	// 测试无效请求
	response = server.HandleRequest("INVALID", "key", "value")
	if response != "Invalid request" {
		t.Errorf("期望响应为 'Invalid request'，实际为 '%s'", response)
	}
}

// TestBenchmarkerCreation 测试压测器创建
func TestBenchmarkerCreation(t *testing.T) {
	benchmarker := NewBenchmarker()

	if benchmarker.server == nil {
		t.Error("压测器的服务器不应该为 nil")
	}
}

// TestSingleThreadBenchmark 测试单线程压测
func TestSingleThreadBenchmark(t *testing.T) {
	benchmarker := NewBenchmarker()

	// 运行小规模的压测
	result := benchmarker.SingleThreadBenchmark(10, 2)

	// 验证基本结果
	if result.TotalRequests != 10 {
		t.Errorf("期望总请求数为 10，实际为 %d", result.TotalRequests)
	}

	if result.TotalTime <= 0 {
		t.Error("总时间应该大于 0")
	}

	if result.RequestsPerSecond <= 0 {
		t.Error("QPS 应该大于 0")
	}
}

// TestMemoryBenchmark 测试内存压测
func TestMemoryBenchmark(t *testing.T) {
	benchmarker := NewBenchmarker()

	// 运行小规模的内存压测
	initialGoroutines := runtime.NumGoroutine()
	benchmarker.MemoryBenchmark(100)
	finalGoroutines := runtime.NumGoroutine()

	// 验证 goroutine 数量没有异常增长
	if finalGoroutines > initialGoroutines+10 {
		t.Errorf("Goroutine 数量异常增长: 从 %d 到 %d", initialGoroutines, finalGoroutines)
	}
}

// TestConnectionBenchmark 测试连接压测
func TestConnectionBenchmark(t *testing.T) {
	benchmarker := NewBenchmarker()

	// 运行小规模的连接压测
	initialGoroutines := runtime.NumGoroutine()
	benchmarker.ConnectionBenchmark(10)
	finalGoroutines := runtime.NumGoroutine()

	// 验证 goroutine 数量合理
	if finalGoroutines > initialGoroutines+20 {
		t.Errorf("Goroutine 数量异常增长: 从 %d 到 %d", initialGoroutines, finalGoroutines)
	}
}

// TestBenchmarkResult 测试压测结果结构
func TestBenchmarkResult(t *testing.T) {
	result := BenchmarkResult{
		TotalRequests:      100,
		SuccessfulRequests: 95,
		FailedRequests:     5,
		TotalTime:          time.Second,
		AverageLatency:     time.Millisecond * 10,
		MinLatency:         time.Millisecond * 5,
		MaxLatency:         time.Millisecond * 20,
		RequestsPerSecond:  100.0,
	}

	// 验证结果一致性
	if result.TotalRequests != result.SuccessfulRequests+result.FailedRequests {
		t.Error("总请求数应该等于成功请求数加失败请求数")
	}

	if result.MinLatency > result.AverageLatency || result.AverageLatency > result.MaxLatency {
		t.Error("延迟时间应该满足: 最小 <= 平均 <= 最大")
	}

	if result.RequestsPerSecond != float64(result.TotalRequests)/result.TotalTime.Seconds() {
		t.Error("QPS 计算不正确")
	}
}

// BenchmarkDatabaseRead 基准测试数据库读取
func BenchmarkDatabaseRead(b *testing.B) {
	db := NewDatabase()
	db.Write("test_key", "test_value")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = db.Read("test_key")
	}
}

// BenchmarkDatabaseWrite 基准测试数据库写入
func BenchmarkDatabaseWrite(b *testing.B) {
	db := NewDatabase()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db.Write("test_key", "test_value")
	}
}

// BenchmarkHTTPServerHandle 基准测试 HTTP 服务器处理
func BenchmarkHTTPServerHandle(b *testing.B) {
	server := NewHTTPServer()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		server.HandleRequest("GET", "test_key", "")
	}
}

// BenchmarkSingleThreadBenchmark 基准测试单线程压测
func BenchmarkSingleThreadBenchmark(b *testing.B) {
	benchmarker := NewBenchmarker()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmarker.SingleThreadBenchmark(100, 5)
	}
}
