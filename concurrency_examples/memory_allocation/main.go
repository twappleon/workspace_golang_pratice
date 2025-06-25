package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
	"time"
)

// 演示内存分配
func memoryAllocationDemo() {
	fmt.Println("=== 内存分配演示 ===")

	// 获取初始内存统计
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("初始内存使用: %d MB\n", m.Alloc/1024/1024)

	// 分配大量内存
	var slices [][]int
	for i := 0; i < 1000; i++ {
		slice := make([]int, 10000)
		slices = append(slices, slice)
	}

	// 再次获取内存统计
	runtime.ReadMemStats(&m)
	fmt.Printf("分配后内存使用: %d MB\n", m.Alloc/1024/1024)

	// 清空切片，触发GC
	slices = nil
	runtime.GC()

	// 最终内存统计
	runtime.ReadMemStats(&m)
	fmt.Printf("GC后内存使用: %d MB\n", m.Alloc/1024/1024)
}

// 演示对象池的使用
func objectPoolDemo() {
	fmt.Println("\n=== 对象池演示 ===")

	// 创建对象池
	type ExpensiveObject struct {
		ID   int
		Data []byte
	}

	// 模拟创建昂贵对象
	createObject := func() *ExpensiveObject {
		return &ExpensiveObject{
			ID:   time.Now().Nanosecond(),
			Data: make([]byte, 1024),
		}
	}

	// 不使用对象池
	fmt.Println("不使用对象池:")
	start := time.Now()
	for i := 0; i < 1000; i++ {
		obj := createObject()
		_ = obj // 使用对象
	}
	fmt.Printf("耗时: %v\n", time.Since(start))

	// 使用对象池
	fmt.Println("\n使用对象池:")
	pool := make(chan *ExpensiveObject, 10)

	// 预填充池
	for i := 0; i < 10; i++ {
		pool <- createObject()
	}

	start = time.Now()
	for i := 0; i < 1000; i++ {
		var obj *ExpensiveObject
		select {
		case obj = <-pool:
			// 从池中获取对象
		default:
			// 池为空，创建新对象
			obj = createObject()
		}

		// 使用对象
		_ = obj

		// 将对象放回池中
		select {
		case pool <- obj:
			// 成功放回池中
		default:
			// 池已满，丢弃对象
		}
	}
	fmt.Printf("耗时: %v\n", time.Since(start))
}

// 演示内存泄漏
func memoryLeakDemo() {
	fmt.Println("\n=== 内存泄漏演示 ===")

	// 模拟内存泄漏
	var leakedSlices [][]int

	go func() {
		for i := 0; i < 100; i++ {
			// 分配内存但不释放
			slice := make([]int, 1000000)
			leakedSlices = append(leakedSlices, slice)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// 监控内存使用
	for i := 0; i < 10; i++ {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("内存使用: %d MB\n", m.Alloc/1024/1024)
		time.Sleep(200 * time.Millisecond)
	}
}

// 演示GC调优
func gcTuningDemo() {
	fmt.Println("\n=== GC调优演示 ===")

	// 设置GC目标百分比
	debug.SetGCPercent(50) // 默认是100

	// 获取当前GC设置
	fmt.Printf("GC目标百分比: %d%%\n", debug.SetGCPercent(-1))

	// 强制GC
	fmt.Println("强制GC...")
	runtime.GC()

	// 获取GC统计信息
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("GC次数: %d\n", m.NumGC)
	fmt.Printf("总GC时间: %v\n", time.Duration(m.PauseTotalNs))
}

// 演示并发内存分配
func concurrentMemoryAllocation() {
	fmt.Println("\n=== 并发内存分配演示 ===")

	var wg sync.WaitGroup
	var mu sync.Mutex
	var totalAllocated int64

	// 启动多个goroutine进行内存分配
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			var localAllocated int64
			for j := 0; j < 100; j++ {
				// 分配内存
				data := make([]byte, 1024*1024) // 1MB
				_ = data
				localAllocated += 1024 * 1024

				time.Sleep(10 * time.Millisecond)
			}

			mu.Lock()
			totalAllocated += localAllocated
			mu.Unlock()

			fmt.Printf("Goroutine %d 分配了 %d MB\n", id, localAllocated/1024/1024)
		}(i)
	}

	wg.Wait()
	fmt.Printf("总分配内存: %d MB\n", totalAllocated/1024/1024)
}

func main() {
	memoryAllocationDemo()
	objectPoolDemo()
	memoryLeakDemo()
	gcTuningDemo()
	concurrentMemoryAllocation()
}
