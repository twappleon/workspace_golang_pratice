package main

import (
	"context"
	"runtime"
	"sync"
	"testing"
	"time"
)

// TestChannelCloseBySender 测试发送者关闭 channel
func TestChannelCloseBySender(t *testing.T) {
	ch := make(chan int, 3)

	// 发送数据
	ch <- 1
	ch <- 2
	ch <- 3

	// 发送者关闭 channel
	close(ch)

	// 接收数据
	values := []int{}
	for value := range ch {
		values = append(values, value)
	}

	expected := []int{1, 2, 3}
	if len(values) != len(expected) {
		t.Errorf("期望接收 %d 个值，实际接收 %d 个", len(expected), len(values))
	}

	for i, v := range values {
		if v != expected[i] {
			t.Errorf("索引 %d: 期望 %d，实际 %d", i, expected[i], v)
		}
	}
}

// TestChannelCloseDetection 测试 channel 关闭检测
func TestChannelCloseDetection(t *testing.T) {
	ch := make(chan int, 2)

	// 发送数据
	ch <- 10
	ch <- 20
	close(ch)

	// 测试关闭检测
	value1, ok1 := <-ch
	if !ok1 || value1 != 10 {
		t.Errorf("第一次接收失败: value=%d, ok=%v", value1, ok1)
	}

	value2, ok2 := <-ch
	if !ok2 || value2 != 20 {
		t.Errorf("第二次接收失败: value=%d, ok=%v", value2, ok2)
	}

	_, ok3 := <-ch
	if ok3 {
		t.Error("第三次接收应该返回 ok=false")
	}
}

// TestSelectWithClosedChannel 测试 select 与关闭的 channel
func TestSelectWithClosedChannel(t *testing.T) {
	ch := make(chan int)
	close(ch)

	// 使用 select 接收
	select {
	case value, ok := <-ch:
		if ok {
			t.Error("从关闭的 channel 接收应该返回 ok=false")
		}
		if value != 0 {
			t.Errorf("从关闭的 channel 接收应该返回零值，实际为 %d", value)
		}
	default:
		t.Error("select 应该能立即从关闭的 channel 接收")
	}
}

// TestContextCancellation 测试 context 取消
func TestContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan int)

	// 启动发送者
	go func() {
		for i := 0; i < 5; i++ {
			select {
			case ch <- i:
				time.Sleep(time.Millisecond * 10)
			case <-ctx.Done():
				return
			}
		}
		close(ch)
	}()

	// 启动接收者
	go func() {
		for value := range ch {
			if value == 2 {
				cancel() // 取消 context
				break
			}
		}
	}()

	// 等待一段时间让 goroutine 完成
	time.Sleep(time.Millisecond * 100)

	// 验证 context 被取消
	select {
	case <-ctx.Done():
		// 正常取消
	default:
		t.Error("context 应该被取消")
	}
}

// TestGoroutineLeakDetection 测试 goroutine 泄漏检测
func TestGoroutineLeakDetection(t *testing.T) {
	initialGoroutines := runtime.NumGoroutine()

	// 创建一些 goroutine
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond * 10)
		}()
	}

	wg.Wait()

	// 等待一段时间让 goroutine 完全结束
	time.Sleep(time.Millisecond * 50)

	finalGoroutines := runtime.NumGoroutine()

	// 验证 goroutine 数量没有异常增长
	if finalGoroutines > initialGoroutines+2 {
		t.Errorf("Goroutine 数量异常增长: 从 %d 到 %d", initialGoroutines, finalGoroutines)
	}
}

// TestChannelRangeLoop 测试 channel range 循环
func TestChannelRangeLoop(t *testing.T) {
	ch := make(chan int, 3)

	// 发送数据
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	// 使用 range 循环接收
	values := []int{}
	for value := range ch {
		values = append(values, value)
	}

	expected := []int{1, 2, 3}
	if len(values) != len(expected) {
		t.Errorf("期望接收 %d 个值，实际接收 %d 个", len(expected), len(values))
	}

	for i, v := range values {
		if v != expected[i] {
			t.Errorf("索引 %d: 期望 %d，实际 %d", i, expected[i], v)
		}
	}
}

// TestNonBlockingChannelOperations 测试非阻塞 channel 操作
func TestNonBlockingChannelOperations(t *testing.T) {
	ch := make(chan int, 1)

	// 非阻塞发送
	select {
	case ch <- 42:
		// 发送成功
	default:
		t.Error("应该能够发送到有缓冲的 channel")
	}

	// 非阻塞接收
	select {
	case value := <-ch:
		if value != 42 {
			t.Errorf("期望接收 42，实际为 %d", value)
		}
	default:
		t.Error("应该能够从有缓冲的 channel 接收")
	}

	// 测试空 channel 的非阻塞接收
	select {
	case <-ch:
		t.Error("空 channel 不应该有数据")
	default:
		// 这是期望的行为
	}
}

// BenchmarkChannelSendReceive 基准测试 channel 发送接收
func BenchmarkChannelSendReceive(b *testing.B) {
	ch := make(chan int, 1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch <- i
		<-ch
	}
}

// BenchmarkChannelClose 基准测试 channel 关闭
func BenchmarkChannelClose(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch := make(chan int, 10)
		for j := 0; j < 10; j++ {
			ch <- j
		}
		close(ch)
		for range ch {
			// 清空 channel
		}
	}
}

// BenchmarkSelectWithChannel 基准测试 select 与 channel
func BenchmarkSelectWithChannel(b *testing.B) {
	ch := make(chan int, 1)
	ch <- 1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		select {
		case value := <-ch:
			ch <- value
		default:
			// 空操作
		}
	}
}
