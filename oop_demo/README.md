# Go语言面向对象编程示例

本目录包含了展示Go语言面向对象特征的完整示例，分为基础概念和设计模式两个部分。

## 目录结构

```
oop_examples/
├── basic_oop/          # 基础面向对象概念
│   ├── go.mod
│   └── main.go
├── design_patterns/    # 高级设计模式
│   ├── go.mod
│   └── main.go
└── README.md
```

## 基础面向对象概念 (basic_oop)

### 1. 封装 (Encapsulation)
- 使用私有字段隐藏数据
- 通过方法提供公共接口
- 构造函数模式

**示例特性：**
- `BankAccount` 结构体：私有字段 `owner` 和 `balance`
- 公共方法：`Deposit()`, `Withdraw()`, `GetBalance()`, `GetOwner()`
- 构造函数：`NewBankAccount()`

### 2. 继承 (Inheritance)
- 通过结构体嵌入实现继承
- 方法重写和扩展

**示例特性：**
- `Animal` 基类：`Name`, `Age` 字段和 `Speak()`, `Move()` 方法
- `Dog` 和 `Cat` 继承 `Animal`，重写 `Speak()` 方法
- 新增特有方法：`Fetch()`, `Climb()`

### 3. 多态 (Polymorphism)
- 通过接口实现多态
- 接口组合

**示例特性：**
- `Speaker`, `Mover` 接口
- `AnimalBehavior` 组合接口
- 多态调用：`animal.Speak()`, `animal.Move()`

### 4. 接口 (Interfaces)
- 接口定义和实现
- 多态性

**示例特性：**
- `Shape` 接口：`Area()`, `Perimeter()` 方法
- `Circle`, `Rectangle` 实现 `Shape` 接口
- 接口切片：`[]Shape`

### 5. 方法接收者 (Method Receivers)
- 值接收者 vs 指针接收者
- 何时使用哪种接收者

**示例特性：**
- `Counter` 结构体
- 值接收者：`GetCount()` - 不修改对象
- 指针接收者：`Increment()`, `Reset()` - 修改对象

### 6. 工厂模式 (Factory Pattern)
- 工厂函数创建对象
- 接口返回类型

**示例特性：**
- `Vehicle` 接口
- `Car`, `Motorcycle` 实现
- `NewVehicle()` 工厂函数

### 7. 单例模式 (Singleton Pattern)
- 全局唯一实例
- 延迟初始化

**示例特性：**
- `Database` 单例
- 全局变量 `dbInstance`
- `GetDatabase()` 获取实例

## 高级设计模式 (design_patterns)

### 1. 观察者模式 (Observer Pattern)
- 发布-订阅模式
- 松耦合设计

**示例特性：**
- `Subject` 和 `Observer` 接口
- `NewsAgency` 主题
- `NewsChannel` 观察者

### 2. 策略模式 (Strategy Pattern)
- 算法族封装
- 运行时切换策略

**示例特性：**
- `PaymentStrategy` 接口
- `CreditCardPayment`, `AlipayPayment`, `WeChatPayment` 策略
- `ShoppingCart` 上下文

### 3. 装饰器模式 (Decorator Pattern)
- 动态扩展功能
- 组合优于继承

**示例特性：**
- `Coffee` 接口
- `SimpleCoffee` 基础组件
- `MilkDecorator`, `SugarDecorator` 装饰器

### 4. 建造者模式 (Builder Pattern)
- 复杂对象构建
- 链式调用

**示例特性：**
- `ComputerBuilder` 接口
- `GamingComputerBuilder` 具体建造者
- 链式方法调用

### 5. 线程安全单例模式 (Thread-Safe Singleton)
- 并发安全
- `sync.Once` 保证唯一性

**示例特性：**
- `ThreadSafeSingleton` 结构体
- `sync.Once` 保证线程安全
- 多goroutine测试

### 6. 组合模式 (Composite Pattern)
- 树形结构
- 统一接口

**示例特性：**
- `FileSystemComponent` 接口
- `File` 叶子节点
- `Directory` 复合节点

### 7. 命令模式 (Command Pattern)
- 请求封装
- 解耦调用者和接收者

**示例特性：**
- `Command` 接口
- `Light` 接收者
- `LightOnCommand`, `LightOffCommand` 具体命令
- `RemoteControl` 调用者

## 运行示例

### 运行基础面向对象示例
```bash
cd basic_oop
go run main.go
```

### 运行设计模式示例
```bash
cd design_patterns
go run main.go
```

### 运行所有示例
```bash
# 在项目根目录执行
./run_oop_examples.sh
```

## 学习要点

### Go语言面向对象特点
1. **组合优于继承**：通过嵌入结构体实现代码复用
2. **接口隐式实现**：只要实现了接口的所有方法，就自动实现了该接口
3. **方法接收者**：值接收者不修改原对象，指针接收者修改原对象
4. **构造函数模式**：使用 `New` 前缀的函数作为构造函数

### 设计模式在Go中的应用
1. **简洁性**：Go的设计模式实现通常比其他语言更简洁
2. **接口驱动**：大量使用接口实现多态和解耦
3. **函数式编程**：结合函数式编程特性
4. **并发安全**：考虑goroutine并发场景

## 扩展练习

1. 为 `BankAccount` 添加利息计算功能
2. 实现更多动物类型（如 `Bird`, `Fish`）
3. 添加更多支付策略（如 `PayPal`, `Bitcoin`）
4. 实现更多装饰器（如 `WhippedCreamDecorator`）
5. 创建更多命令类型（如 `StereoOnCommand`, `GarageDoorCommand`）

这些示例展示了Go语言如何通过结构体、方法、接口等特性实现面向对象编程，以及如何应用经典设计模式解决实际问题。 