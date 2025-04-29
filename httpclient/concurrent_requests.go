package httpclient

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// DemonstrateConcurrentRequests 展示并发请求
func DemonstrateConcurrentRequests() {
	// 基本并发请求
	fmt.Println("6.1 基本并发请求")
	basicConcurrentRequests()

	// 使用工作池
	fmt.Println("\n6.2 使用工作池")
	requestsWithWorkerPool()

	// 限制并发数
	fmt.Println("\n6.3 限制并发数")
	requestsWithRateLimiting()
}

// 基本并发请求
func basicConcurrentRequests() {
	// 创建要请求的URL列表
	urls := []string{
		"https://httpbin.org/get",
		"https://httpbin.org/ip",
		"https://httpbin.org/headers",
		"https://httpbin.org/user-agent",
		"https://httpbin.org/uuid",
	}

	// 创建等待组
	var wg sync.WaitGroup
	wg.Add(len(urls))

	// 记录开始时间
	start := time.Now()

	// 为每个URL创建一个goroutine
	for i, url := range urls {
		// 启动goroutine
		go func(i int, url string) {
			defer wg.Done()

			// 发送请求
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("请求 %s 失败: %v\n", url, err)
				return
			}
			defer resp.Body.Close()

			// 读取响应
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("读取响应失败: %v\n", err)
				return
			}

			// 打印结果
			fmt.Printf("请求 %d (%s) - 状态码: %d, 响应大小: %d 字节\n",
				i+1, url, resp.StatusCode, len(body))
		}(i, url)
	}

	// 等待所有goroutine完成
	wg.Wait()

	// 计算总耗时
	duration := time.Since(start)
	fmt.Printf("并发请求总耗时: %v\n", duration)
}

// 使用工作池
func requestsWithWorkerPool() {
	// 创建要请求的URL列表
	urls := []string{
		"https://httpbin.org/get",
		"https://httpbin.org/ip",
		"https://httpbin.org/headers",
		"https://httpbin.org/user-agent",
		"https://httpbin.org/uuid",
		"https://httpbin.org/delay/1",
		"https://httpbin.org/delay/2",
		"https://httpbin.org/delay/1",
		"https://httpbin.org/delay/2",
		"https://httpbin.org/delay/1",
	}

	// 设置工作池大小
	workerCount := 3
	fmt.Printf("使用 %d 个工作线程处理 %d 个请求\n", workerCount, len(urls))

	// 创建任务通道
	tasks := make(chan string, len(urls))
	results := make(chan string, len(urls))

	// 启动工作线程
	var wg sync.WaitGroup
	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go func(workerId int) {
			defer wg.Done()
			for url := range tasks {
				// 发送请求
				start := time.Now()
				resp, err := http.Get(url)
				if err != nil {
					results <- fmt.Sprintf("工作线程 %d - 请求 %s 失败: %v", workerId, url, err)
					continue
				}

				// 读取响应
				body, err := io.ReadAll(resp.Body)
				resp.Body.Close()
				if err != nil {
					results <- fmt.Sprintf("工作线程 %d - 读取响应失败: %v", workerId, err)
					continue
				}

				// 发送结果
				duration := time.Since(start)
				results <- fmt.Sprintf("工作线程 %d - 请求 %s - 状态码: %d, 响应大小: %d 字节, 耗时: %v",
					workerId, url, resp.StatusCode, len(body), duration)
			}
		}(i + 1)
	}

	// 记录开始时间
	start := time.Now()

	// 发送任务
	for _, url := range urls {
		tasks <- url
	}
	close(tasks)

	// 等待所有工作线程完成
	go func() {
		wg.Wait()
		close(results)
	}()

	// 收集结果
	for result := range results {
		fmt.Println(result)
	}

	// 计算总耗时
	duration := time.Since(start)
	fmt.Printf("工作池总耗时: %v\n", duration)
}

// 限制并发数
func requestsWithRateLimiting() {
	// 创建要请求的URL列表
	urls := []string{
		"https://httpbin.org/get",
		"https://httpbin.org/ip",
		"https://httpbin.org/headers",
		"https://httpbin.org/user-agent",
		"https://httpbin.org/uuid",
		"https://httpbin.org/delay/1",
		"https://httpbin.org/delay/1",
		"https://httpbin.org/delay/1",
		"https://httpbin.org/delay/1",
		"https://httpbin.org/delay/1",
	}

	// 设置限流参数
	maxConcurrent := 2                        // 最大并发数
	requestInterval := 500 * time.Millisecond // 请求间隔

	fmt.Printf("限制最大并发数为 %d，请求间隔为 %v\n", maxConcurrent, requestInterval)

	// 创建信号量通道
	semaphore := make(chan struct{}, maxConcurrent)
	var wg sync.WaitGroup

	// 记录开始时间
	start := time.Now()

	// 发送请求
	for i, url := range urls {
		wg.Add(1)

		// 限制请求速率
		time.Sleep(requestInterval)

		go func(i int, url string) {
			defer wg.Done()

			// 获取信号量
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			// 发送请求
			requestStart := time.Now()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("请求 %d (%s) 失败: %v\n", i+1, url, err)
				return
			}
			defer resp.Body.Close()

			// 读取响应
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("读取响应失败: %v\n", err)
				return
			}

			// 计算请求耗时
			requestDuration := time.Since(requestStart)

			// 打印结果
			fmt.Printf("请求 %d (%s) - 状态码: %d, 响应大小: %d 字节, 耗时: %v\n",
				i+1, url, resp.StatusCode, len(body), requestDuration)
		}(i, url)
	}

	// 等待所有请求完成
	wg.Wait()

	// 计算总耗时
	duration := time.Since(start)
	fmt.Printf("限流请求总耗时: %v\n", duration)
}
