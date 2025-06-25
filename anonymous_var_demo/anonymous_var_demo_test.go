package main

import (
	"testing"
)

// TestAnonymousVariableReturn 测试匿名变量处理返回值
func TestAnonymousVariableReturn(t *testing.T) {
	// 测试只获取第一个返回值
	first := getMultipleValues()

	if first != 42 {
		t.Errorf("期望第一个返回值为 42，实际为 %d", first)
	}
}

// TestAnonymousVariableStruct 测试匿名变量处理结构体
func TestAnonymousVariableStruct(t *testing.T) {
	// 测试只获取结构体的 Name 字段
	name := getPersonName()

	if name != "Alice" {
		t.Errorf("期望姓名为 'Alice'，实际为 '%s'", name)
	}
}

// TestAnonymousVariableMap 测试匿名变量处理 map
func TestAnonymousVariableMap(t *testing.T) {
	// 测试遍历 map 时忽略值
	colors := getColorMap()

	expectedKeys := map[string]bool{"red": true, "green": true, "blue": true}
	for color := range colors {
		if !expectedKeys[color] {
			t.Errorf("意外的颜色键: %s", color)
		}
	}
}

// TestAnonymousVariableTypeAssertion 测试匿名变量处理类型断言
func TestAnonymousVariableTypeAssertion(t *testing.T) {
	// 测试类型断言时忽略 ok 值
	result := getTypeAssertionResult()

	if result != "hello" {
		t.Errorf("期望结果为 'hello'，实际为 '%s'", result)
	}
}

// TestAnonymousVariableError 测试匿名变量处理错误
func TestAnonymousVariableError(t *testing.T) {
	// 测试忽略错误返回值
	value := getValueIgnoringError()

	if value != 100 {
		t.Errorf("期望值为 100，实际为 %d", value)
	}
}

// TestAnonymousVariableChannel 测试匿名变量处理 channel
func TestAnonymousVariableChannel(t *testing.T) {
	// 测试忽略 channel 接收的值
	ch := make(chan int, 1)
	ch <- 42

	// 只检查 channel 是否有数据，忽略值
	select {
	case <-ch:
		// 成功接收到数据
	default:
		t.Error("应该能接收到数据")
	}
}

// 辅助函数，模拟原代码中的函数
func getMultipleValues() int {
	return 42
}

func getPersonName() string {
	return "Alice"
}

func getColorMap() map[string]int {
	return map[string]int{
		"red":   0,
		"green": 1,
		"blue":  2,
	}
}

func getTypeAssertionResult() string {
	var i interface{} = "hello"
	return i.(string)
}

func getValueIgnoringError() int {
	return 100
}

// BenchmarkAnonymousVariable 基准测试匿名变量使用
func BenchmarkAnonymousVariable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = getMultipleValues()
		_ = getPersonName()
		_ = getValueIgnoringError()
	}
}

// BenchmarkTypeAssertion 基准测试类型断言
func BenchmarkTypeAssertion(b *testing.B) {
	var val interface{} = "hello"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = val.(string)
	}
}

// BenchmarkMapIteration 基准测试 map 遍历
func BenchmarkMapIteration(b *testing.B) {
	m := map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for range m {
			// 忽略值
		}
	}
}
