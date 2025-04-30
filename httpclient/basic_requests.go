package httpclient

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// DemonstrateBasicRequests 展示基本HTTP请求
func DemonstrateBasicRequests() {
	// GET请求示例
	fmt.Println("1.1 GET请求")
	performGETRequest()

	// POST请求示例
	fmt.Println("\n1.2 POST请求")
	performPOSTRequest()

	// 带查询参数的请求
	fmt.Println("\n1.3 带查询参数的请求")
	performRequestWithQueryParams()

	// 带自定义头的请求
	fmt.Println("\n1.4 带自定义头的请求")
	performRequestWithHeaders()
}

// 执行GET请求
func performGETRequest() {
	// 创建一个HTTP客户端
	client := &http.Client{}

	// 创建请求
	req, err := http.NewRequest("GET", "https://httpbun.com/get", nil)
	if err != nil {
		fmt.Printf("创建请求失败: %v\n", err)
		return
	}

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

	// 打印响应状态和部分内容
	fmt.Printf("状态码: %d\n", resp.StatusCode)
	fmt.Printf("响应头: %v\n", resp.Header)
	if len(body) > 200 {
		fmt.Printf("响应体 (前200字节): %s...\n", body[:200])
	} else {
		fmt.Printf("响应体: %s\n", body)
	}
}

// 执行POST请求
func performPOSTRequest() {
	// 创建请求体
	data := strings.NewReader(`{"name":"张三","age":30}`)

	// 创建请求
	resp, err := http.Post("https://httpbun.com/post", "application/json", data)
	if err != nil {
		fmt.Printf("发送POST请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return
	}

	// 打印响应状态和部分内容
	fmt.Printf("状态码: %d\n", resp.StatusCode)
	if len(body) > 200 {
		fmt.Printf("响应体 (前200字节): %s...\n", body[:200])
	} else {
		fmt.Printf("响应体: %s\n", body)
	}
}

// 执行带查询参数的请求
func performRequestWithQueryParams() {
	// 创建基础URL
	baseURL := "https://httpbun.com/get"

	// 创建查询参数
	params := url.Values{}
	params.Add("name", "张三")
	params.Add("age", "30")
	params.Add("city", "北京")

	// 将查询参数添加到URL
	requestURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// 发送请求
	resp, err := http.Get(requestURL)
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
	fmt.Printf("带查询参数的URL: %s\n", requestURL)
	fmt.Printf("状态码: %d\n", resp.StatusCode)
	if len(body) > 200 {
		fmt.Printf("响应体 (前200字节): %s...\n", body[:200])
	} else {
		fmt.Printf("响应体: %s\n", body)
	}
}

// 执行带自定义头的请求
func performRequestWithHeaders() {
	// 创建一个HTTP客户端
	client := &http.Client{}

	// 创建请求
	req, err := http.NewRequest("GET", "https://httpbun.com/headers", nil)
	if err != nil {
		fmt.Printf("创建请求失败: %v\n", err)
		return
	}

	// 添加自定义头
	req.Header.Add("X-Custom-Header", "自定义值")
	req.Header.Add("User-Agent", "Go-HTTP-Client/1.0")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")

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
