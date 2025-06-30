package main

import (
	"fmt"
)

func basicChannelDemo() {
	fmt.Println("=== Basic Channel Demo ===")

	// 1. 创建无缓冲channel
	ch := make(chan int)

	// 2. 发送数据到channel (在goroutine中)
	go func() {
		fmt.Println("Sending data to channel...")
		ch <- 42
		fmt.Println("Data sent to channel")
	}()

	// 3. 从channel接收数据
	fmt.Println("Waiting to receive data...")
	value := <-ch
	fmt.Printf("Received: %d\n", value)

	// 4. 有缓冲channel
	bufferedCh := make(chan string, 3)

	// 可以发送多个数据而不阻塞
	bufferedCh <- "Hello"
	bufferedCh <- "World"
	bufferedCh <- "Go"

	fmt.Println("Buffered channel contents:")
	fmt.Println(<-bufferedCh) // Hello
	fmt.Println(<-bufferedCh) // World
	fmt.Println(<-bufferedCh) // Go

	// 5. 关闭channel
	close(bufferedCh)

	// 6. 检查channel是否关闭
	value2, ok := <-bufferedCh
	fmt.Printf("Value: %s, Channel open: %t\n", value2, ok)
}
