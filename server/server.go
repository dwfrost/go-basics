package server

import "fmt"

// DemonstrateServer 展示Go服务器相关功能
func DemonstrateServer() {
	fmt.Println("=== Go Http服务器学习 ===")

	fmt.Println("\n1. 中间件示例")
	// DemonstrateMiddleware()
	// DemonstrateGin()

	fmt.Println("\n2. RESTful API示例")
	DemoRESTful()
}
