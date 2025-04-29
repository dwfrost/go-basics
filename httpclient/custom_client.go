package httpclient

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// DemonstrateCustomClient 展示自定义HTTP客户端
func DemonstrateCustomClient() {
	// 创建带超时的客户端
	fmt.Println("3.1 带超时的客户端")
	clientWithTimeout()

	// 创建带代理的客户端
	fmt.Println("\n3.2 带代理的客户端 (示例代码)")
	fmt.Println("// 实际代码已注释，因为需要可用的代理服务器")

	// 创建带自定义传输的客户端
	fmt.Println("\n3.3 带自定义传输的客户端")
	clientWithCustomTransport()
}

// 创建带超时的客户端
func clientWithTimeout() {
	// 创建一个带超时的HTTP客户端
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// 发送请求
	start := time.Now()
	resp, err := client.Get("https://httpbin.org/delay/2") // 延迟2秒的接口
	duration := time.Since(start)

	if err != nil {
		fmt.Printf("请求失败: %v\n", err)
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
	fmt.Printf("请求耗时: %v\n", duration)
	fmt.Printf("状态码: %d\n", resp.StatusCode)
	if len(body) > 100 {
		fmt.Printf("响应体 (前100字节): %s...\n", body[:100])
	} else {
		fmt.Printf("响应体: %s\n", body)
	}
}

// 创建带代理的客户端 (注释掉，因为需要可用的代理服务器)
/*
func clientWithProxy() {
	// 创建代理URL
	proxyURL, err := url.Parse("http://your-proxy-server:8080")
	if err != nil {
		fmt.Printf("解析代理URL失败: %v\n", err)
		return
	}

	// 创建带代理的传输
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	// 创建使用该传输的客户端
	client := &http.Client{
		Transport: transport,
	}

	// 发送请求
	resp, err := client.Get("https://httpbin.org/ip")
	if err != nil {
		fmt.Printf("请求失败: %v\n", err)
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
*/

// 创建带自定义传输的客户端
func clientWithCustomTransport() {
	// 创建自定义传输
	transport := &http.Transport{
		MaxIdleConns:        100,              // 最大空闲连接数
		MaxIdleConnsPerHost: 10,               // 每个主机的最大空闲连接数
		IdleConnTimeout:     90 * time.Second, // 空闲连接超时
	}

	// 创建使用该传输的客户端
	client := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}

	// 发送请求
	resp, err := client.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Printf("请求失败: %v\n", err)
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
	if len(body) > 100 {
		fmt.Printf("响应体 (前100字节): %s...\n", body[:100])
	} else {
		fmt.Printf("响应体: %s\n", body)
	}

	// 打印传输配置
	fmt.Println("传输配置:")
	fmt.Printf("- 最大空闲连接数: %d\n", transport.MaxIdleConns)
	fmt.Printf("- 每个主机的最大空闲连接数: %d\n", transport.MaxIdleConnsPerHost)
	fmt.Printf("- 空闲连接超时: %v\n", transport.IdleConnTimeout)
}
