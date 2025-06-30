package main

import (
	"fmt"
	"math"
)

// 1. 封装 - 使用结构体和私有字段
type BankAccount struct {
	owner   string  // 私有字段
	balance float64 // 私有字段
}

// 构造函数
func NewBankAccount(owner string, initialBalance float64) *BankAccount {
	return &BankAccount{
		owner:   owner,
		balance: initialBalance,
	}
}

// 方法 - 封装行为
func (ba *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("存款金额必须大于0")
	}
	ba.balance += amount
	return nil
}

func (ba *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("取款金额必须大于0")
	}
	if amount > ba.balance {
		return fmt.Errorf("余额不足")
	}
	ba.balance -= amount
	return nil
}

// Getter方法
func (ba *BankAccount) GetBalance() float64 {
	return ba.balance
}

func (ba *BankAccount) GetOwner() string {
	return ba.owner
}

// 2. 继承 - 通过嵌入结构体实现
type Animal struct {
	Name string
	Age  int
}

func (a *Animal) Speak() {
	fmt.Printf("%s 发出声音\n", a.Name)
}

func (a *Animal) Move() {
	fmt.Printf("%s 在移动\n", a.Name)
}

// Dog继承Animal
type Dog struct {
	Animal // 嵌入结构体，实现继承
	Breed  string
}

// 重写方法
func (d *Dog) Speak() {
	fmt.Printf("%s 汪汪叫\n", d.Name)
}

// 新增方法
func (d *Dog) Fetch() {
	fmt.Printf("%s 在捡球\n", d.Name)
}

// Cat继承Animal
type Cat struct {
	Animal
	Color string
}

func (c *Cat) Speak() {
	fmt.Printf("%s 喵喵叫\n", c.Name)
}

func (c *Cat) Climb() {
	fmt.Printf("%s 在爬树\n", c.Name)
}

// 3. 多态 - 通过接口实现
type Speaker interface {
	Speak()
}

type Mover interface {
	Move()
}

// 组合接口
type AnimalBehavior interface {
	Speaker
	Mover
}

// 4. 接口实现
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 5. 方法接收者类型
type Counter struct {
	count int
}

// 值接收者 - 不修改原对象
func (c Counter) GetCount() int {
	return c.count
}

// 指针接收者 - 修改原对象
func (c *Counter) Increment() {
	c.count++
}

func (c *Counter) Reset() {
	c.count = 0
}

// 6. 工厂模式
type Vehicle interface {
	Drive()
	GetType() string
}

type Car struct {
	Brand string
}

func (c *Car) Drive() {
	fmt.Printf("驾驶 %s 汽车\n", c.Brand)
}

func (c *Car) GetType() string {
	return "汽车"
}

type Motorcycle struct {
	Brand string
}

func (m *Motorcycle) Drive() {
	fmt.Printf("驾驶 %s 摩托车\n", m.Brand)
}

func (m *Motorcycle) GetType() string {
	return "摩托车"
}

// 工厂函数
func NewVehicle(vehicleType, brand string) Vehicle {
	switch vehicleType {
	case "car":
		return &Car{Brand: brand}
	case "motorcycle":
		return &Motorcycle{Brand: brand}
	default:
		return nil
	}
}

// 7. 单例模式
type Database struct {
	connection string
}

var dbInstance *Database

func GetDatabase() *Database {
	if dbInstance == nil {
		dbInstance = &Database{connection: "mysql://localhost:3306"}
		fmt.Println("创建数据库连接")
	}
	return dbInstance
}

func (db *Database) Query(sql string) {
	fmt.Printf("执行查询: %s\n", sql)
}

// 演示函数
func demonstrateEncapsulation() {
	fmt.Println("=== 封装演示 ===")
	account := NewBankAccount("张三", 1000.0)

	err := account.Deposit(500.0)
	if err != nil {
		fmt.Println("存款失败:", err)
	}

	err = account.Withdraw(200.0)
	if err != nil {
		fmt.Println("取款失败:", err)
	}

	fmt.Printf("账户: %s, 余额: %.2f\n", account.GetOwner(), account.GetBalance())
}

func demonstrateInheritance() {
	fmt.Println("\n=== 继承演示 ===")

	dog := &Dog{
		Animal: Animal{Name: "旺财", Age: 3},
		Breed:  "金毛",
	}

	cat := &Cat{
		Animal: Animal{Name: "咪咪", Age: 2},
		Color:  "橘色",
	}

	dog.Speak() // 重写的方法
	dog.Move()  // 继承的方法
	dog.Fetch() // 新增的方法

	cat.Speak() // 重写的方法
	cat.Move()  // 继承的方法
	cat.Climb() // 新增的方法
}

func demonstratePolymorphism() {
	fmt.Println("\n=== 多态演示 ===")

	animals := []AnimalBehavior{
		&Dog{Animal: Animal{Name: "旺财", Age: 3}},
		&Cat{Animal: Animal{Name: "咪咪", Age: 2}},
	}

	for _, animal := range animals {
		animal.Speak() // 多态调用
		animal.Move()  // 多态调用
	}
}

func demonstrateInterfaces() {
	fmt.Println("\n=== 接口演示 ===")

	shapes := []Shape{
		&Circle{Radius: 5.0},
		&Rectangle{Width: 4.0, Height: 6.0},
	}

	for _, shape := range shapes {
		fmt.Printf("形状类型: %T\n", shape)
		fmt.Printf("面积: %.2f\n", shape.Area())
		fmt.Printf("周长: %.2f\n", shape.Perimeter())
		fmt.Println()
	}
}

func demonstrateMethodReceivers() {
	fmt.Println("=== 方法接收者演示 ===")

	counter := &Counter{}

	fmt.Printf("初始计数: %d\n", counter.GetCount())

	counter.Increment()
	counter.Increment()

	fmt.Printf("增加后计数: %d\n", counter.GetCount())

	counter.Reset()
	fmt.Printf("重置后计数: %d\n", counter.GetCount())
}

func demonstrateFactoryPattern() {
	fmt.Println("\n=== 工厂模式演示 ===")

	vehicles := []Vehicle{
		NewVehicle("car", "丰田"),
		NewVehicle("motorcycle", "本田"),
	}

	for _, vehicle := range vehicles {
		fmt.Printf("车辆类型: %s\n", vehicle.GetType())
		vehicle.Drive()
		fmt.Println()
	}
}

func demonstrateSingletonPattern() {
	fmt.Println("=== 单例模式演示 ===")

	db1 := GetDatabase()
	db2 := GetDatabase()

	fmt.Printf("db1 地址: %p\n", db1)
	fmt.Printf("db2 地址: %p\n", db2)
	fmt.Printf("是否为同一实例: %t\n", db1 == db2)

	db1.Query("SELECT * FROM users")
}

func main() {
	demonstrateEncapsulation()
	demonstrateInheritance()
	demonstratePolymorphism()
	demonstrateInterfaces()
	demonstrateMethodReceivers()
	demonstrateFactoryPattern()
	demonstrateSingletonPattern()
}
