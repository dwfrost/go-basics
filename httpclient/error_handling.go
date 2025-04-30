package httpclient

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"time"
)

// DemonstrateErrorHandling 展示错误处理和重试
func DemonstrateErrorHandling() {
	// 基本错误处理
	fmt.Println("7.1 基本错误处理")
	basicErrorHandling()

	// 超时处理
	fmt.Println("\n7.2 超时处理")
	timeoutHandling()

	// 重试机制
	fmt.Println("\n7.3 重试机制")
	retryMechanism()

	// 断路器模式
	fmt.Println("\n7.4 断路器模式 (概念示例)")
	fmt.Println("断路器模式是一种保护系统免受级联故障影响的设计模式。")
	fmt.Println("当系统检测到持续的错误时，断路器会'断开'，阻止进一步的请求。")
	fmt.Println("在一段时间后，断路器会进入'半开'状态，允许少量请求通过。")
	fmt.Println("如果这些请求成功，断路器会'闭合'，恢复正常操作。")
}

// 基本错误处理
func basicErrorHandling() {
	// 创建一个不存在的URL
	url := "https://httpbun.com/status/404"

	// 发送请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("请求返回非成功状态码: %d\n", resp.StatusCode)

		// 根据状态码处理不同情况
		switch resp.StatusCode {
		case http.StatusNotFound:
			fmt.Println("资源不存在 (404)")
		case http.StatusUnauthorized:
			fmt.Println("未授权 (401)")
		case http.StatusForbidden:
			fmt.Println("禁止访问 (403)")
		case http.StatusInternalServerError:
			fmt.Println("服务器内部错误 (500)")
		default:
			fmt.Printf("其他错误状态码: %d\n", resp.StatusCode)
		}

		// 读取错误响应
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("读取错误响应失败: %v\n", err)
			return
		}

		fmt.Printf("错误响应内容: %s\n", body)
		return
	}

	// 读取成功响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return
	}

	fmt.Printf("成功响应: %s\n", body)
}

// 超时处理
func timeoutHandling() {
	// 创建上下文
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 创建请求
	req, err := http.NewRequestWithContext(ctx, "GET", "https://httpbun.com/delay/3", nil)
	if err != nil {
		fmt.Printf("创建请求失败: %v\n", err)
		return
	}

	// 发送请求
	start := time.Now()
	resp, err := http.DefaultClient.Do(req)
	duration := time.Since(start)

	// 检查是否超时
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Printf("请求超时 (耗时: %v): %v\n", duration, err)
		} else {
			fmt.Printf("请求失败: %v\n", err)
		}
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return
	}

	fmt.Printf("请求成功 (耗时: %v), 状态码: %d, 响应大小: %d 字节\n",
		duration, resp.StatusCode, len(body))
}

// 重试机制
func retryMechanism() {
	// 设置重试参数
	maxRetries := 3
	baseDelay := 1 * time.Second

	// 创建一个不稳定的URL (模拟随机失败)
	urls := []string{
		"https://httpbun.com/status/200",
		"https://httpbun.com/status/500",
		"https://httpbun.com/status/503",
	}

	// 随机选择一个URL
	url := urls[rand.Intn(len(urls))]

	fmt.Printf("请求URL: %s (可能随机失败)\n", url)

	// 创建HTTP客户端
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// 执行带重试的请求
	var (
		resp    *http.Response
		err     error
		attempt int
	)

	for attempt = 0; attempt <= maxRetries; attempt++ {
		if attempt > 0 {
			// 计算退避时间 (指数退避)
			delay := time.Duration(math.Pow(2, float64(attempt-1))) * baseDelay

			// 添加抖动 (随机因素)
			jitter := time.Duration(rand.Int63n(int64(baseDelay)))
			delay += jitter

			fmt.Printf("尝试 %d 失败，等待 %v 后重试...\n", attempt, delay)
			time.Sleep(delay)
		}

		// 发送请求
		resp, err = client.Get(url)

		// 检查是否成功
		if err == nil && resp.StatusCode == http.StatusOK {
			break
		}

		// 处理错误
		if err != nil {
			fmt.Printf("尝试 %d 请求错误: %v\n", attempt+1, err)
		} else {
			fmt.Printf("尝试 %d 状态码错误: %d\n", attempt+1, resp.StatusCode)
			resp.Body.Close()
		}
	}

	// 检查最终结果
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Printf("在 %d 次尝试后请求失败\n", attempt)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return
	}

	fmt.Printf("在尝试 %d 次后请求成功，状态码: %d, 响应大小: %d 字节\n",
		attempt+1, resp.StatusCode, len(body))
}
