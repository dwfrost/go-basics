package httpclient

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

// DemonstrateAuthentication 展示HTTP认证
func DemonstrateAuthentication() {
	// 基本认证
	fmt.Println("4.1 基本认证")
	basicAuth()

	// Bearer令牌认证
	fmt.Println("\n4.2 Bearer令牌认证")
	bearerTokenAuth()

	// API密钥认证
	fmt.Println("\n4.3 API密钥认证")
	apiKeyAuth()

	// OAuth2认证 (仅展示概念)
	fmt.Println("\n4.4 OAuth2认证 (概念示例)")
	fmt.Println("OAuth2认证流程较为复杂，通常需要以下步骤:")
	fmt.Println("1. 获取授权码")
	fmt.Println("2. 使用授权码获取访问令牌")
	fmt.Println("3. 使用访问令牌调用API")
	fmt.Println("4. 刷新访问令牌")
}

// 基本认证
func basicAuth() {
	// 创建一个HTTP客户端
	client := &http.Client{}

	// 创建请求
	req, err := http.NewRequest("GET", "https://httpbin.org/basic-auth/user/passwd", nil)
	if err != nil {
		fmt.Printf("创建请求失败: %v\n", err)
		return
	}

	// 添加基本认证头
	username := "user"
	password := "passwd"
	auth := username + ":" + password
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", "Basic "+encodedAuth)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("发送请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return
	}

	// 打印响应
	fmt.Printf("状态码: %d\n", resp.StatusCode)
	fmt.Printf("响应体: %s\n", body)
}

// Bearer令牌认证
func bearerTokenAuth() {
	// 创建一个HTTP客户端
	client := &http.Client{}

	// 创建请求
	req, err := http.NewRequest("GET", "https://httpbin.org/bearer", nil)
	if err != nil {
		fmt.Printf("创建请求失败: %v\n", err)
		return
	}

	// 添加Bearer令牌头
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	req.Header.Add("Authorization", "Bearer "+token)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("发送请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return
	}

	// 打印响应
	fmt.Printf("状态码: %d\n", resp.StatusCode)
	fmt.Printf("响应体: %s\n", body)
}

// API密钥认证
func apiKeyAuth() {
	// 方法1: 在URL中添加API密钥
	apiKey := "abcdef123456"
	url := fmt.Sprintf("https://httpbin.org/get?api_key=%s", apiKey)

	// 发送请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("发送请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return
	}

	// 打印响应
	fmt.Printf("方法1 (URL参数) - 状态码: %d\n", resp.StatusCode)
	if len(body) > 100 {
		fmt.Printf("响应体 (前100字节): %s...\n", body[:100])
	} else {
		fmt.Printf("响应体: %s\n", body)
	}

	// 方法2: 在请求头中添加API密钥
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://httpbin.org/get", nil)
	if err != nil {
		fmt.Printf("创建请求失败: %v\n", err)
		return
	}

	// 添加API密钥到请求头
	req.Header.Add("X-API-Key", apiKey)

	// 发送请求
	resp2, err := client.Do(req)
	if err != nil {
		fmt.Printf("发送请求失败: %v\n", err)
		return
	}
	defer resp2.Body.Close()

	// 读取响应
	body2, err := io.ReadAll(resp2.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return
	}

	// 打印响应
	fmt.Printf("\n方法2 (请求头) - 状态码: %d\n", resp2.StatusCode)
	if len(body2) > 100 {
		fmt.Printf("响应体 (前100字节): %s...\n", body2[:100])
	} else {
		fmt.Printf("响应体: %s\n", body2)
	}
}
