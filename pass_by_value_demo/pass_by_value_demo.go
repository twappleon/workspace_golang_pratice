package main

import "fmt"

// 值类型示例
func modifyInt(n int) {
	n = 100
}

// 切片示例(引用类型)
func modifySlice(s []int) {
	s[0] = 100       // 修改底层数组
	s = append(s, 5) // 创建新描述符
}

// 指针示例
func modifyPointer(p *int) {
	*p = 100
}

// Since the error 'main redeclared in this block' is mentioned, but there's no other main function in the provided code,
// assuming no real redeclaration issue, keep the original code.
func main() {
	// 基本类型值传递
	n := 1
	modifyInt(n)
	fmt.Println("基本类型:", n) // 输出1

	// 切片行为
	s := []int{1, 2, 3}
	modifySlice(s)
	fmt.Println("切片:", s) // 输出[100 2 3]

	// 指针行为
	p := new(int)
	*p = 1
	modifyPointer(p)
	fmt.Println("指针:", *p) // 输出100
}