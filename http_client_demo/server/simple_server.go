package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// User 用户结构体
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

// APIResponse API响应结构体
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SimpleServer 简化的HTTP服务器
type SimpleServer struct {
	users map[int]*User
	port  string
}

// NewSimpleServer 创建新的简化服务器
func NewSimpleServer(port string) *SimpleServer {
	return &SimpleServer{
		users: make(map[int]*User),
		port:  port,
	}
}

// 初始化测试数据
func (s *SimpleServer) initTestData() {
	s.users[1] = &User{
		ID:    1,
		Name:  "张三",
		Email: "zhangsan@example.com",
	}
	s.users[2] = &User{
		ID:    2,
		Name:  "李四",
		Email: "lisi@example.com",
	}
	s.users[3] = &User{
		ID:    3,
		Name:  "王五",
		Email: "wangwu@example.com",
	}
}

// 启动服务器
func (s *SimpleServer) Start() {
	s.initTestData()

	// 设置路由
	http.HandleFunc("/users/", s.handleUsers)
	http.HandleFunc("/login", s.handleLogin)
	http.HandleFunc("/upload", s.handleUpload)
	http.HandleFunc("/users/encrypted", s.handleEncryptedUser)

	log.Printf("简化服务器启动在端口 %s", s.port)
	log.Fatal(http.ListenAndServe(":"+s.port, nil))
}

// 处理用户相关请求
func (s *SimpleServer) handleUsers(w http.ResponseWriter, r *http.Request) {
	// 验证API Key
	if !s.validateAPIKey(r) {
		s.sendResponse(w, http.StatusUnauthorized, APIResponse{
			Success: false,
			Message: "无效的API Key",
		})
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/users/")

	switch r.Method {
	case "GET":
		if path != "" {
			s.getUser(w, r, path)
		}
	case "POST":
		if path == "" {
			s.createUser(w, r)
		}
	default:
		s.sendResponse(w, http.StatusMethodNotAllowed, APIResponse{
			Success: false,
			Message: "不支持的HTTP方法",
		})
	}
}

// 获取用户
func (s *SimpleServer) getUser(w http.ResponseWriter, r *http.Request, idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		s.sendResponse(w, http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "无效的用户ID",
		})
		return
	}

	user, exists := s.users[id]
	if !exists {
		s.sendResponse(w, http.StatusNotFound, APIResponse{
			Success: false,
			Message: "用户不存在",
		})
		return
	}

	s.sendResponse(w, http.StatusOK, APIResponse{
		Success: true,
		Message: "获取用户成功",
		Data:    user,
	})
}

// 创建用户
func (s *SimpleServer) createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		s.sendResponse(w, http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "无效的请求数据",
		})
		return
	}

	// 生成新ID
	user.ID = len(s.users) + 1
	s.users[user.ID] = &user

	s.sendResponse(w, http.StatusCreated, APIResponse{
		Success: true,
		Message: "创建用户成功",
		Data:    user,
	})
}

// 处理登录
func (s *SimpleServer) handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		s.sendResponse(w, http.StatusMethodNotAllowed, APIResponse{
			Success: false,
			Message: "只支持POST方法",
		})
		return
	}

	// 解析表单数据
	if err := r.ParseForm(); err != nil {
		s.sendResponse(w, http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "解析表单失败",
		})
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		s.sendResponse(w, http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "用户名和密码不能为空",
		})
		return
	}

	// 简单的用户验证
	if username == "zhangsan@example.com" && password == "password123" {
		// 生成简单的token
		token := fmt.Sprintf("token_%s_%d", username, time.Now().Unix())
		s.sendResponse(w, http.StatusOK, APIResponse{
			Success: true,
			Message: "登录成功",
			Data:    token,
		})
	} else {
		s.sendResponse(w, http.StatusUnauthorized, APIResponse{
			Success: false,
			Message: "用户名或密码错误",
		})
	}
}

// 处理文件上传
func (s *SimpleServer) handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		s.sendResponse(w, http.StatusMethodNotAllowed, APIResponse{
			Success: false,
			Message: "只支持POST方法",
		})
		return
	}

	// 验证API Key
	if !s.validateAPIKey(r) {
		s.sendResponse(w, http.StatusUnauthorized, APIResponse{
			Success: false,
			Message: "无效的API Key",
		})
		return
	}

	filename := r.Header.Get("X-Filename")
	if filename == "" {
		s.sendResponse(w, http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "缺少文件名",
		})
		return
	}

	// 读取文件内容
	body, err := io.ReadAll(r.Body)
	if err != nil {
		s.sendResponse(w, http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "读取文件失败",
		})
		return
	}

	log.Printf("接收到文件上传: %s, 大小: %d bytes", filename, len(body))

	s.sendResponse(w, http.StatusOK, APIResponse{
		Success: true,
		Message: "文件上传成功",
		Data: map[string]interface{}{
			"filename": filename,
			"size":     len(body),
		},
	})
}

// 处理加密用户创建
func (s *SimpleServer) handleEncryptedUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		s.sendResponse(w, http.StatusMethodNotAllowed, APIResponse{
			Success: false,
			Message: "只支持POST方法",
		})
		return
	}

	// 验证API Key
	if !s.validateAPIKey(r) {
		s.sendResponse(w, http.StatusUnauthorized, APIResponse{
			Success: false,
			Message: "无效的API Key",
		})
		return
	}

	var encryptedRequest struct {
		Data string `json:"encrypted_data"`
	}

	if err := json.NewDecoder(r.Body).Decode(&encryptedRequest); err != nil {
		s.sendResponse(w, http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "无效的请求数据",
		})
		return
	}

	// 为了演示，直接创建一个新用户
	user := &User{
		ID:    len(s.users) + 1,
		Name:  "加密用户",
		Email: "encrypted@example.com",
	}

	s.users[user.ID] = user

	s.sendResponse(w, http.StatusCreated, APIResponse{
		Success: true,
		Message: "创建加密用户成功",
		Data:    user,
	})
}

// 验证API Key
func (s *SimpleServer) validateAPIKey(r *http.Request) bool {
	authHeader := r.Header.Get("Authorization")
	return strings.HasPrefix(authHeader, "Bearer your-api-key-here")
}

// 发送响应
func (s *SimpleServer) sendResponse(w http.ResponseWriter, statusCode int, response APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
