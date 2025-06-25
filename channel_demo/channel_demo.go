package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 无缓冲Channel的阻塞行为
	unbuffered := make(chan int)
	go func() {
		time.Sleep(1 * time.Second) // 模拟延迟
		fmt.Println("无缓冲Channel: 接收者准备就绪")
		fmt.Println("接收值:", <-unbuffered)
	}()
	fmt.Println("无缓冲Channel: 发送者尝试发送")
	unbuffered <- 42 // 会阻塞直到接收者就绪

	// 2. 有缓冲Channel的阻塞行为
	buffered := make(chan int, 2)
	buffered <- 1
	buffered <- 2
	fmt.Println("有缓冲Channel: 前两个发送成功")
	go func() {
		time.Sleep(1 * time.Second)
		<-buffered // 腾出一个位置
	}()
	fmt.Println("有缓冲Channel: 第三个发送将阻塞")
	buffered <- 3 // 会阻塞直到有空间

	// 3. 非阻塞操作(select + default)
	select {
	case val := <-buffered:
		fmt.Println("非阻塞接收成功:", val)
	default:
		fmt.Println("非阻塞接收: 无数据可接收")
	}

	select {
	case buffered <- 4:
		fmt.Println("非阻塞发送成功")
	default:
		fmt.Println("非阻塞发送: 无空间可发送")
	}

	// 4. Channel关闭行为
	ch := make(chan int, 3)
	ch <- 10
	ch <- 20
	close(ch) // 正确关闭Channel

	// 从已关闭Channel接收
	val1, ok1 := <-ch
	fmt.Printf("从关闭Channel接收1: val=%v, ok=%v\n", val1, ok1)
	val2, ok2 := <-ch
	fmt.Printf("从关闭Channel接收2: val=%v, ok=%v\n", val2, ok2)
	val3, ok3 := <-ch
	fmt.Printf("从关闭Channel接收3: val=%v, ok=%v\n", val3, ok3)

	// 5. 错误操作演示(注释掉避免panic)

	// 向已关闭Channel发送会panic
	// Since 'ch' is already declared above, we remove this redundant declaration.
	// We will just use the existing 'ch' variable later.
	ch = make(chan int, 1)
	ch <- 30

	// 重复关闭Channel会panic
	close(ch)

	val4, ok4 := <-ch
	fmt.Printf("从关闭Channel接收4: val=%v, ok=%v\n", val4, ok4)

}
