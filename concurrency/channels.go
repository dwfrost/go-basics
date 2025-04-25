package concurrency

import (
	"fmt"
	"time"
)

// DemonstrateChannels 展示Channel的基本用法
func DemonstrateChannels() {
	fmt.Println("1. 基本Channel")
	// 创建一个无缓冲的字符串channel
	ch := make(chan string)

	// 在Goroutine中发送数据
	go func() {
		ch <- "你好，Channel!"
	}()

	// 从channel接收数据
	msg := <-ch
	fmt.Println(msg)

	fmt.Println("\n2. 带缓冲的Channel")
	// 创建一个容量为3的缓冲channel
	bufferedCh := make(chan int, 3)

	// 发送数据到缓冲channel
	bufferedCh <- 1
	bufferedCh <- 2
	bufferedCh <- 3
	// 此时缓冲区已满，再发送会阻塞

	// 接收数据
	fmt.Println(<-bufferedCh)
	fmt.Println(<-bufferedCh)
	fmt.Println(<-bufferedCh)

	fmt.Println("\n3. 使用Channel进行同步")
	done := make(chan bool)

	go worker(done)

	// 等待worker完成
	<-done
	fmt.Println("主函数结束")

	fmt.Println("\n4. 单向Channel")
	demonstrateDirectionalChannels()

	fmt.Println("\n5. 使用select多路复用")
	demonstrateSelect()

	fmt.Println("\n6. 关闭Channel和遍历")
	demonstrateCloseAndRange()
}

// worker 模拟一个工作Goroutine
func worker(done chan bool) {
	fmt.Println("工作中...")
	time.Sleep(2 * time.Second)
	fmt.Println("工作完成!")
	// 发送完成信号
	done <- true
}

// demonstrateDirectionalChannels 展示单向Channel
func demonstrateDirectionalChannels() {
	// 创建一个双向channel
	ch := make(chan int)

	// 启动发送者和接收者
	go sender(ch)
	go receiver(ch)

	// 等待足够时间让goroutines完成
	time.Sleep(2 * time.Second)
}

// sender 只向channel发送数据
func sender(ch chan<- int) { // chan<- 表示只发送channel
	for i := 1; i <= 5; i++ {
		ch <- i
		time.Sleep(200 * time.Millisecond)
	}
}

// receiver 只从channel接收数据
func receiver(ch <-chan int) { // <-chan 表示只接收channel
	for i := 1; i <= 5; i++ {
		val := <-ch
		fmt.Printf("接收到: %d\n", val)
	}
}

// demonstrateSelect 展示select语句的使用
func demonstrateSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 在两个goroutine中发送数据
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "来自channel 1的消息"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "来自channel 2的消息"
	}()

	// 使用select等待两个channel
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}

// demonstrateCloseAndRange 展示关闭channel和使用range遍历
func demonstrateCloseAndRange() {
	ch := make(chan int, 5)

	// 发送数据并关闭channel
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch) // 关闭channel
	}()

	// 使用range遍历channel直到它关闭
	for num := range ch {
		fmt.Printf("接收到: %d\n", num)
	}
}
