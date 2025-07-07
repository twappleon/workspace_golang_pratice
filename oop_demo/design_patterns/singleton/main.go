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

func (s *Singleton) DoSomething() {
	fmt.Println("Doing something")
}

func main() {
	instance := GetInstance(1)
	fmt.Println(&instance)
	instance.DoSomething()
	instance2 := GetInstance(2)
	fmt.Println(&instance2)
	instance2.DoSomething()
}
