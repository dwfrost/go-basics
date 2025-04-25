package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// DemonstratePracticalConcurrency 展示并发的实际应用
func DemonstratePracticalConcurrency() {
	fmt.Println("1. 并行计算示例")
	numbers := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	result := parallelSum(numbers)
	fmt.Printf("并行计算总和: %d\n", result)

	// fmt.Println("\n2. 工作池模式")
	// demonstrateWorkerPool()

	// fmt.Println("\n3. 使用WaitGroup同步")
	// demonstrateWaitGroup()

	// fmt.Println("\n4. 使用互斥锁")
	// demonstrateMutex()

	fmt.Println("\n5. 使用Context控制取消")
	demonstrateContext()
}

// parallelSum 并行计算切片中所有数字的平方和
func parallelSum(numbers []int) int {
	// 创建一个channel来接收结果
	results := make(chan int)

	// 启动多个goroutine计算部分和
	for _, num := range numbers {
		go func(n int) {
			// 计算平方并发送到channel
			results <- n * n
		}(num)
	}

	// 收集所有结果
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += <-results
	}

	return sum
}

// demonstrateWorkerPool 展示工作池模式
func demonstrateWorkerPool() {
	const numJobs = 10
	const numWorkers = 3

	// 创建任务和结果channel
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 启动工作者
	for w := 1; w <= numWorkers; w++ {
		go worker2(w, jobs, results)
	}

	// 发送任务
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// 收集结果
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

// worker2 是工作池中的工作者
func worker2(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("工作者 %d 开始处理任务 %d\n", id, j)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		fmt.Printf("工作者 %d 完成任务 %d\n", id, j)
		results <- j * 2
	}
}

// demonstrateWaitGroup 展示如何使用WaitGroup
func demonstrateWaitGroup() {
	var wg sync.WaitGroup

	// 添加3个等待项
	wg.Add(3)

	// 启动3个goroutine
	for i := 1; i <= 3; i++ {
		go func(id int) {
			defer wg.Done() // 完成时减少计数
			fmt.Printf("Goroutine %d 开始工作\n", id)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Printf("Goroutine %d 完成工作\n", id)
		}(i)
	}

	// 等待所有goroutine完成
	wg.Wait()
	fmt.Println("所有goroutine已完成")
}

// demonstrateMutex 展示互斥锁的使用
func demonstrateMutex() {
	// 共享变量
	var counter int
	// 互斥锁保护共享变量
	var mu sync.Mutex
	var wg sync.WaitGroup

	// 启动1000个goroutine增加计数器
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 加锁
			mu.Lock()
			counter++
			// 解锁
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Printf("计数器最终值: %d\n", counter)
}

// demonstrateContext 展示Context的使用
func demonstrateContext() {
	// 在实际应用中，我们会使用context包
	// 这里用一个简化版本展示概念
	done := make(chan struct{})

	// 启动可取消的操作
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("操作被取消")
				return
			default:
				fmt.Println("操作进行中...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// 让操作运行一段时间
	time.Sleep(2 * time.Second)

	// 发送取消信号
	close(done)
	time.Sleep(500 * time.Millisecond) // 给goroutine时间响应取消
}
