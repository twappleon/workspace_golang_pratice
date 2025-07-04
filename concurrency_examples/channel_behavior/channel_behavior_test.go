package main

import (
	"testing"
	"time"
)

// TestUnbufferedChannel 测试无缓冲 channel
func TestUnbufferedChannel(t *testing.T) {
	ch := make(chan int)

	// 启动发送者
	go func() {
		ch <- 42
	}()

	// 接收者
	value := <-ch

	if value != 42 {
		t.Errorf("期望值为 42，实际为 %d", value)
	}
}

// TestBufferedChannel 测试有缓冲 channel
func TestBufferedChannel(t *testing.T) {
	ch := make(chan int, 3)

	// 发送数据到缓冲 channel
	ch <- 1
	ch <- 2
	ch <- 3

	// 接收数据
	values := []int{}
	for i := 0; i < 3; i++ {
		values = append(values, <-ch)
	}

	expected := []int{1, 2, 3}
	for i, v := range values {
		if v != expected[i] {
			t.Errorf("索引 %d: 期望 %d，实际 %d", i, expected[i], v)
		}
	}
}

// TestChannelClose 测试 channel 关闭
func TestChannelClose(t *testing.T) {
	ch := make(chan int, 2)

	// 发送数据
	ch <- 10
	ch <- 20

	// 关闭 channel
	close(ch)

	// 接收数据
	value1, ok1 := <-ch
	value2, ok2 := <-ch
	_, ok3 := <-ch

	if !ok1 || value1 != 10 {
		t.Errorf("第一次接收失败: value=%d, ok=%v", value1, ok1)
	}

	if !ok2 || value2 != 20 {
		t.Errorf("第二次接收失败: value=%d, ok=%v", value2, ok2)
	}

	if ok3 {
		t.Errorf("第三次接收应该失败，但 ok=%v", ok3)
	}
}

// TestSelectStatement 测试 select 语句
func TestSelectStatement(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 启动发送者
	go func() {
		time.Sleep(time.Millisecond * 10)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(time.Millisecond * 20)
		ch2 <- 2
	}()

	// 使用 select 接收
	received := false
	select {
	case value := <-ch1:
		if value != 1 {
			t.Errorf("期望从 ch1 接收 1，实际为 %d", value)
		}
		received = true
	case value := <-ch2:
		if value != 2 {
			t.Errorf("期望从 ch2 接收 2，实际为 %d", value)
		}
		received = true
	case <-time.After(time.Millisecond * 100):
		t.Error("select 超时")
	}

	if !received {
		t.Error("应该接收到数据")
	}
}

// TestChannelRange 测试 channel range 循环
func TestChannelRange(t *testing.T) {
	ch := make(chan int, 3)

	// 发送数据
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	// 使用 range 接收
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

// BenchmarkChannelSendReceive 基准测试 channel 发送接收
func BenchmarkChannelSendReceive(b *testing.B) {
	ch := make(chan int, 1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch <- i
		<-ch
	}
}

// BenchmarkBufferedChannel 基准测试有缓冲 channel
func BenchmarkBufferedChannel(b *testing.B) {
	ch := make(chan int, 1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch <- i
	}

	// 清空 channel
	for len(ch) > 0 {
		<-ch
	}
}
