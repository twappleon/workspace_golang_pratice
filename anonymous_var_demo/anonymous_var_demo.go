package main

import (
	"fmt"
	_ "os" // 仅执行init函数
)

func multiReturn() (int, string) {
	return 42, "answer"
}

type Person struct {
	Name string
	Age  int
}

func main() {
	// 场景1：忽略部分返回值
	num, _ := multiReturn()
	fmt.Println("只获取第一个返回值:", num)

	// 场景2：结构体解构忽略字段
	p := Person{"Alice", 30}
	var name string
	name, _ = p.Name, p.Age
	fmt.Println("只获取Name字段:", name)

	// 场景3：循环中忽略索引
	colors := []string{"red", "green", "blue"}
	for idx, _ := range colors {
		fmt.Println("颜色:", colors[idx], idx)
	}

	// 场景4：类型断言忽略ok值
	var i interface{} = "hello"
	str, _ := i.(string)
	fmt.Println("类型断言结果:", str)
}
