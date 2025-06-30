package main

import (
	"fmt"
	"sync"
)

// 1. 观察者模式
type Observer interface {
	Update(message string)
}

type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify(message string)
}

type NewsAgency struct {
	observers []Observer
	news      string
}

func NewNewsAgency() *NewsAgency {
	return &NewsAgency{
		observers: make([]Observer, 0),
	}
}

func (na *NewsAgency) Attach(observer Observer) {
	na.observers = append(na.observers, observer)
}

func (na *NewsAgency) Detach(observer Observer) {
	for i, obs := range na.observers {
		if obs == observer {
			na.observers = append(na.observers[:i], na.observers[i+1:]...)
			break
		}
	}
}

func (na *NewsAgency) Notify(message string) {
	na.news = message
	for _, observer := range na.observers {
		observer.Update(message)
	}
}

type NewsChannel struct {
	name string
}

func NewNewsChannel(name string) *NewsChannel {
	return &NewsChannel{name: name}
}

func (nc *NewsChannel) Update(message string) {
	fmt.Printf("[%s] 收到新闻: %s\n", nc.name, message)
}

// 2. 策略模式
type PaymentStrategy interface {
	Pay(amount float64) string
}

type CreditCardPayment struct{}

func (cc *CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("使用信用卡支付 %.2f 元", amount)
}

type AlipayPayment struct{}

func (ap *AlipayPayment) Pay(amount float64) string {
	return fmt.Sprintf("使用支付宝支付 %.2f 元", amount)
}

type WeChatPayment struct{}

func (wp *WeChatPayment) Pay(amount float64) string {
	return fmt.Sprintf("使用微信支付 %.2f 元", amount)
}

type ShoppingCart struct {
	paymentStrategy PaymentStrategy
}

func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{}
}

func (sc *ShoppingCart) SetPaymentStrategy(strategy PaymentStrategy) {
	sc.paymentStrategy = strategy
}

func (sc *ShoppingCart) Checkout(amount float64) string {
	if sc.paymentStrategy == nil {
		return "请选择支付方式"
	}
	return sc.paymentStrategy.Pay(amount)
}

// 3. 装饰器模式
type Coffee interface {
	Cost() float64
	Description() string
}

type SimpleCoffee struct{}

func (sc *SimpleCoffee) Cost() float64 {
	return 10.0
}

func (sc *SimpleCoffee) Description() string {
	return "简单咖啡"
}

type MilkDecorator struct {
	coffee Coffee
}

func NewMilkDecorator(coffee Coffee) *MilkDecorator {
	return &MilkDecorator{coffee: coffee}
}

func (md *MilkDecorator) Cost() float64 {
	return md.coffee.Cost() + 2.0
}

func (md *MilkDecorator) Description() string {
	return md.coffee.Description() + " + 牛奶"
}

type SugarDecorator struct {
	coffee Coffee
}

func NewSugarDecorator(coffee Coffee) *SugarDecorator {
	return &SugarDecorator{coffee: coffee}
}

func (sd *SugarDecorator) Cost() float64 {
	return sd.coffee.Cost() + 1.0
}

func (sd *SugarDecorator) Description() string {
	return sd.coffee.Description() + " + 糖"
}

// 4. 建造者模式
type Computer struct {
	CPU      string
	Memory   string
	Storage  string
	Graphics string
}

type ComputerBuilder interface {
	SetCPU(cpu string) ComputerBuilder
	SetMemory(memory string) ComputerBuilder
	SetStorage(storage string) ComputerBuilder
	SetGraphics(graphics string) ComputerBuilder
	Build() *Computer
}

type GamingComputerBuilder struct {
	computer *Computer
}

func NewGamingComputerBuilder() *GamingComputerBuilder {
	return &GamingComputerBuilder{
		computer: &Computer{},
	}
}

func (gcb *GamingComputerBuilder) SetCPU(cpu string) ComputerBuilder {
	gcb.computer.CPU = cpu
	return gcb
}

func (gcb *GamingComputerBuilder) SetMemory(memory string) ComputerBuilder {
	gcb.computer.Memory = memory
	return gcb
}

