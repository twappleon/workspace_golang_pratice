package main

import (
	"fmt"
)

// 包级变量初始化
var msg = func() string {
	fmt.Println("变量初始化")
	return "Hello"
}()

// 第一个init函数
func init() {
	fmt.Println("init函数1执行")
}

// 第二个init函数
func init() {
	fmt.Println("init函数2执行")
}

func main() {
	fmt.Println("main函数执行:", msg)
}
