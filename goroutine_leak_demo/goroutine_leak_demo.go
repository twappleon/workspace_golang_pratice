package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 1. 使用context控制Goroutine退出
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done(): // 监听取消信号
				fmt.Println("Goroutine1: 收到取消信号，退出")
				return
			default:
				fmt.Println("Goroutine1: 执行工作...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// 2. 使用专用退出channel
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-quit: // 监听退出信号
				fmt.Println("Goroutine2: 收到退出信号，退出")
				return
			default:
				fmt.Println("Goroutine2: 执行工作...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	// 3. 确保channel被正确关闭
	dataCh := make(chan int)
	done := make(chan struct{})
	go func() {
		defer close(done)          // 确保退出时关闭done channel
		for data := range dataCh { // 自动检测channel关闭
			fmt.Printf("Goroutine3: 处理数据 %d\n", data)
		}
		fmt.Println("Goroutine3: dataCh已关闭，退出")
	}()

	// 模拟工作
	time.Sleep(2 * time.Second)

	// 发送取消信号
	fmt.Println("主程序: 发送取消信号")
	cancel()
	close(quit)

	// 发送数据并关闭channel
	dataCh <- 1
	dataCh <- 2
	close(dataCh) // 正确关闭channel

	// 等待所有Goroutine退出
	<-done
	time.Sleep(100 * time.Millisecond) // 确保日志输出
	fmt.Println("主程序: 所有Goroutine已退出")
}
