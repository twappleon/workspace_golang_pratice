package main

import (
	"sync"
	"testing"
	"time"
)

// TestWaitGroupBasic 测试基本的 WaitGroup 使用
func TestWaitGroupBasic(t *testing.T) {
	var wg sync.WaitGroup
	var counter int

	// 启动多个 goroutine
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}

	wg.Wait()

	if counter != 3 {
		t.Errorf("期望计数器为 3，实际为 %d", counter)
	}
}

// TestWaitGroupWithData 测试带数据的 WaitGroup
func TestWaitGroupWithData(t *testing.T) {
	var wg sync.WaitGroup
	results := make([]int, 0, 3)
	var mu sync.Mutex

	// 启动多个 goroutine 处理数据
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Millisecond * 10) // 模拟工作
			mu.Lock()
			results = append(results, id)
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	if len(results) != 3 {
		t.Errorf("期望结果数量为 3，实际为 %d", len(results))
	}

	// 检查是否包含所有 ID
	expected := map[int]bool{0: true, 1: true, 2: true}
	for _, id := range results {
		if !expected[id] {
			t.Errorf("意外的 ID: %d", id)
		}
	}
}

// TestWaitGroupTimeout 测试 WaitGroup 超时
func TestWaitGroupTimeout(t *testing.T) {
	var wg sync.WaitGroup
	done := make(chan bool)

	// 启动一个可能阻塞的 goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 50) // 模拟长时间工作
	}()

	// 等待完成或超时
	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case <-done:
		// 正常完成
	case <-time.After(time.Millisecond * 100):
		t.Error("WaitGroup 等待超时")
	}
}

// TestWaitGroupPanic 测试 WaitGroup 的 panic 情况
func TestWaitGroupPanic(t *testing.T) {
	var wg sync.WaitGroup

	// 测试 Add 负数
	defer func() {
		if r := recover(); r == nil {
			t.Error("期望 panic，但没有发生")
		}
	}()

	wg.Add(-1)
}

// TestWaitGroupDoneWithoutAdd 测试没有 Add 就调用 Done
func TestWaitGroupDoneWithoutAdd(t *testing.T) {
	var wg sync.WaitGroup

	defer func() {
		if r := recover(); r == nil {
			t.Error("期望 panic，但没有发生")
		}
	}()

	wg.Done()
}

// TestWaitGroupConcurrent 测试并发使用 WaitGroup
func TestWaitGroupConcurrent(t *testing.T) {
	var wg sync.WaitGroup
	var counter int
	var mu sync.Mutex

	// 启动多个 goroutine 并发增加计数器
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()

	if counter != 100 {
		t.Errorf("期望计数器为 100，实际为 %d", counter)
	}
}

// BenchmarkWaitGroup 基准测试 WaitGroup
func BenchmarkWaitGroup(b *testing.B) {
	var wg sync.WaitGroup

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 模拟一些工作
			_ = i * 2
		}()
	}
	wg.Wait()
}

// BenchmarkWaitGroupWithWork 基准测试带工作的 WaitGroup
func BenchmarkWaitGroupWithWork(b *testing.B) {
	var wg sync.WaitGroup

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 模拟实际工作
			time.Sleep(time.Microsecond)
		}()
	}
	wg.Wait()
}
