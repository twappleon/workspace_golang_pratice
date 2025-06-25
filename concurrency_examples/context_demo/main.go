package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 演示基本的context使用
func basicContextDemo() {
	fmt.Println("=== 基本Context演示 ===")

	// 创建带超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 启动goroutine
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context已取消:", ctx.Err())
				return
			case <-time.After(500 * time.Millisecond):
				fmt.Println("执行任务...")
			}
		}
	}()

	// 等待context取消
	<-ctx.Done()
	fmt.Println("主程序收到取消信号")
}

// 演示context传递
func contextPropagationDemo() {
	fmt.Println("\n=== Context传递演示 ===")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 启动多个层级的goroutine
	go func() {
		fmt.Println("第一层goroutine启动")

		go func() {
			fmt.Println("第二层goroutine启动")

			go func() {
				fmt.Println("第三层goroutine启动")

				select {
				case <-ctx.Done():
					fmt.Println("第三层收到取消信号:", ctx.Err())
				case <-time.After(5 * time.Second):
					fmt.Println("第三层任务完成")
				}
			}()

			select {
			case <-ctx.Done():
				fmt.Println("第二层收到取消信号:", ctx.Err())
			case <-time.After(4 * time.Second):
				fmt.Println("第二层任务完成")
			}
		}()

		select {
		case <-ctx.Done():
			fmt.Println("第一层收到取消信号:", ctx.Err())
		case <-time.After(3 * time.Second):
			fmt.Println("第一层任务完成")
		}
	}()

	// 等待context取消
	<-ctx.Done()
	time.Sleep(100 * time.Millisecond) // 给goroutine时间退出
}

// 演示context值传递
func contextValueDemo() {
	fmt.Println("\n=== Context值传递演示 ===")

	// 创建带值的context
	ctx := context.WithValue(context.Background(), "user_id", "12345")
	ctx = context.WithValue(ctx, "request_id", "req_67890")

	// 启动goroutine
	go func() {
		// 获取context中的值
		userID := ctx.Value("user_id")
		requestID := ctx.Value("request_id")

		fmt.Printf("用户ID: %v\n", userID)
		fmt.Printf("请求ID: %v\n", requestID)

		// 模拟工作
		time.Sleep(1 * time.Second)
		fmt.Println("任务完成")
	}()

	time.Sleep(2 * time.Second)
}

// 演示context取消
func contextCancellationDemo() {
	fmt.Println("\n=== Context取消演示 ===")

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	// 启动多个goroutine
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					fmt.Printf("Goroutine %d 收到取消信号: %v\n", id, ctx.Err())
					return
				case <-time.After(500 * time.Millisecond):
					fmt.Printf("Goroutine %d 执行任务\n", id)
				}
			}
		}(i)
	}

	// 让goroutine运行一段时间
	time.Sleep(2 * time.Second)

	// 取消所有goroutine
	fmt.Println("发送取消信号...")
	cancel()

	// 等待所有goroutine完成
	wg.Wait()
	fmt.Println("所有goroutine已退出")
}

// 演示context超时
func contextTimeoutDemo() {
	fmt.Println("\n=== Context超时演示 ===")

	// 创建带超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// 启动长时间运行的goroutine
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("任务因超时而取消:", ctx.Err())
		case <-time.After(3 * time.Second):
			fmt.Println("任务完成")
		}
	}()

	// 等待context取消
	<-ctx.Done()
	fmt.Println("主程序收到超时信号")
}

// 演示context死线
func contextDeadlineDemo() {
	fmt.Println("\n=== Context死线演示 ===")

	// 创建带死线的context
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	fmt.Printf("死线时间: %v\n", deadline)

	// 启动goroutine
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("任务因死线而取消:", ctx.Err())
				return
			case <-time.After(500 * time.Millisecond):
				fmt.Println("执行任务...")
			}
		}
	}()

	// 等待context取消
	<-ctx.Done()
	fmt.Println("主程序收到死线信号")
}

func main() {
	basicContextDemo()
	contextPropagationDemo()
	contextValueDemo()
	contextCancellationDemo()
	contextTimeoutDemo()
	contextDeadlineDemo()
}
