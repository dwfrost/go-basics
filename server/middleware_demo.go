package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// DemonstrateMiddleware 展示中间件的使用
func DemonstrateMiddleware() {
	// 创建路由
	mux := http.NewServeMux()

	// 注册处理函数，使用中间件链
	mux.HandleFunc("/", Chain(
		homeHandler,
		loggingMiddleware,
		authMiddleware,
		timingMiddleware,
	))

	// 启动服务器
	fmt.Println("服务器启动在 http://localhost:8080")
	fmt.Println("可以尝试访问以下URL：")
	fmt.Println("1. http://localhost:8080/ (无token)")
	fmt.Println("2. http://localhost:8080/?token=valid (带有效token)")
	fmt.Println("按 Ctrl+C 停止服务器")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

// 处理函数类型
type HandlerFunc func(http.ResponseWriter, *http.Request)

// 中间件类型
type Middleware func(HandlerFunc) HandlerFunc

// Chain 将多个中间件连接成一个链
func Chain(handler HandlerFunc, middlewares ...Middleware) HandlerFunc {
	for _, m := range middlewares {
		handler = m(handler)
	}
	return handler
}

// 基础处理函数
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!\n")
}

// 日志中间件
func loggingMiddleware(next HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 请求前的日志
		log.Printf("开始处理请求: %s %s", r.Method, r.URL.Path)

		// 调用下一个处理函数
		next(w, r)

		// 请求后的日志
		log.Printf("请求处理完成: %s %s", r.Method, r.URL.Path)
	}
}

// 认证中间件
func authMiddleware(next HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 检查token
		token := r.URL.Query().Get("token")
		if token != "valid" {
			http.Error(w, "未授权访问", http.StatusUnauthorized)
			return
		}

		// 验证通过，继续处理
		next(w, r)
	}
}

// 计时中间件
func timingMiddleware(next HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// 调用下一个处理函数
		next(w, r)

		// 计算处理时间
		duration := time.Since(start)
		log.Printf("请求处理耗时: %v", duration)
	}
}
