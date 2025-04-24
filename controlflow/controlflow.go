package controlflow

import (
	"fmt"
	"time"
)

// DemonstrateControlFlow 展示Go语言中的控制流结构
func DemonstrateControlFlow() {
	// 1. 条件语句
	demonstrateConditionals()
	
	// 2. 循环语句
	demonstrateLoops()
	
	// 3. 跳转语句
	demonstrateJumpStatements()
	
	// 4. 延迟执行
	demonstrateDefer()
}

// demonstrateConditionals 展示条件语句
func demonstrateConditionals() {
	fmt.Println("--- 条件语句 ---")
	
	// 基本if语句
	x := 10
	if x > 5 {
		fmt.Println("x大于5")
	}
	
	// if-else语句
	if x > 20 {
		fmt.Println("x大于20")
	} else {
		fmt.Println("x不大于20")
	}
	
	// if-else if-else语句
	if x > 20 {
		fmt.Println("x大于20")
	} else if x > 10 {
		fmt.Println("x大于10但不大于20")
	} else {
		fmt.Println("x不大于10")
	}
	
	// if语句中的短变量声明
	if y := 15; y > x {
		fmt.Printf("y(%d)大于x(%d)\n", y, x)
	} else {
		fmt.Printf("y(%d)不大于x(%d)\n", y, x)
	}
	
	// switch语句
	fmt.Println("\nswitch语句:")
	
	day := time.Now().Weekday()
	switch day {
	case time.Saturday, time.Sunday:
		fmt.Println("今天是周末")
	default:
		fmt.Println("今天是工作日")
	}
	
	// 不带表达式的switch
	hour := time.Now().Hour()
	switch {
	case hour < 12:
		fmt.Println("现在是上午")
	case hour < 18:
		fmt.Println("现在是下午")
	default:
		fmt.Println("现在是晚上")
	}
	
	// 带fallthrough的switch
	num := 75
	switch {
	case num >= 90:
		fmt.Println("成绩优秀")
	case num >= 80:
		fmt.Println("成绩良好")
	case num >= 70:
		fmt.Println("成绩中等")
		fallthrough // 继续执行下一个case
	case num >= 60:
		fmt.Println("成绩及格")
	default:
		fmt.Println("成绩不及格")
	}
}

// demonstrateLoops 展示循环语句
func demonstrateLoops() {
	fmt.Println("\n--- 循环语句 ---")
	
	// 基本for循环
	fmt.Println("基本for循环:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	
	// 类似while的for循环
	fmt.Println("\n类似while的for循环:")
	j := 1
	for j <= 5 {
		fmt.Printf("%d ", j)
		j++
	}
	fmt.Println()
	
	// 无限循环(这里用break控制)
	fmt.Println("\n带break的循环:")
	k := 1
	for {
		fmt.Printf("%d ", k)
		k++
		if k > 5 {
			break
		}
	}
	fmt.Println()
	
	// for-range循环遍历切片
	fmt.Println("\nfor-range遍历切片:")
	fruits := []string{"苹果", "香蕉", "橙子", "葡萄"}
	for index, value := range fruits {
		fmt.Printf("fruits[%d]=%s\n", index, value)
	}
	
	// for-range遍历map
	fmt.Println("\nfor-range遍历map:")
	scores := map[string]int{
		"数学": 90,
		"语文": 85,
		"英语": 95,
	}
	for subject, score := range scores {
		fmt.Printf("%s: %d分\n", subject, score)
	}
	
	// 忽略索引或值
	fmt.Println("\n忽略索引:")
	for _, fruit := range fruits {
		fmt.Printf("%s ", fruit)
	}
	fmt.Println()
	
	fmt.Println("\n忽略值:")
	for i, _ := range fruits {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

// demonstrateJumpStatements 展示跳转语句
func demonstrateJumpStatements() {
	fmt.Println("\n--- 跳转语句 ---")
	
	// break语句
	fmt.Println("break语句:")
	for i := 1; i <= 10; i++ {
		if i == 6 {
			break // 当i等于6时跳出循环
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	
	// continue语句
	fmt.Println("\ncontinue语句:")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // 跳过偶数
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	
	// 带标签的break
	fmt.Println("\n带标签的break:")
OuterLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i*j >= 4 {
				break OuterLoop // 跳出外层循环
			}
			fmt.Printf("(%d,%d) ", i, j)
		}
		fmt.Println()
	}
	fmt.Println()
	
	// goto语句(谨慎使用)
	fmt.Println("\ngoto语句:")
	i := 1
Start:
	if i <= 5 {
		fmt.Printf("%d ", i)
		i++
		goto Start
	}
	fmt.Println()
}

// demonstrateDefer 展示defer语句
func demonstrateDefer() {
	fmt.Println("\n--- defer语句 ---")
	
	// 基本defer用法
	defer fmt.Println("1. 这句话会在函数结束时打印")
	fmt.Println("2. 这句话会立即打印")
	
	// 多个defer (LIFO顺序执行)
	defer fmt.Println("3. 第二个defer")
	defer fmt.Println("4. 第三个defer")
	
	fmt.Println("5. 函数主体")
	
	// defer中的参数在声明时计算
	a := 10
	defer fmt.Printf("6. defer中的a = %d\n", a)
	a = 20
	fmt.Printf("7. 函数主体中的a = %d\n", a)
}