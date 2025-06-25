package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 演示goroutine泄漏
func goroutineLeakDemo() {
	fmt.Println("=== Goroutine泄漏演示 ===")

	// 获取初始goroutine数量
	initialGoroutines := runtime.NumGoroutine()
	fmt.Printf("初始goroutine数量: %d\n", initialGoroutines)

	// 启动泄漏的goroutine
	for i := 0; i < 10; i++ {
		go func(id int) {
			// 这个goroutine永远不会退出
			select {
			case <-time.After(time.Hour):
				// 永远不会到达这里
			}
		}(i)
	}

	// 等待一段时间让goroutine启动
	time.Sleep(100 * time.Millisecond)

	// 检查goroutine数量
	currentGoroutines := runtime.NumGoroutine()
	fmt.Printf("泄漏后goroutine数量: %d\n", currentGoroutines)
	fmt.Printf("泄漏的goroutine数量: %d\n", currentGoroutines-initialGoroutines)
}

// 演示正确的goroutine管理
func properGoroutineManagement() {
	fmt.Println("\n=== 正确的Goroutine管理 ===")

	initialGoroutines := runtime.NumGoroutine()
	fmt.Printf("初始goroutine数量: %d\n", initialGoroutines)

	var wg sync.WaitGroup

	// 启动正确管理的goroutine
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// 模拟工作
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Goroutine %d 完成工作\n", id)
		}(i)
	}

	// 等待所有goroutine完成
	wg.Wait()

	// 检查goroutine数量
	currentGoroutines := runtime.NumGoroutine()
	fmt.Printf("完成后goroutine数量: %d\n", currentGoroutines)
}

// 演示使用context控制goroutine
func contextControlDemo() {
	fmt.Println("\n=== Context控制Goroutine演示 ===")

	initialGoroutines := runtime.NumGoroutine()
	fmt.Printf("初始goroutine数量: %d\n", initialGoroutines)

	// 创建一个可以取消的context
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	// 启动受context控制的goroutine
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					fmt.Printf("Goroutine %d 收到取消信号\n", id)
					return
				case <-time.After(200 * time.Millisecond):
					fmt.Printf("Goroutine %d 执行任务\n", id)
				}
			}
		}(i)
	}

	// 让goroutine运行一段时间
	time.Sleep(1 * time.Second)

	// 取消所有goroutine
	fmt.Println("发送取消信号...")
	cancel()

	// 等待所有goroutine完成
	wg.Wait()

	currentGoroutines := runtime.NumGoroutine()
	fmt.Printf("取消后goroutine数量: %d\n", currentGoroutines)
}

// 演示channel泄漏
func channelLeakDemo() {
	fmt.Println("\n=== Channel泄漏演示 ===")

	initialGoroutines := runtime.NumGoroutine()
	fmt.Printf("初始goroutine数量: %d\n", initialGoroutines)

	// 创建channel
	ch := make(chan int)

	// 启动发送方goroutine
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}
		// 忘记关闭channel
	}()

	// 启动接收方goroutine
	go func() {
		for value := range ch {
			fmt.Printf("接收到: %d\n", value)
		}
	}()

	// 等待一段时间
	time.Sleep(1 * time.Second)

	currentGoroutines := runtime.NumGoroutine()
	fmt.Printf("Channel泄漏后goroutine数量: %d\n", currentGoroutines)
	fmt.Printf("泄漏的goroutine数量: %d\n", currentGoroutines-initialGoroutines)
}

// 演示正确的channel使用
func properChannelUsage() {
	fmt.Println("\n=== 正确的Channel使用 ===")

	initialGoroutines := runtime.NumGoroutine()
	fmt.Printf("初始goroutine数量: %d\n", initialGoroutines)

	ch := make(chan int)

	// 启动发送方goroutine
	go func() {
		defer close(ch) // 正确关闭channel
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// 启动接收方goroutine
	go func() {
		for value := range ch {
			fmt.Printf("接收到: %d\n", value)
		}
	}()

	// 等待完成
	time.Sleep(1 * time.Second)

	currentGoroutines := runtime.NumGoroutine()
	fmt.Printf("正确使用后goroutine数量: %d\n", currentGoroutines)
}

func main() {
	goroutineLeakDemo()
	properGoroutineManagement()
	contextControlDemo()
	channelLeakDemo()
	properChannelUsage()
}
