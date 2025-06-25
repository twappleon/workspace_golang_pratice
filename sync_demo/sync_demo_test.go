package main

import (
	"sync"
	"testing"
	"time"
)

// TestMutex 测试互斥锁
func TestMutex(t *testing.T) {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

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

// TestRWMutex 测试读写锁
func TestRWMutex(t *testing.T) {
	var data string
	var rwMu sync.RWMutex
	var wg sync.WaitGroup

	// 启动多个读取者
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			rwMu.RLock()
			_ = data // 读取数据
			rwMu.RUnlock()
		}(i)
	}

	// 启动一个写入者
	wg.Add(1)
	go func() {
		defer wg.Done()
		rwMu.Lock()
		data = "new data"
		rwMu.Unlock()
	}()

	wg.Wait()

	if data != "new data" {
		t.Errorf("期望数据为 'new data'，实际为 '%s'", data)
	}
}

// TestWaitGroup 测试等待组
func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	var counter int

	// 启动多个 goroutine
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}

	wg.Wait()

	if counter != 5 {
		t.Errorf("期望计数器为 5，实际为 %d", counter)
	}
}

// TestOnce 测试一次性执行
func TestOnce(t *testing.T) {
	var once sync.Once
	var counter int

	// 多次调用，但只执行一次
	for i := 0; i < 10; i++ {
		once.Do(func() {
			counter++
		})
	}

	if counter != 1 {
		t.Errorf("期望计数器为 1，实际为 %d", counter)
	}
}

// TestChannelSync 测试 channel 同步
func TestChannelSync(t *testing.T) {
	ch := make(chan int)

	// 发送者
	go func() {
		ch <- 42
	}()

	// 接收者
	value := <-ch

	if value != 42 {
		t.Errorf("期望值为 42，实际为 %d", value)
	}
}

// TestRaceCondition 测试竞态条件
func TestRaceCondition(t *testing.T) {
	var counter int
	var mu sync.Mutex

	// 不使用锁的情况（可能产生竞态条件）
	var unsafeCounter int
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			unsafeCounter++
		}()
	}

	wg.Wait()

	// 使用锁的情况
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()

	if counter != 1000 {
		t.Errorf("使用锁的计数器期望为 1000，实际为 %d", counter)
	}

	// 注意：unsafeCounter 可能不等于 1000，这是竞态条件的结果
	t.Logf("不安全计数器值: %d", unsafeCounter)
}

// BenchmarkMutex 基准测试互斥锁
func BenchmarkMutex(b *testing.B) {
	var counter int
	var mu sync.Mutex

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
}

// BenchmarkRWMutex 基准测试读写锁
func BenchmarkRWMutex(b *testing.B) {
	var data string
	var rwMu sync.RWMutex

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rwMu.RLock()
		_ = data
		rwMu.RUnlock()
	}
}

// BenchmarkWaitGroup 基准测试等待组
func BenchmarkWaitGroup(b *testing.B) {
	var wg sync.WaitGroup

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Microsecond)
		}()
	}
	wg.Wait()
}
