package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 演示竞态条件问题
func raceConditionDemo() {
	fmt.Println("=== 竞态条件演示 ===")

	var counter int
	var wg sync.WaitGroup

	// 启动多个goroutine同时修改counter
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter++ // 这里存在竞态条件
			}
		}()
	}

	wg.Wait()
	fmt.Printf("预期结果: 10000, 实际结果: %d\n", counter)
	fmt.Println("注意：由于竞态条件，结果通常小于10000")
}

// 使用互斥锁解决竞态条件
func mutexSolution() {
	fmt.Println("\n=== 使用互斥锁解决竞态条件 ===")

	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("预期结果: 10000, 实际结果: %d\n", counter)
}

// 使用原子操作解决竞态条件
func atomicSolution() {
	fmt.Println("\n=== 使用原子操作解决竞态条件 ===")

	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Printf("预期结果: 10000, 实际结果: %d\n", counter)
}

// 演示读写锁的使用
func rwMutexDemo() {
	fmt.Println("\n=== 读写锁演示 ===")

	var data map[string]int = make(map[string]int)
	var rwMu sync.RWMutex
	var wg sync.WaitGroup

	// 写入goroutine
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				key := fmt.Sprintf("key_%d_%d", id, j)
				rwMu.Lock()
				data[key] = j
				rwMu.Unlock()
			}
		}(i)
	}

	// 读取goroutine
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 50; j++ {
				key := fmt.Sprintf("key_%d_%d", id, j)
				rwMu.RLock()
				_, exists := data[key]
				rwMu.RUnlock()
				_ = exists
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("最终数据项数量: %d\n", len(data))
}

func main() {
	raceConditionDemo()
	mutexSolution()
	atomicSolution()
	rwMutexDemo()
}
