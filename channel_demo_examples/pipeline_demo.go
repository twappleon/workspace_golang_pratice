package main

import (
	"fmt"
	"math"
)

// 第一个stage: 生成数字
func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

// 第二个stage: 计算平方
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

// 第三个stage: 过滤偶数
func filterEven(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			if n%2 == 0 {
				out <- n
			}
		}
	}()
	return out
}

// 第四个stage: 计算平方根
func sqrt(in <-chan int) <-chan float64 {
	out := make(chan float64)
	go func() {
		defer close(out)
		for n := range in {
			out <- math.Sqrt(float64(n))
		}
	}()
	return out
}

func pipelineDemo() {
	fmt.Println("=== Pipeline Demo ===")

	// 构建pipeline: generate -> square -> filterEven -> sqrt
	numbers := generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	squared := square(numbers)
	evenSquared := filterEven(squared)
	sqrtEvenSquared := sqrt(evenSquared)

	// 消费结果
	fmt.Println("Pipeline results:")
	for result := range sqrtEvenSquared {
		fmt.Printf("Result: %.2f\n", result)
	}

	fmt.Println("\nPipeline explanation:")
	fmt.Println("1. Generate: 1,2,3,4,5,6,7,8,9,10")
	fmt.Println("2. Square: 1,4,9,16,25,36,49,64,81,100")
	fmt.Println("3. Filter even: 4,16,36,64,100")
	fmt.Println("4. Sqrt: 2.00,4.00,6.00,8.00,10.00")
}
