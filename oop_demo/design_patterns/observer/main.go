package main

import "fmt"

type Listener interface {
	OnTeacherComing()
}

type Notifer interface {
	AddListener(listener Listener)
	RemoveListener(listener Listener)
	Notify()
}

// 实现层
type StudentZhang3 struct {
	Badthing string
}

func (s *StudentZhang3) OnTeacherComing() {
	fmt.Println("张3 停止 ", s.Badthing)
}

func (s *StudentZhang3) DoBadthing() {
	fmt.Println("张3 正在", s.Badthing)
}

type StudentLi4 struct {
	Badthing string
}

func (s *StudentLi4) OnTeacherComing() {
	fmt.Println("赵4 停止 ", s.Badthing)
}

func (s *StudentLi4) DoBadthing() {
	fmt.Println("赵4 正在", s.Badthing)
}

type ClassMonitor struct {
	listener []Listener
}

func (c *ClassMonitor) AddListener(listener Listener) {
	c.listener = append(c.listener, listener)
}

func (c *ClassMonitor) RemoveListener(listener Listener) {
	for i, l := range c.listener {
		if l == listener {
			c.listener = append(c.listener[:i], c.listener[i+1:]...)
		}
	}
}

func (c *ClassMonitor) Notify() {
	for _, listener := range c.listener {
		listener.OnTeacherComing()
	}
}

func main() {
	zhang3 := &StudentZhang3{
		Badthing: "抄作业",
	}
	li4 := &StudentLi4{
		Badthing: "玩王者荣耀",
	}

	fmt.Println("上课了，但是老师没有来，学生们都在忙自己的事...")
	zhang3.DoBadthing()
	li4.DoBadthing()

	classMonitor := &ClassMonitor{}
	classMonitor.AddListener(zhang3)
	classMonitor.AddListener(li4)

	fmt.Println("这时候老师来了，班长给学什么使了一个眼神...")
	classMonitor.Notify()
}
