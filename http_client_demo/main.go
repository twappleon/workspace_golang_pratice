package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func main() {
	// 启动模拟服务器
	go func() {
		// 使用系统命令启动独立的服务器
		cmd := exec.Command("go", "run", "server/.")
		cmd.Dir = "."
		cmd.Start()
	}()

	// 等待服务器启动
	time.Sleep(3 * time.Second)

	// 创建HTTP客户端
	client := NewHTTPClient("http://localhost:8080", "your-api-key-here")

	fmt.Println("=== GET请求示例 ===")
	user, err := client.GetUser(1)
	if err != nil {
		log.Printf("获取用户失败: %v", err)
	} else {
		fmt.Printf("获取到用户: %+v\n", user)
	}

	fmt.Println("\n=== POST JSON请求示例 ===")
	newUser := &User{
		Name:     "张三",
		Email:    "zhangsan@example.com",
		Password: "password123",
	}
	createdUser, err := client.CreateUser(newUser)
	if err != nil {
		log.Printf("创建用户失败: %v", err)
	} else {
		fmt.Printf("创建用户成功: %+v\n", createdUser)
	}

	fmt.Println("\n=== POST Form请求示例 ===")
	token, err := client.LoginWithForm("zhangsan@example.com", "password123")
	if err != nil {
		log.Printf("登录失败: %v", err)
	} else {
		fmt.Printf("登录成功，获得token: %s\n", token)
	}

	fmt.Println("\n=== POST Raw数据请求示例 ===")
	fileData := []byte("这是文件内容")
	err = client.UploadFile("test.txt", fileData)
	if err != nil {
		log.Printf("上传文件失败: %v", err)
	} else {
		fmt.Println("文件上传成功")
	}

	fmt.Println("\n=== 带加密的POST请求示例 ===")
	encryptionKey := []byte("your-32-byte-encryption-key-here")
	encryptedUser, err := client.CreateUserWithEncryption(newUser, encryptionKey)
	if err != nil {
		log.Printf("创建加密用户失败: %v", err)
	} else {
		fmt.Printf("创建加密用户成功: %+v\n", encryptedUser)
	}

	fmt.Println("\n=== 带自定义Token认证的请求示例 ===")
	customTokenUser, err := client.GetUserWithCustomToken(1, "your-secret-key")
	if err != nil {
		log.Printf("自定义Token认证获取用户失败: %v", err)
	} else {
		fmt.Printf("自定义Token认证获取用户成功: %+v\n", customTokenUser)
	}

	fmt.Println("\n=== 带重试机制的请求示例 ===")
	retryUser, err := client.GetUserWithRetry(1, 3)
	if err != nil {
		log.Printf("重试获取用户失败: %v", err)
	} else {
		fmt.Printf("重试获取用户成功: %+v\n", retryUser)
	}

	fmt.Println("\n=== 批量请求示例 ===")
	userIDs := []int{1, 2, 3, 4, 5}
	users, err := client.GetUsersBatch(userIDs)
	if err != nil {
		log.Printf("批量获取用户失败: %v", err)
	} else {
		fmt.Printf("批量获取用户成功，共获取%d个用户\n", len(users))
		for i, user := range users {
			fmt.Printf("用户%d: %+v\n", i+1, user)
		}
	}

	fmt.Println("\n=== 哈希生成示例 ===")
	password := "password123"
	hash := generateHash(password)
	fmt.Printf("密码: %s\n哈希值: %s\n", password, hash)
}
