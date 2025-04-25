package functions

import "fmt"

// DemonstrateFunctions 展示Go语言中的函数相关特性
func DemonstrateFunctions() {
	fmt.Println("1. 基本函数")
	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)

	fmt.Println("\n2. 多返回值")
	sum, diff := calculateSumAndDiff(10, 4)
	fmt.Printf("10 + 4 = %d, 10 - 4 = %d\n", sum, diff)

	fmt.Println("\n3. 命名返回值")
	area, perimeter := calculateRectangleProperties(5, 3)
	fmt.Printf("矩形面积和周长: %d, %d\n", area, perimeter)

	fmt.Println("\n4. 可变参数")
	fmt.Printf("1 + 2 + 3 + 4 = %d\n", sumNumbers(1, 2, 3, 4))

	fmt.Println("\n5. 函数作为值")
	operation := add
	fmt.Printf("通过函数变量调用: 6 + 4 = %d\n", operation(6, 4))

	fmt.Println("\n6. 匿名函数")
	func(x, y int) {
		fmt.Printf("匿名函数: %d + %d = %d\n", x, y, x+y)
	}(7, 3)

	fmt.Println("\n7. 闭包")
	counter := createCounter()
	fmt.Printf("计数器: %d\n", counter())
	fmt.Printf("计数器: %d\n", counter())
	fmt.Printf("计数器: %d\n", counter())
}

// 基本函数
func add(a, b int) int {
	return a + b
}

// 多返回值函数
func calculateSumAndDiff(a, b int) (int, int) {
	return a + b, a - b
}

// 命名返回值函数
func calculateRectangleProperties(length, width int) (area int, perimeter int) {
	area = length * width
	perimeter = 2 * (length + width)
	return // 裸返回
}

// 可变参数函数
func sumNumbers(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 闭包示例
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
