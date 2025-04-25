package concurrency

import (
	"context"
	"fmt"
	"time"
)

// DemonstrateContext 展示Context包的使用
func DemonstrateContext() {
	// fmt.Println("1. 基本的Context取消")
	// demonstrateContextCancel()

	// fmt.Println("\n2. Context超时")
	// demonstrateContextTimeout()

	// fmt.Println("\n3. Context截止时间")
	// demonstrateContextDeadline()

	// fmt.Println("\n4. Context值传递")
	// demonstrateContextValue()

	fmt.Println("\n5. 在HTTP请求中使用Context")
	demonstrateHTTPContext()
}

// demonstrateContextCancel 展示如何使用可取消的Context
func demonstrateContextCancel() {
	// 创建一个可取消的Context
	ctx, cancel := context.WithCancel(context.Background())

	// 启动一个工作goroutine
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("工作被取消")
				return
			default:
				fmt.Println("工作进行中...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// 让工作运行一段时间
	time.Sleep(2 * time.Second)

	// 取消工作
	fmt.Println("发送取消信号")
	cancel()
	time.Sleep(1 * time.Second) // 给goroutine时间响应取消
}

// demonstrateContextTimeout 展示如何使用带超时的Context
func demonstrateContextTimeout() {
	// 创建一个2秒超时的Context
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // 良好实践：总是调用cancel，即使超时已经发生

	// 启动一个工作goroutine
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("工作超时: %v\n", ctx.Err())
				return
			default:
				fmt.Println("工作进行中...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// 等待超时发生
	<-ctx.Done()
	fmt.Println("主函数感知到超时")
	time.Sleep(1 * time.Second) // 给goroutine时间响应
}

// demonstrateContextDeadline 展示如何使用带截止时间的Context
func demonstrateContextDeadline() {
	// 创建一个有截止时间的Context
	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	// 启动一个工作goroutine
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("截止时间到达: %v\n", ctx.Err())
				return
			default:
				// 检查还有多长时间到截止时间
				if deadline, ok := ctx.Deadline(); ok {
					timeLeft := time.Until(deadline)
					fmt.Printf("距离截止时间还有: %.2f秒\n", timeLeft.Seconds())
				}
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// 等待截止时间到达
	<-ctx.Done()
	fmt.Println("主函数感知到截止时间到达")
	time.Sleep(500 * time.Millisecond) // 给goroutine时间响应
}

// 定义一些类型安全的键
type userIDKey struct{}
type authTokenKey struct{}

// demonstrateContextValue 展示如何使用Context传递值
func demonstrateContextValue() {

	// 创建带值的Context
	ctx := context.Background()
	ctx = context.WithValue(ctx, userIDKey{}, "user-123")
	ctx = context.WithValue(ctx, authTokenKey{}, "auth-token-456")

	// 启动一个使用这些值的goroutine
	go processRequest(ctx)

	// 等待goroutine完成
	time.Sleep(1 * time.Second)
}

// processRequest 模拟处理请求
func processRequest(ctx context.Context) {

	userID, ok := ctx.Value(userIDKey{}).(string)
	if !ok {
		fmt.Println("无法获取用户ID")
		return
	}

	authToken, ok := ctx.Value(authTokenKey{}).(string)
	if !ok {
		fmt.Println("无法获取认证令牌")
		return
	}

	fmt.Printf("处理用户ID为 %s 的请求，认证令牌: %s\n", userID, authToken)
}

// demonstrateHTTPContext 展示在HTTP请求中使用Context
func demonstrateHTTPContext() {
	// 在实际应用中，这会是一个HTTP处理函数
	// 这里我们只是模拟HTTP请求处理过程

	// 模拟创建一个请求Context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 模拟数据库查询
	result, err := simulateDBQuery(ctx, "SELECT * FROM users")
	if err != nil {
		fmt.Printf("数据库查询错误: %v\n", err)
		return
	}

	fmt.Printf("查询结果: %s\n", result)
	time.Sleep(5 * time.Second)
}

// simulateDBQuery 模拟一个数据库查询
func simulateDBQuery(ctx context.Context, query string) (string, error) {
	// 创建一个channel来传递结果
	resultCh := make(chan string)
	errCh := make(chan error)

	go func() {
		// 模拟查询耗时
		// time.Sleep(2 * time.Second)
		time.Sleep(4 * time.Second) // 模拟超时
		resultCh <- "用户数据"
	}()

	// 等待查询完成或Context取消
	select {
	case result := <-resultCh:
		return result, nil
	case err := <-errCh:
		return "", err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}
