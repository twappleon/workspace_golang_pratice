package main

import (
	"fmt"
	"time"
)

func worker(id int, done chan struct{}) {
	defer fmt.Printf("Worker %d: 工作完成\n", id)
	
	fmt.Printf("Worker %d: 开始工作...\n", id)
	time.Sleep(time.Duration(id) * time.Second) // 模拟工作
	
	// 发送完成信号
	done <- struct{}{}
}

func main() {
	done := make(chan struct{}, 2) // 缓冲大小为worker数量
	
	// 启动两个worker
	go worker(1, done)
	go worker(2, done)
	
	// 等待所有worker完成
	for i := 0; i < 2; i++ {
		<-done // 等待每个worker发送完成信号
	}
	
	fmt.Println("主程序: 所有worker已完成")
}