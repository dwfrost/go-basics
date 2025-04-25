package stdlib

import (
	"fmt"
	"time"
)

// DemonstrateTime 展示time包的使用
func DemonstrateTime() {
	// 获取当前时间
	fmt.Println("5.1 获取当前时间")
	now := time.Now()
	fmt.Printf("当前时间: %v\n", now)
	fmt.Printf("Unix时间戳: %d\n", now.Unix())
	fmt.Printf("纳秒时间戳: %d\n", now.UnixNano())

	// 时间格式化
	fmt.Println("\n5.2 时间格式化")
	fmt.Printf("标准格式: %s\n", now.Format(time.RFC3339))
	fmt.Printf("自定义格式: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("日期: %s\n", now.Format("2006年01月02日"))
	fmt.Printf("时间: %s\n", now.Format("15:04:05"))

	// 时间解析
	fmt.Println("\n5.3 时间解析")
	t, _ := time.Parse("2006-01-02", "2023-05-15")
	fmt.Printf("解析的时间: %v\n", t)

	// 时间计算
	fmt.Println("\n5.4 时间计算")
	future := now.Add(24 * time.Hour)
	fmt.Printf("明天这个时间: %v\n", future)

	past := now.AddDate(0, 0, -7)
	fmt.Printf("一周前: %v\n", past)

	diff := future.Sub(now)
	fmt.Printf("时间差: %v\n", diff)

	// 时间比较
	fmt.Println("\n5.5 时间比较")
	fmt.Printf("future在now之后: %t\n", future.After(now))
	fmt.Printf("past在now之前: %t\n", past.Before(now))
	fmt.Printf("创建的时间t和now相等: %t\n", t.Equal(now))

	// 定时器和打点器 (注释掉以避免阻塞)
	fmt.Println("\n5.6 定时器和打点器 (代码已注释)")
	/*
		// 定时器 - 等待一段时间后执行
		timer := time.NewTimer(2 * time.Second)
		fmt.Println("定时器启动...")
		<-timer.C
		fmt.Println("定时器触发!")

		// 打点器 - 定期执行
		ticker := time.NewTicker(1 * time.Second)
		count := 0
		for {
			<-ticker.C
			count++
			fmt.Println("滴答!")
			if count >= 5 {
				ticker.Stop()
				break
			}
		}
	*/

	// 时区
	fmt.Println("\n5.7 时区")
	local := time.Now()
	fmt.Printf("本地时间: %v\n", local)

	utc := local.UTC()
	fmt.Printf("UTC时间: %v\n", utc)

	// 加载特定时区
	location, err := time.LoadLocation("America/New_York")
	if err == nil {
		nyTime := local.In(location)
		fmt.Printf("纽约时间: %v\n", nyTime)
	}
}
