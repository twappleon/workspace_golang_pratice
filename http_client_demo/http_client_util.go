package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// 5. 带加密的POST请求示例
func (c *HTTPClient) CreateUserWithEncryption(user *User, encryptionKey []byte) (*User, error) {
	encryptedData, err := encryptData(user, encryptionKey)
	if err != nil {
		return nil, fmt.Errorf("加密数据失败: %v", err)
	}

	encryptedRequest := struct {
		Data string `json:"encrypted_data"`
	}{
		Data: base64.StdEncoding.EncodeToString(encryptedData),
	}

	jsonData, err := json.Marshal(encryptedRequest)
	if err != nil {
		return nil, fmt.Errorf("序列化加密请求失败: %v", err)
	}

	url := fmt.Sprintf("%s/users/encrypted", c.baseURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Encryption", "AES-256-CBC")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("创建用户失败，状态码: %d", resp.StatusCode)
	}

	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v", err)
	}

	if !apiResp.Success {
		return nil, fmt.Errorf("API错误: %s", apiResp.Message)
	}

	userData, err := json.Marshal(apiResp.Data)
	if err != nil {
		return nil, fmt.Errorf("序列化用户数据失败: %v", err)
	}

	var createdUser User
	if err := json.Unmarshal(userData, &createdUser); err != nil {
		return nil, fmt.Errorf("反序列化用户数据失败: %v", err)
	}

	return &createdUser, nil
}

// 6. 带自定义Token认证的请求示例
func (c *HTTPClient) GetUserWithCustomToken(userID int, secretKey string) (*User, error) {
	tokenData := fmt.Sprintf("user_%d_%d_%s", userID, time.Now().Unix(), secretKey)
	tokenString := generateHash(tokenData)

	url := fmt.Sprintf("%s/users/%d", c.baseURL, userID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+tokenString)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("请求失败，状态码: %d", resp.StatusCode)
	}

	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v", err)
	}

	if !apiResp.Success {
		return nil, fmt.Errorf("API错误: %s", apiResp.Message)
	}

	userData, err := json.Marshal(apiResp.Data)
	if err != nil {
		return nil, fmt.Errorf("序列化用户数据失败: %v", err)
	}

	var user User
	if err := json.Unmarshal(userData, &user); err != nil {
		return nil, fmt.Errorf("反序列化用户数据失败: %v", err)
	}

	return &user, nil
}

// 7. 带重试机制的请求示例
func (c *HTTPClient) GetUserWithRetry(userID int, maxRetries int) (*User, error) {
	var lastErr error
	for i := 0; i <= maxRetries; i++ {
		user, err := c.GetUser(userID)
		if err == nil {
			return user, nil
		}
		lastErr = err
		log.Printf("第%d次尝试失败: %v", i+1, err)
		if i < maxRetries {
			backoff := time.Duration(1<<uint(i)) * time.Second
			time.Sleep(backoff)
		}
	}
	return nil, fmt.Errorf("重试%d次后仍然失败: %v", maxRetries, lastErr)
}

// 8. 批量请求示例
func (c *HTTPClient) GetUsersBatch(userIDs []int) ([]*User, error) {
	users := make([]*User, 0, len(userIDs))
	errors := make([]error, 0)
	type result struct {
		user  *User
		err   error
		index int
	}
	resultChan := make(chan result, len(userIDs))
	for i, userID := range userIDs {
		go func(id int, index int) {
			user, err := c.GetUser(id)
			resultChan <- result{user: user, err: err, index: index}
		}(userID, i)
	}
	for i := 0; i < len(userIDs); i++ {
		res := <-resultChan
		if res.err != nil {
			errors = append(errors, fmt.Errorf("用户ID %d: %v", userIDs[res.index], res.err))
		} else {
			users = append(users, res.user)
		}
	}
	if len(errors) > 0 {
		return users, fmt.Errorf("部分请求失败: %v", errors)
	}
	return users, nil
}

// 加密数据
func encryptData(data interface{}, key []byte) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	ciphertext := gcm.Seal(nonce, nonce, jsonData, nil)
	return ciphertext, nil
}

// 生成SHA256哈希
func generateHash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}
