package main

import (
	"log"
)

func main() {
	// 启动模拟服务器
	server := NewSimpleServer("8080")
	log.Println("启动HTTP模拟服务器...")
	server.Start()
}
