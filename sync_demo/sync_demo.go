package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 1. 使用Channel同步
	ch := make(chan int)
	go func() {
		fmt.Println("Channel: Goroutine发送数据")
		ch <- 42
	}()
	fmt.Println("Channel: 主程序接收数据:", <-ch)

	// 2. 使用WaitGroup等待多个Goroutine
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("WaitGroup: Goroutine %d 执行\n", id)
		}(i)
	}
	wg.Wait()

	// 3. 使用Mutex保护共享数据
	var counter int
	var mu sync.Mutex
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("Mutex: 最终计数器值:", counter)

	// 4. 使用RWMutex读写锁
	var rwMu sync.RWMutex
	var data string
	// 写操作
	go func() {
		rwMu.Lock()
		data = "新数据"
		rwMu.Unlock()
	}()
	// 读操作
	go func() {
		rwMu.RLock()
		fmt.Println("RWMutex: 读取数据:", data)
		rwMu.RUnlock()
	}()
	time.Sleep(100 * time.Millisecond)

	// 5. 使用Once确保只执行一次
	var once sync.Once
	for i := 0; i < 3; i++ {
		go func() {
			once.Do(func() {
				fmt.Println("Once: 这段代码只会执行一次")
			})
		}()
	}
	time.Sleep(100 * time.Millisecond)
}