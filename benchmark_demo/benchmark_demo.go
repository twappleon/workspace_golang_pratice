package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 模拟数据库操作
type Database struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewDatabase() *Database {
	return &Database{
		data: make(map[string]string),
	}
}

func (db *Database) Read(key string) string {
	db.mu.RLock()
	defer db.mu.RUnlock()
	return db.data[key]
}

func (db *Database) Write(key, value string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.data[key] = value
}

// 模拟HTTP服务器
type HTTPServer struct {
	db *Database
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{
		db: NewDatabase(),
	}
}

func (s *HTTPServer) HandleRequest(reqType string, key, value string) string {
	switch reqType {
	case "GET":
		return s.db.Read(key)
	case "POST":
		s.db.Write(key, value)
		return "OK"
	default:
		return "Invalid request"
	}
}

// 压测结果统计
type BenchmarkResult struct {
	TotalRequests      int64
	SuccessfulRequests int64
	FailedRequests     int64
	TotalTime          time.Duration
	AverageLatency     time.Duration
	MinLatency         time.Duration
	MaxLatency         time.Duration
	RequestsPerSecond  float64
}

// 压测器
type Benchmarker struct {
	server *HTTPServer
}

func NewBenchmarker() *Benchmarker {
	return &Benchmarker{
		server: NewHTTPServer(),
	}
}

// 单线程压测
func (b *Benchmarker) SingleThreadBenchmark(requests int, concurrency int) BenchmarkResult {
	var wg sync.WaitGroup
	var mu sync.Mutex

	result := BenchmarkResult{}
	latencies := make([]time.Duration, 0, requests)

	// 预热
	for i := 0; i < 100; i++ {
		b.server.HandleRequest("GET", "key", "")
	}

	startTime := time.Now()

	// 创建工作协程
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < requests/concurrency; j++ {
				reqStart := time.Now()

				// 模拟不同类型的请求
				reqType := "GET"
				if j%3 == 0 {
					reqType = "POST"
				}

				key := fmt.Sprintf("key_%d", j)
				value := fmt.Sprintf("value_%d", j)

				response := b.server.HandleRequest(reqType, key, value)

				latency := time.Since(reqStart)

				mu.Lock()
				result.TotalRequests++
				if response != "" {
					result.SuccessfulRequests++
				} else {
					result.FailedRequests++
				}
				latencies = append(latencies, latency)
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	result.TotalTime = time.Since(startTime)

	// 计算统计信息
	if len(latencies) > 0 {
		result.MinLatency = latencies[0]
		result.MaxLatency = latencies[0]
		var totalLatency time.Duration

		for _, latency := range latencies {
			totalLatency += latency
			if latency < result.MinLatency {
				result.MinLatency = latency
			}
			if latency > result.MaxLatency {
				result.MaxLatency = latency
			}
		}

		result.AverageLatency = totalLatency / time.Duration(len(latencies))
	}

	result.RequestsPerSecond = float64(result.TotalRequests) / result.TotalTime.Seconds()

	return result
}

// 内存压测
func (b *Benchmarker) MemoryBenchmark(iterations int) {
	fmt.Println("开始内存压测...")

	var mu sync.Mutex
	memoryUsage := make([]int, 0, iterations)

	for i := 0; i < iterations; i++ {
		// 模拟内存分配
		data := make([]byte, 1024*1024) // 1MB
		rand.Read(data)

		mu.Lock()
		memoryUsage = append(memoryUsage, len(data))
		mu.Unlock()

		if i%1000 == 0 {
			fmt.Printf("内存压测进度: %d/%d\n", i, iterations)
		}
	}

	fmt.Printf("内存压测完成，分配了 %d 个 1MB 块\n", len(memoryUsage))
}

// 并发连接压测
func (b *Benchmarker) ConnectionBenchmark(maxConnections int) {
	fmt.Println("开始并发连接压测...")

	var wg sync.WaitGroup
	activeConnections := make(chan struct{}, maxConnections)

	for i := 0; i < maxConnections*2; i++ {
		wg.Add(1)
		go func(connID int) {
			defer wg.Done()

			// 获取连接
			activeConnections <- struct{}{}
			defer func() { <-activeConnections }()

			// 模拟连接处理
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))

			if connID%100 == 0 {
				fmt.Printf("连接 %d 处理完成\n", connID)
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("并发连接压测完成，最大并发连接数: %d\n", maxConnections)
}

// 打印压测结果
func (b *Benchmarker) PrintResult(result BenchmarkResult, testName string) {
	fmt.Printf("\n=== %s 压测结果 ===\n", testName)
	fmt.Printf("总请求数: %d\n", result.TotalRequests)
	fmt.Printf("成功请求数: %d\n", result.SuccessfulRequests)
	fmt.Printf("失败请求数: %d\n", result.FailedRequests)
	fmt.Printf("总耗时: %v\n", result.TotalTime)
	fmt.Printf("平均延迟: %v\n", result.AverageLatency)
	fmt.Printf("最小延迟: %v\n", result.MinLatency)
	fmt.Printf("最大延迟: %v\n", result.MaxLatency)
	fmt.Printf("QPS: %.2f\n", result.RequestsPerSecond)
	fmt.Printf("成功率: %.2f%%\n", float64(result.SuccessfulRequests)/float64(result.TotalRequests)*100)
}

func main() {
	fmt.Println("🚀 Go 语言压测演示")
	fmt.Println("==================")

	benchmarker := NewBenchmarker()

	// 1. 基础压测
	fmt.Println("\n1. 基础压测 (1000 请求，10 并发)")
	result1 := benchmarker.SingleThreadBenchmark(1000, 10)
	benchmarker.PrintResult(result1, "基础压测")

	// 2. 高并发压测
	fmt.Println("\n2. 高并发压测 (5000 请求，50 并发)")
	result2 := benchmarker.SingleThreadBenchmark(5000, 50)
	benchmarker.PrintResult(result2, "高并发压测")

	// 3. 内存压测
	fmt.Println("\n3. 内存压测")
	benchmarker.MemoryBenchmark(10000)

	// 4. 连接压测
	fmt.Println("\n4. 并发连接压测")
	benchmarker.ConnectionBenchmark(100)

	fmt.Println("\n✅ 所有压测完成！")
}
