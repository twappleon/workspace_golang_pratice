package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// 压测配置
type BenchmarkConfig struct {
	URL         string
	Requests    int
	Concurrency int
	Duration    time.Duration
	Method      string
	Headers     map[string]string
	Timeout     time.Duration
}

// 压测结果
type BenchmarkResult struct {
	TotalRequests      int64
	SuccessfulRequests int64
	FailedRequests     int64
	TotalTime          time.Duration
	AverageLatency     time.Duration
	MinLatency         time.Duration
	MaxLatency         time.Duration
	RequestsPerSecond  float64
	LatencyPercentiles map[string]time.Duration
}

// HTTP 压测器
type HTTPBenchmarker struct {
	client *http.Client
}

func NewHTTPBenchmarker(timeout time.Duration) *HTTPBenchmarker {
	return &HTTPBenchmarker{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

// 执行压测
func (b *HTTPBenchmarker) RunBenchmark(config BenchmarkConfig) BenchmarkResult {
	var wg sync.WaitGroup
	var mu sync.Mutex

	result := BenchmarkResult{}
	latencies := make([]time.Duration, 0, config.Requests)

	// 创建请求
	req, err := http.NewRequest(config.Method, config.URL, nil)
	if err != nil {
		log.Fatal("创建请求失败:", err)
	}

	// 添加请求头
	for key, value := range config.Headers {
		req.Header.Set(key, value)
	}

	startTime := time.Now()

	// 创建工作协程
	for i := 0; i < config.Concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			requestsPerWorker := config.Requests / config.Concurrency
			for j := 0; j < requestsPerWorker; j++ {
				reqStart := time.Now()

				resp, err := b.client.Do(req.Clone(req.Context()))

				latency := time.Since(reqStart)

				mu.Lock()
				result.TotalRequests++
				if err == nil && resp.StatusCode < 400 {
					result.SuccessfulRequests++
				} else {
					result.FailedRequests++
				}
				latencies = append(latencies, latency)
				mu.Unlock()

				if resp != nil {
					resp.Body.Close()
				}
			}
		}()
	}

	wg.Wait()
	result.TotalTime = time.Since(startTime)

	// 计算统计信息
	b.calculateStatistics(&result, latencies)

	return result
}

// 计算统计信息
func (b *HTTPBenchmarker) calculateStatistics(result *BenchmarkResult, latencies []time.Duration) {
	if len(latencies) == 0 {
		return
	}

	// 排序延迟时间
	for i := 0; i < len(latencies); i++ {
		for j := i + 1; j < len(latencies); j++ {
			if latencies[i] > latencies[j] {
				latencies[i], latencies[j] = latencies[j], latencies[i]
			}
		}
	}

	result.MinLatency = latencies[0]
	result.MaxLatency = latencies[len(latencies)-1]

	var totalLatency time.Duration
	for _, latency := range latencies {
		totalLatency += latency
	}
	result.AverageLatency = totalLatency / time.Duration(len(latencies))

	// 计算百分位数
	result.LatencyPercentiles = make(map[string]time.Duration)
	result.LatencyPercentiles["50%"] = latencies[len(latencies)*50/100]
	result.LatencyPercentiles["90%"] = latencies[len(latencies)*90/100]
	result.LatencyPercentiles["95%"] = latencies[len(latencies)*95/100]
	result.LatencyPercentiles["99%"] = latencies[len(latencies)*99/100]

	result.RequestsPerSecond = float64(result.TotalRequests) / result.TotalTime.Seconds()
}

// 打印结果
func (b *HTTPBenchmarker) PrintResult(result BenchmarkResult) {
	fmt.Printf("\n=== 压测结果 ===\n")
	fmt.Printf("总请求数: %d\n", result.TotalRequests)
	fmt.Printf("成功请求数: %d\n", result.SuccessfulRequests)
	fmt.Printf("失败请求数: %d\n", result.FailedRequests)
	fmt.Printf("总耗时: %v\n", result.TotalTime)
	fmt.Printf("平均延迟: %v\n", result.AverageLatency)
	fmt.Printf("最小延迟: %v\n", result.MinLatency)
	fmt.Printf("最大延迟: %v\n", result.MaxLatency)
	fmt.Printf("QPS: %.2f\n", result.RequestsPerSecond)
	fmt.Printf("成功率: %.2f%%\n", float64(result.SuccessfulRequests)/float64(result.TotalRequests)*100)

	fmt.Printf("\n延迟百分位数:\n")
	fmt.Printf("  50%%: %v\n", result.LatencyPercentiles["50%"])
	fmt.Printf("  90%%: %v\n", result.LatencyPercentiles["90%"])
	fmt.Printf("  95%%: %v\n", result.LatencyPercentiles["95%"])
	fmt.Printf("  99%%: %v\n", result.LatencyPercentiles["99%"])
}

func main() {
	// 命令行参数
	var (
		url         = flag.String("url", "http://localhost:8080", "目标URL")
		requests    = flag.Int("n", 1000, "总请求数")
		concurrency = flag.Int("c", 10, "并发数")
		method      = flag.String("method", "GET", "HTTP方法")
		timeout     = flag.Duration("timeout", 30*time.Second, "请求超时时间")
	)
	flag.Parse()

	fmt.Println("🚀 HTTP 压测工具")
	fmt.Println("================")
	fmt.Printf("目标URL: %s\n", *url)
	fmt.Printf("总请求数: %d\n", *requests)
	fmt.Printf("并发数: %d\n", *concurrency)
	fmt.Printf("HTTP方法: %s\n", *method)
	fmt.Printf("超时时间: %v\n", *timeout)

	// 创建压测器
	benchmarker := NewHTTPBenchmarker(*timeout)

	// 配置压测参数
	config := BenchmarkConfig{
		URL:         *url,
		Requests:    *requests,
		Concurrency: *concurrency,
		Method:      *method,
		Headers: map[string]string{
			"User-Agent": "Benchmark-Tool/1.0",
		},
		Timeout: *timeout,
	}

	// 执行压测
	fmt.Println("\n开始压测...")
	result := benchmarker.RunBenchmark(config)

	// 打印结果
	benchmarker.PrintResult(result)
}
