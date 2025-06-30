package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("Channel Demo Examples")
	fmt.Println("=====================")

	// 运行基础channel demo
	basicChannelDemo()

	fmt.Println("\n" + strings.Repeat("-", 50) + "\n")
	time.Sleep(1 * time.Second)

	// 运行select demo
	selectDemo()

	fmt.Println("\n" + strings.Repeat("-", 50) + "\n")
	time.Sleep(1 * time.Second)

	// 运行pipeline demo
	pipelineDemo()

	fmt.Println("\n" + strings.Repeat("-", 50) + "\n")
	time.Sleep(1 * time.Second)

	// 运行worker pool demo
	workerPoolDemo()

	fmt.Println("\nAll demos completed!")
}
