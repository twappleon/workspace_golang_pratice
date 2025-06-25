package main

import (
	"testing"
	"time"
)

// TestInitFunctionOrder 测试 init 函数的执行顺序
func TestInitFunctionOrder(t *testing.T) {
	// 这个测试主要验证 init 函数是否正常执行
	// 由于 init 函数在包加载时执行，我们只能验证包级变量是否正确初始化

	if msg == "" {
		t.Error("包级变量 msg 应该被正确初始化")
	}

	if msg != "Hello" {
		t.Errorf("期望 msg 为 'Hello'，实际为 '%s'", msg)
	}
}

// TestMainFunction 测试 main 函数的逻辑
func TestMainFunction(t *testing.T) {
	// 模拟 main 函数的逻辑
	expectedMsg := "Hello"
	if msg != expectedMsg {
		t.Errorf("期望消息为 '%s'，实际为 '%s'", expectedMsg, msg)
	}
}

// BenchmarkVariableInitialization 基准测试包级变量初始化
func BenchmarkVariableInitialization(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 重置包级变量（在实际情况下这不可行，这里只是演示）
		_ = msg
	}
}

// TestPackageLevelVariable 测试包级变量的行为
func TestPackageLevelVariable(t *testing.T) {
	// 验证包级变量在多次访问时保持一致
	firstAccess := msg
	time.Sleep(time.Millisecond) // 模拟一些时间间隔
	secondAccess := msg

	if firstAccess != secondAccess {
		t.Error("包级变量应该在多次访问时保持一致")
	}
}