func (gcb *GamingComputerBuilder) SetStorage(storage string) ComputerBuilder {
	gcb.computer.Storage = storage
	return gcb
}

func (gcb *GamingComputerBuilder) SetGraphics(graphics string) ComputerBuilder {
	gcb.computer.Graphics = graphics
	return gcb
}

func (gcb *GamingComputerBuilder) Build() *Computer {
	return gcb.computer
}

// 5. 线程安全的单例模式
type ThreadSafeSingleton struct {
	data string
}

var (
	instance *ThreadSafeSingleton
	once     sync.Once
)

func GetThreadSafeInstance() *ThreadSafeSingleton {
	once.Do(func() {
		instance = &ThreadSafeSingleton{data: "线程安全单例"}
		fmt.Println("创建线程安全单例实例")
	})
	return instance
}

func (ts *ThreadSafeSingleton) GetData() string {
	return ts.data
}

// 6. 组合模式
type FileSystemComponent interface {
	Display(indent string)
}

type File struct {
	name string
	size int
}

func NewFile(name string, size int) *File {
	return &File{name: name, size: size}
}

func (f *File) Display(indent string) {
	fmt.Printf("%s文件: %s (大小: %d KB)\n", indent, f.name, f.size)
}

type Directory struct {
	name     string
	children []FileSystemComponent
}

func NewDirectory(name string) *Directory {
	return &Directory{
		name:     name,
		children: make([]FileSystemComponent, 0),
	}
}

func (d *Directory) Add(component FileSystemComponent) {
	d.children = append(d.children, component)
}

func (d *Directory) Remove(component FileSystemComponent) {
	for i, child := range d.children {
		if child == component {
			d.children = append(d.children[:i], d.children[i+1:]...)
			break
		}
	}
}

func (d *Directory) Display(indent string) {
	fmt.Printf("%s目录: %s\n", indent, d.name)
	for _, child := range d.children {
		child.Display(indent + "  ")
	}
}

// 7. 命令模式
type Command interface {
	Execute()
}

type Light struct {
	location string
}

func NewLight(location string) *Light {
	return &Light{location: location}
}

func (l *Light) TurnOn() {
	fmt.Printf("%s 的灯打开了\n", l.location)
}

func (l *Light) TurnOff() {
	fmt.Printf("%s 的灯关闭了\n", l.location)
}

type LightOnCommand struct {
	light *Light
}

func NewLightOnCommand(light *Light) *LightOnCommand {
	return &LightOnCommand{light: light}
}

func (loc *LightOnCommand) Execute() {
	loc.light.TurnOn()
}

type LightOffCommand struct {
	light *Light
}

func NewLightOffCommand(light *Light) *LightOffCommand {
	return &LightOffCommand{light: light}
}

func (loc *LightOffCommand) Execute() {
	loc.light.TurnOff()
}

type RemoteControl struct {
	commands []Command
}

func NewRemoteControl() *RemoteControl {
	return &RemoteControl{
		commands: make([]Command, 0),
	}
}

func (rc *RemoteControl) AddCommand(command Command) {
	rc.commands = append(rc.commands, command)
}

func (rc *RemoteControl) PressButton(index int) {
	if index < len(rc.commands) {
		rc.commands[index].Execute()
	}
}

// 演示函数
func demonstrateObserverPattern() {
	fmt.Println("=== 观察者模式演示 ===")

	newsAgency := NewNewsAgency()

	channel1 := NewNewsChannel("CCTV")
	channel2 := NewNewsChannel("BBC")
	channel3 := NewNewsChannel("CNN")

	newsAgency.Attach(channel1)
	newsAgency.Attach(channel2)
	newsAgency.Attach(channel3)

	newsAgency.Notify("重大新闻：Go语言发布新版本！")

	newsAgency.Detach(channel2)
	newsAgency.Notify("BBC已取消订阅")
}

