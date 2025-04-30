package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// DemonstrateJSONHandling 展示JSON处理
func DemonstrateJSONHandling() {
	// 发送JSON数据
	fmt.Println("2.1 发送JSON数据")
	sendJSONData()

	// 解析JSON响应
	fmt.Println("\n2.2 解析JSON响应")
	parseJSONResponse()

	// 使用结构体处理JSON
	fmt.Println("\n2.3 使用结构体处理JSON")
	useStructsWithJSON()
}

// 用户数据结构
type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email,omitempty"`
	Address struct {
		City    string `json:"city"`
		Country string `json:"country"`
	} `json:"address"`
}

// 发送JSON数据
func sendJSONData() {
	// 创建用户数据
	user := User{
		Name: "李四",
		Age:  35,
		Address: struct {
			City    string `json:"city"`
			Country string `json:"country"`
		}{
			City:    "上海",
			Country: "中国",
		},
	}

	// 将结构体转换为JSON
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("JSON编码失败: %v\n", err)
		return
	}

	// 创建请求
	req, err := http.NewRequest("POST", "https://httpbun.com/anything", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("创建请求失败: %v\n", err)
		return
	}

	// 设置内容类型
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
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
	fmt.Printf("发送的JSON数据: %s\n", jsonData)
	fmt.Printf("状态码: %d\n", resp.StatusCode)
	if len(body) > 200 {
		fmt.Printf("响应体 (前200字节): %s...\n", body[:200])
	} else {
		fmt.Printf("响应体: %s\n", body)
	}
}

// 解析JSON响应
func parseJSONResponse() {
	// 发送请求
	resp, err := http.Get("https://httpbun.com/anything")
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

	// 解析JSON
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Printf("JSON解析失败: %v\n", err)
		return
	}

	// 打印解析结果
	fmt.Println("解析的JSON数据:")
	for key, value := range result {
		fmt.Printf("- %s: %v\n", key, value)
	}
}

// 使用结构体处理JSON
func useStructsWithJSON() {
	// 定义响应结构
	type HTTPResponse struct {
		Method  string            `json:"method"`
		URL     string            `json:"url"`
		Headers map[string]string `json:"headers"`
		Origin  string            `json:"origin"`
		Data    string            `json:"data,omitempty"`
	}

	// 创建请求
	req, err := http.NewRequest("GET", "https://httpbun.com/anything", nil)
	if err != nil {
		fmt.Printf("创建请求失败: %v\n", err)
		return
	}

	// 添加自定义请求头
	req.Header.Add("X-Custom-Header", "Hello")
	req.Header.Add("User-Agent", "Go-HTTP-Client")

	// 发送请求
	client := &http.Client{}
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

	// 解析JSON到结构体
	var response HTTPResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("JSON解析失败: %v\n", err)
		return
	}

	// 打印结构体数据
	fmt.Println("使用结构体解析的JSON数据:")
	fmt.Printf("- 请求方法: %s\n", response.Method)
	fmt.Printf("- 请求URL: %s\n", response.URL)
	fmt.Printf("- 来源IP: %s\n", response.Origin)
	fmt.Printf("- 返回data: %s\n", response.Data)
	fmt.Println("- 请求头:")
	for key, value := range response.Headers {
		fmt.Printf("  %s: %s\n", key, value)
	}
}
