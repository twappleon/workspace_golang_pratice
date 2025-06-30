package main

import (
	"fmt"
	"time"
)

func selectDemo() {
	fmt.Println("=== Select Demo ===")

	ch1 := make(chan string)
	ch2 := make(chan string)

	// 启动两个goroutine发送数据
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "from channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "from channel 2"
	}()

	// 使用select等待多个channel
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("Received: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("Received: %s\n", msg2)
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout!")
		}
	}

	// Select with default (non-blocking)
	fmt.Println("\n=== Non-blocking Select ===")
	ch3 := make(chan int)

	select {
	case <-ch3:
		fmt.Println("Received from ch3")
	default:
		fmt.Println("No data available, continuing...")
	}

	// 发送数据到ch3
	go func() {
		ch3 <- 100
	}()

	// 再次尝试接收
	select {
	case val := <-ch3:
		fmt.Printf("Received: %d\n", val)
	default:
		fmt.Println("No data available")
	}
}
