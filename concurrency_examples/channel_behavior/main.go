package main

import (
	"fmt"
	"time"
)

// 演示无缓冲channel的行为
func unbufferedChannelDemo() {
	fmt.Println("=== 无缓冲Channel演示 ===")

	ch := make(chan int)

	// 发送方goroutine
	go func() {
		fmt.Println("发送方准备发送数据...")
		ch <- 42
		fmt.Println("数据已发送")
	}()

	// 接收方
	fmt.Println("接收方准备接收数据...")
	value := <-ch
	fmt.Printf("接收到数据: %d\n", value)
}

// 演示有缓冲channel的行为
func bufferedChannelDemo() {
	fmt.Println("\n=== 有缓冲Channel演示 ===")

	ch := make(chan int, 2) // 缓冲区大小为2

	// 可以发送多个数据而不阻塞
	fmt.Println("发送数据到有缓冲channel...")
	ch <- 1
	ch <- 2
	fmt.Println("已发送2个数据")

	// 接收数据
	fmt.Printf("接收数据1: %d\n", <-ch)
	fmt.Printf("接收数据2: %d\n", <-ch)
}

// 演示channel的阻塞行为
func channelBlockingDemo() {
	fmt.Println("\n=== Channel阻塞行为演示 ===")

	ch := make(chan int)

	// 启动接收方
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("接收方开始接收...")
		value := <-ch
		fmt.Printf("接收到: %d\n", value)
	}()

	// 发送方会阻塞直到接收方准备好
	fmt.Println("发送方准备发送...")
	ch <- 100
	fmt.Println("发送完成")
}

// 演示select语句的使用
func selectDemo() {
	fmt.Println("\n=== Select语句演示 ===")

	ch1 := make(chan string)
	ch2 := make(chan string)

	// 启动两个goroutine
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "来自channel1的消息"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "来自channel2的消息"
	}()

	// 使用select等待第一个可用的channel
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("接收到: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("接收到: %s\n", msg2)
		case <-time.After(3 * time.Second):
			fmt.Println("超时")
		}
	}
}

// 演示channel的关闭
func channelCloseDemo() {
	fmt.Println("\n=== Channel关闭演示 ===")

	ch := make(chan int, 3)

	// 发送数据
	ch <- 1
	ch <- 2
	ch <- 3

	// 关闭channel
	close(ch)

	// 从关闭的channel读取数据
	for i := 0; i < 4; i++ {
		value, ok := <-ch
		if ok {
			fmt.Printf("读取到: %d\n", value)
		} else {
			fmt.Println("channel已关闭")
		}
	}
}

// 演示range遍历channel
func channelRangeDemo() {
	fmt.Println("\n=== Channel Range遍历演示 ===")

	ch := make(chan int, 5)

	// 发送数据
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(ch)
	}()

	// 使用range遍历channel
	for value := range ch {
		fmt.Printf("接收到: %d\n", value)
	}
}

// 演示channel的方向性
func channelDirectionDemo() {
	fmt.Println("\n=== Channel方向性演示 ===")

	ch := make(chan int, 3)

	// 只发送的channel
	sendOnly := func(ch chan<- int) {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}

	// 只接收的channel
	receiveOnly := func(ch <-chan int) {
		for value := range ch {
			fmt.Printf("接收到: %d\n", value)
		}
	}

	// 启动goroutine
	go sendOnly(ch)
	receiveOnly(ch)
}

func main() {
	unbufferedChannelDemo()
	bufferedChannelDemo()
	channelBlockingDemo()
	selectDemo()
	channelCloseDemo()
	channelRangeDemo()
	channelDirectionDemo()
}
