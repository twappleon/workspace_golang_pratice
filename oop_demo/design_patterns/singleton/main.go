package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type Singleton struct {
	Name int
}

var singleton *Singleton

func GetInstance(num int) *Singleton {
	once.Do(func() {
		singleton = &Singleton{num}
	})
	return singleton
}

func main() {
	instance := GetInstance(1)
	fmt.Println(instance)
	instance2 := GetInstance(2)
	fmt.Println(instance2)
}