func demonstrateStrategyPattern() {
	fmt.Println("\n=== 策略模式演示 ===")

	cart := NewShoppingCart()

	// 使用不同支付策略
	strategies := []PaymentStrategy{
		&CreditCardPayment{},
		&AlipayPayment{},
		&WeChatPayment{},
	}

	for _, strategy := range strategies {
		cart.SetPaymentStrategy(strategy)
		fmt.Println(cart.Checkout(100.0))
	}
}

func demonstrateDecoratorPattern() {
	fmt.Println("\n=== 装饰器模式演示 ===")

	coffee := &SimpleCoffee{}
	fmt.Printf("基础咖啡: %s, 价格: %.2f\n", coffee.Description(), coffee.Cost())

	// 添加牛奶
	milkCoffee := NewMilkDecorator(coffee)
	fmt.Printf("加牛奶: %s, 价格: %.2f\n", milkCoffee.Description(), milkCoffee.Cost())

	// 添加糖
	sugarMilkCoffee := NewSugarDecorator(milkCoffee)
	fmt.Printf("加糖加牛奶: %s, 价格: %.2f\n", sugarMilkCoffee.Description(), sugarMilkCoffee.Cost())
}

func demonstrateBuilderPattern() {
	fmt.Println("\n=== 建造者模式演示 ===")

	builder := NewGamingComputerBuilder()
	computer := builder.
		SetCPU("Intel i9-12900K").
		SetMemory("32GB DDR5").
		SetStorage("2TB NVMe SSD").
		SetGraphics("RTX 4090").
		Build()

	fmt.Printf("游戏电脑配置:\n")
	fmt.Printf("CPU: %s\n", computer.CPU)
	fmt.Printf("内存: %s\n", computer.Memory)
	fmt.Printf("存储: %s\n", computer.Storage)
	fmt.Printf("显卡: %s\n", computer.Graphics)
}

func demonstrateThreadSafeSingleton() {
	fmt.Println("\n=== 线程安全单例模式演示 ===")

	// 模拟多个goroutine同时获取实例
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			instance := GetThreadSafeInstance()
			fmt.Printf("Goroutine %d 获取实例: %s\n", id, instance.GetData())
		}(i)
	}

	wg.Wait()
}

func demonstrateCompositePattern() {
	fmt.Println("\n=== 组合模式演示 ===")

	// 创建文件系统结构
	root := NewDirectory("根目录")

	documents := NewDirectory("文档")
	documents.Add(NewFile("报告.txt", 100))
	documents.Add(NewFile("计划.docx", 250))

	pictures := NewDirectory("图片")
	pictures.Add(NewFile("照片1.jpg", 2048))
	pictures.Add(NewFile("照片2.jpg", 1536))

	root.Add(documents)
	root.Add(pictures)
	root.Add(NewFile("README.md", 50))

	// 显示文件系统结构
	root.Display("")
}

func demonstrateCommandPattern() {
	fmt.Println("\n=== 命令模式演示 ===")

	// 创建接收者
	livingRoomLight := NewLight("客厅")
	kitchenLight := NewLight("厨房")

	// 创建命令
	livingRoomLightOn := NewLightOnCommand(livingRoomLight)
	livingRoomLightOff := NewLightOffCommand(livingRoomLight)
	kitchenLightOn := NewLightOnCommand(kitchenLight)
	kitchenLightOff := NewLightOffCommand(kitchenLight)

	// 创建调用者
	remote := NewRemoteControl()
	remote.AddCommand(livingRoomLightOn)
	remote.AddCommand(livingRoomLightOff)
	remote.AddCommand(kitchenLightOn)
	remote.AddCommand(kitchenLightOff)

	// 执行命令
	fmt.Println("按下按钮 0:")
	remote.PressButton(0)

	fmt.Println("按下按钮 1:")
	remote.PressButton(1)

	fmt.Println("按下按钮 2:")
	remote.PressButton(2)

	fmt.Println("按下按钮 3:")
	remote.PressButton(3)
}

func main() {
	demonstrateObserverPattern()
	demonstrateStrategyPattern()
	demonstrateDecoratorPattern()
	demonstrateBuilderPattern()
	demonstrateThreadSafeSingleton()
	demonstrateCompositePattern()
	demonstrateCommandPattern()
}
