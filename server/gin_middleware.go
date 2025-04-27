package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// DemonstrateGin 展示Gin框架的中间件使用
func DemonstrateGin() {
	// 创建 Gin 引擎实例
	r := gin.Default()

	// 1. 全局中间件
	r.Use(LoggerMiddleware())
	r.Use(CORSMiddleware())

	// 2. 路由组中间件
	authorized := r.Group("/auth")
	authorized.Use(AuthMiddleware())
	{
		authorized.GET("/profile", func(c *gin.Context) {
			// 从上下文中获取用户信息
			user := c.MustGet("user").(string)
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("你好, %s", user),
				"status":  "已认证",
			})
		})
	}

	// 3. 单个路由中间件
	r.GET("/test", RequestTimeMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "这是一个测试端点",
		})
	})

	// 4. 基础路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "欢迎使用 Gin 框架",
		})
	})

	// 启动服务器
	fmt.Println("Gin服务器启动在 http://localhost:8080")
	fmt.Println("可以尝试访问以下URL：")
	fmt.Println("1. http://localhost:8080/ (首页)")
	fmt.Println("2. http://localhost:8080/test (测试延迟中间件)")
	fmt.Println("3. http://localhost:8080/auth/profile?token=valid (需要认证)")
	fmt.Println("4. http://localhost:8080/auth/profile (无token将被拒绝)")
	fmt.Println("按 Ctrl+C 停止服务器")

	r.Run(":8080")
}

// LoggerMiddleware 日志中间件
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求前
		startTime := time.Now()
		path := c.Request.URL.Path

		// 处理请求
		c.Next()

		// 请求后
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		// 获取状态码
		statusCode := c.Writer.Status()

		log.Printf("[GIN] %s | %d | %v | %s",
			c.Request.Method,
			statusCode,
			latency,
			path,
		)
	}
}

// AuthMiddleware 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token != "valid" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "未授权访问",
			})
			c.Abort() // 终止后续中间件的执行
			return
		}

		// 设置用户信息到上下文
		c.Set("user", "测试用户")
		c.Next()
	}
}

// CORSMiddleware 跨域中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// RequestTimeMiddleware 请求时间中间件
func RequestTimeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 模拟处理延迟
		time.Sleep(200 * time.Millisecond)
		c.Next()
	}
}
