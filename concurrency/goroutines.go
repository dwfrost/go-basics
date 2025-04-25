package concurrency

import (
	"fmt"
	"time"
)

// DemonstrateGoroutines 展示Goroutine的基本用法
func DemonstrateGoroutines() {
	fmt.Println("1. 基本Goroutine")
	// 启动一个Goroutine
	go sayHello("世界")
	// 主函数继续执行
	fmt.Println("主函数继续执行")
	// 等待一秒，否则程序可能在Goroutine执行前就结束了
	time.Sleep(time.Second)

	fmt.Println("\n2. 多个Goroutine")
	// 启动多个Goroutine
	go count("羊", 5)
	go count("鸭", 5)
	// 等待足够长的时间让Goroutines完成
	time.Sleep(3 * time.Second)

	fmt.Println("\n3. 匿名Goroutine")
	// 启动一个匿名函数作为Goroutine
	go func() {
		fmt.Println("这是一个匿名Goroutine")
	}()
	time.Sleep(time.Second)
}

// sayHello 是一个简单的函数，将被作为Goroutine运行
func sayHello(name string) {
	fmt.Printf("你好，%s!\n", name)
}

// count 函数会打印name指定次数
func count(name string, n int) {
	for i := 1; i <= n; i++ {
		fmt.Printf("%s %d\n", name, i)
		// 让出CPU时间，让其他Goroutine有机会运行
		time.Sleep(500 * time.Millisecond)
	}
}
