package variables

import "fmt"

// DemonstrateVariables 展示Go语言中变量的声明和使用
func DemonstrateVariables() {
	// 1. 使用var关键字声明变量
	var a int // 声明变量但不初始化，默认为0
	fmt.Println("默认int值:", a)

	// 2. 声明变量并初始化
	var b int = 10
	fmt.Println("初始化int值:", b)

	// 3. 类型推断 - 省略类型
	var c = 20 // 编译器根据值自动推断类型
	fmt.Println("类型推断:", c)

	// 4. 短变量声明 (:=) - 只能在函数内部使用
	d := 30
	fmt.Println("短变量声明:", d)

	// 5. 多变量声明
	var e, f int = 40, 50
	fmt.Println("多变量声明:", e, f)

	// 6. 不同类型的多变量声明
	var (
		name   string  = "小明"
		age    int     = 25
		height float64 = 175.5
	)
	fmt.Printf("姓名: %s, 年龄: %d, 身高: %.1f\n", name, age, height)
	// fmt.Printf("姓名: %v, 年龄: %v, 身高: %v\n", name, age, height)

	// 7. 变量赋值
	var g int
	g = 60
	fmt.Println("变量赋值:", g)

	// 8. 变量交换 - Go特有
	x, y := 100, 200
	fmt.Println("交换前:", x, y)
	x, y = y, x
	fmt.Println("交换后:", x, y)
}
