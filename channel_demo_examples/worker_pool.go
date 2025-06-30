package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker function that processes jobs from a channel
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(100 * time.Millisecond) // 模拟工作
		results <- job * 2                 // 返回处理结果
	}
}

func workerPoolDemo() {
	fmt.Println("=== Worker Pool Demo ===")

	const numJobs = 10
	const numWorkers = 3

	// 创建job和result channels
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	// 启动workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// 发送jobs
	go func() {
		for i := 1; i <= numJobs; i++ {
			jobs <- i
		}
		close(jobs) // 关闭jobs channel
	}()

	// 等待所有workers完成
	go func() {
		wg.Wait()
		close(results) // 关闭results channel
	}()

	// 收集结果
	fmt.Println("Results:")
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}

	fmt.Println("All jobs completed!")
}
