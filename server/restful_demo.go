package server

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Product 产品结构体
type Product struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Price       float64   `json:"price" binding:"required,gt=0"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ProductResponse 产品响应结构体
type ProductResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// DemoRESTful 展示RESTful API的设计与实现
func DemoRESTful() {
	// 创建路由引擎
	r := gin.Default()

	// 添加自定义响应中间件
	r.Use(ResponseMiddleware())

	// API版本控制
	v1 := r.Group("/api/v1")
	{
		// 产品相关路由
		products := v1.Group("/products")
		{
			products.GET("", ListProducts)          // 获取产品列表
			products.GET("/:id", GetProduct)        // 获取单个产品
			products.POST("", CreateProduct)        // 创建产品
			products.PUT("/:id", UpdateProduct)     // 更新产品
			products.DELETE("/:id", DeleteProduct)  // 删除产品
			products.GET("/search", SearchProducts) // 搜索产品
		}
	}

	// API文档路由
	r.GET("/docs/*any", gin.WrapH(swaggerHandler()))

	fmt.Println("=== RESTful API 示例 ===")
	fmt.Println("服务器运行在 http://localhost:8080")
	fmt.Println("\n可用的API端点：")
	fmt.Println("1. GET    /api/v1/products     - 获取产品列表")
	fmt.Println("2. GET    /api/v1/products/:id - 获取单个产品")
	fmt.Println("3. POST   /api/v1/products     - 创建产品")
	fmt.Println("4. PUT    /api/v1/products/:id - 更新产品")
	fmt.Println("5. DELETE /api/v1/products/:id - 删除产品")
	fmt.Println("6. GET    /api/v1/products/search?q=关键词 - 搜索产品")
	fmt.Println("\n文档地址：http://localhost:8080/docs/")

	r.Run(":8080")
}

// 模拟数据库
var products = []Product{
	{
		ID:          1,
		Name:        "Go编程实战",
		Description: "深入学习Go语言的实践指南",
		Price:       99.00,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}

// ResponseMiddleware 统一响应格式中间件
func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 如果已经有响应，则不处理
		if c.Writer.Written() {
			return
		}

		// 获取处理结果
		data, exists := c.Get("data")
		if !exists {
			data = gin.H{}
		}

		// 构造响应
		response := ProductResponse{
			Code:    c.Writer.Status(),
			Message: http.StatusText(c.Writer.Status()),
			Data:    data,
		}

		c.JSON(c.Writer.Status(), response)
	}
}

// ListProducts 获取产品列表
func ListProducts(c *gin.Context) {
	// 分页参数
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	// TODO: 实现分页逻辑

	c.Set("data", gin.H{
		"products": products,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": len(products),
		},
	})
}

// GetProduct 获取单个产品
func GetProduct(c *gin.Context) {
	id := c.Param("id")
	for _, product := range products {
		if fmt.Sprint(product.ID) == id {
			c.Set("data", product)
			return
		}
	}
	c.Status(http.StatusNotFound)
}

// CreateProduct 创建产品
func CreateProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 设置创建时间和更新时间
	product.ID = uint(len(products) + 1)
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	products = append(products, product)
	c.Status(http.StatusCreated)
	c.Set("data", product)
}

// UpdateProduct 更新产品
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, p := range products {
		if fmt.Sprint(p.ID) == id {
			product.ID = p.ID
			product.CreatedAt = p.CreatedAt
			product.UpdatedAt = time.Now()
			products[i] = product
			c.Set("data", product)
			return
		}
	}
	c.Status(http.StatusNotFound)
}

// DeleteProduct 删除产品
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	for i, product := range products {
		if fmt.Sprint(product.ID) == id {
			products = append(products[:i], products[i+1:]...)
			c.Status(http.StatusOK)
			return
		}
	}
	c.Status(http.StatusNotFound)
}

// SearchProducts 搜索产品
func SearchProducts(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.Set("data", products)
		return
	}

	// 简单的搜索实现
	var results []Product
	for _, product := range products {
		if strings.Contains(strings.ToLower(product.Name), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(product.Description), strings.ToLower(query)) {
			results = append(results, product)
		}
	}

	c.Set("data", results)
}

// swaggerHandler 返回Swagger UI处理器
func swaggerHandler() http.Handler {
	// TODO: 实现Swagger UI
	return http.NotFoundHandler()
}
