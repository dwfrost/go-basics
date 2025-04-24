package constants

import "fmt"

// DemonstrateConstants 展示Go语言中常量的声明和使用
func DemonstrateConstants() {
	// 1. 基本常量声明
	const PI = 3.14159
	fmt.Println("PI值:", PI)

	// 2. 声明多个常量
	const (
		StatusOK       = 200
		StatusNotFound = 404
	)
	fmt.Println("HTTP状态码:", StatusOK, StatusNotFound)

	// 3. 类型常量
	const MaxUsers int = 1000
	fmt.Println("最大用户数:", MaxUsers)

	// 4. iota 常量生成器
	const (
		Sunday    = iota // 0
		Monday           // 1
		Tuesday          // 2
		Wednesday        // 3
		Thursday         // 4
		Friday           // 5
		Saturday         // 6
	)
	fmt.Println("星期常量:", Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	// 5. iota 在同一行
	const (
		_  = iota             // 0 (忽略)
		KB = 1 << (10 * iota) // 1 << 10 = 1024
		MB                    // 1 << 20 = 1048576
		GB                    // 1 << 30 = 1073741824
	)
	fmt.Printf("存储单位: KB=%d, MB=%d, GB=%d\n", KB, MB, GB)

	// 6. 常量表达式
	const (
		Year    = 365
		Decade  = Year * 10
		Century = Decade * 10
	)
	fmt.Println("时间常量: 年=", Year, "十年=", Decade, "世纪=", Century)
}
