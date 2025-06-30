package main

import (
	"fmt"
	"time"
)

type worker struct {
	ID int
}

type WorkerManager struct {
	workerChan chan *worker
	mWorkers   int
}

func NewWorkerManager(mWorker int) *WorkerManager {
	return &WorkerManager{
		workerChan: make(chan *worker, mWorker),
		mWorkers:   mWorker,
	}
}

func main() {
	ch := make(chan int, 100)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println(a)
		}
	}()
	//defer close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 10)
}
