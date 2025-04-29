package httpclient

import (
	"fmt"
	"time"
)

// DemonstrateHTTPClient 展示HTTP客户端的使用
func DemonstrateHTTPClient() {
	fmt.Println("=== HTTP客户端学习 ===")

	fmt.Println("\n1. 基本HTTP请求")
	DemonstrateBasicRequests()

	fmt.Println("\n2. 处理JSON响应")
	DemonstrateJSONHandling()

	fmt.Println("\n3. 自定义HTTP客户端")
	DemonstrateCustomClient()

	fmt.Println("\n4. 处理认证")
	DemonstrateAuthentication()

	fmt.Println("\n5. 调用第三方API示例")
	DemonstrateThirdPartyAPIs()

	fmt.Println("\n6. 并发请求")
	DemonstrateConcurrentRequests()

	fmt.Println("\n7. 错误处理和重试")
	DemonstrateErrorHandling()
}

// 设置超时时间常量
const (
	DefaultTimeout      = 10 * time.Second
	DefaultRetryCount   = 3
	DefaultRetryBackoff = 1 * time.Second
)
