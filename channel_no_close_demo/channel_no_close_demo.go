package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 演示1: Channel 不关闭导致的内存泄漏
func memoryLeakDemo() {
	fmt.Println("=== Channel 不关闭导致的内存泄漏 ===")

	// 创建一个 channel
	ch := make(chan int)

	// 启动一个 goroutine 发送数据
	go func() {
		for i := 0; i < 1000; i++ {
			ch <- i
			time.Sleep(time.Millisecond)
		}
		// 注意：这里没有关闭 channel
	}()

	// 主 goroutine 只接收前 100 个数据
	for i := 0; i < 100; i++ {
		value := <-ch
		fmt.Printf("接收: %d\n", value)
	}

	// 此时发送 goroutine 还在运行，但主 goroutine 已经退出
	// 这会导致发送 goroutine 永远阻塞在 ch <- i 上
	fmt.Println("主 goroutine 退出，但发送 goroutine 仍在运行...")

	// 显示当前 goroutine 数量
	fmt.Printf("当前 goroutine 数量: %d\n", runtime.NumGoroutine())
}

// 演示2: 多个接收者等待关闭信号
func multipleReceiversDemo() {
	fmt.Println("\n=== 多个接收者等待关闭信号 ===")

	ch := make(chan int)
	var wg sync.WaitGroup

	// 启动多个接收者
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				value, ok := <-ch
				if !ok {
					fmt.Printf("接收者 %d: channel 已关闭，退出\n", id)
					return
				}
				fmt.Printf("接收者 %d: 收到 %d\n", id, value)
			}
		}(i)
	}

	// 发送一些数据
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 100)
	}

	// 关闭 channel，通知所有接收者
	close(ch)
	fmt.Println("Channel 已关闭")

	// 等待所有接收者完成
	wg.Wait()
	fmt.Printf("当前 goroutine 数量: %d\n", runtime.NumGoroutine())
}

// 演示3: 使用 select 避免阻塞
func selectDemo() {
	fmt.Println("\n=== 使用 select 避免阻塞 ===")

	ch := make(chan int)
	done := make(chan bool)

	// 发送者
	go func() {
		for i := 0; i < 10; i++ {
			select {
			case ch <- i:
				fmt.Printf("发送: %d\n", i)
			case <-done:
				fmt.Println("发送者收到退出信号")
				return
			}
			time.Sleep(time.Millisecond * 100)
		}
	}()

	// 接收者
	go func() {
		for i := 0; i < 5; i++ {
			select {
			case value := <-ch:
				fmt.Printf("接收: %d\n", value)
			case <-time.After(time.Second):
				fmt.Println("接收超时，退出")
				done <- true
				return
			}
		}
		done <- true
	}()

	time.Sleep(time.Second * 2)
	fmt.Printf("当前 goroutine 数量: %d\n", runtime.NumGoroutine())
}

// 演示4: 正确的 Channel 使用模式
func correctUsageDemo() {
	fmt.Println("\n=== 正确的 Channel 使用模式 ===")

	ch := make(chan int)
	var wg sync.WaitGroup

	// 发送者
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch) // 发送者负责关闭 channel

		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Millisecond * 100)
		}
		fmt.Println("发送者完成，关闭 channel")
	}()

	// 接收者
	wg.Add(1)
	go func() {
		defer wg.Done()

		for value := range ch { // 使用 range 循环，会自动检测 channel 关闭
			fmt.Printf("接收: %d\n", value)
		}
		fmt.Println("接收者完成")
	}()

	wg.Wait()
	fmt.Printf("当前 goroutine 数量: %d\n", runtime.NumGoroutine())
}

// 演示5: 检测 goroutine 泄漏
func goroutineLeakDetection() {
	fmt.Println("\n=== Goroutine 泄漏检测 ===")

	initialGoroutines := runtime.NumGoroutine()
	fmt.Printf("初始 goroutine 数量: %d\n", initialGoroutines)

	// 模拟有问题的代码
	ch := make(chan int)
	go func() {
		ch <- 1
		// 没有关闭 channel，goroutine 会一直存在
	}()

	value := <-ch
	fmt.Printf("接收到: %d\n", value)

	// 等待一段时间让 goroutine 有机会退出
	time.Sleep(time.Millisecond * 100)

	finalGoroutines := runtime.NumGoroutine()
	fmt.Printf("最终 goroutine 数量: %d\n", finalGoroutines)

	if finalGoroutines > initialGoroutines {
		fmt.Printf("⚠️  检测到 goroutine 泄漏！增加了 %d 个 goroutine\n",
			finalGoroutines-initialGoroutines)
	} else {
		fmt.Println("✅ 没有检测到 goroutine 泄漏")
	}
}

func main() {

	fmt.Println("🚀 Channel 不关闭的问题演示")
	fmt.Println("==========================")

	// 注意：这些演示可能会导致 goroutine 泄漏
	// 在实际运行中，建议逐个运行并观察结果

	// 演示1: 内存泄漏
	memoryLeakDemo()

	// 演示2: 多个接收者
	multipleReceiversDemo()

	// 演示3: select 使用
	selectDemo()

	// 演示4: 正确用法
	correctUsageDemo()

	// 演示5: 泄漏检测
	goroutineLeakDetection()

	fmt.Println("\n=== 总结 ===")
	fmt.Println("Channel 不关闭可能导致的问题：")
	fmt.Println("1. Goroutine 泄漏 - 发送者或接收者永远阻塞")
	fmt.Println("2. 内存泄漏 - 相关的 goroutine 无法被垃圾回收")
	fmt.Println("3. 程序行为异常 - 接收者无法知道何时停止等待")
	fmt.Println("4. 资源浪费 - 占用系统资源")

	fmt.Println("\n最佳实践：")
	fmt.Println("1. 发送者负责关闭 channel")
	fmt.Println("2. 接收者使用 range 循环或检查 ok 值")
	fmt.Println("3. 使用 select 避免阻塞")
	fmt.Println("4. 使用 context 或 done channel 控制生命周期")
}
