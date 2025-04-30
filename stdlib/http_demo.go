package stdlib

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// DemonstrateHTTP 展示net/http包的使用
func DemonstrateHTTP() {
	// HTTP客户端 - 发送GET请求
	fmt.Println("4.1 HTTP客户端 - GET请求")
	fmt.Println("发送GET请求到 https://httpbun.com/get")
	resp, err := http.Get("https://httpbun.com/get")
	if err == nil {
		defer resp.Body.Close()
		fmt.Printf("状态码: %d\n", resp.StatusCode)
		fmt.Printf("内容类型: %s\n", resp.Header.Get("Content-Type"))

		// 读取前100个字符的响应
		body, _ := io.ReadAll(resp.Body)
		if len(body) > 100 {
			fmt.Printf("响应体 (前100字符): %s...\n", body[:100])
		} else {
			fmt.Printf("响应体: %s\n", body)
		}
	} else {
		fmt.Printf("请求失败: %v\n", err)
	}

	// 自定义HTTP客户端
	fmt.Println("\n4.2 自定义HTTP客户端")
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, _ := http.NewRequest("GET", "https://httpbun.com/headers", nil)
	req.Header.Add("User-Agent", "Go-HTTP-Client/1.1")
	req.Header.Add("Custom-Header", "CustomValue")

	fmt.Println("发送带自定义头的请求到 https://httpbun.com/headers")
	resp, err = client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		if len(body) > 100 {
			fmt.Printf("响应体 (前100字符): %s...\n", body[:100])
		} else {
			fmt.Printf("响应体: %s\n", body)
		}
	}

	// HTTP服务器 (注释掉以避免阻塞)
	fmt.Println("\n4.3 HTTP服务器 (代码已注释)")
	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, 你正在访问: %s\n", r.URL.Path)
		})

		http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, `{"message": "这是一个API端点", "status": "success"}`)
		})

		fmt.Println("启动HTTP服务器在 http://localhost:8080/")
		http.ListenAndServe(":8080", nil)
	*/
}
