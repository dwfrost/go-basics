package interfaces

import "fmt"

// DemonstrateInterfaces 展示Go语言中的接口相关特性
func DemonstrateInterfaces() {
	fmt.Println("1. 基本接口")
	var s Shape = Circle{Radius: 5}
	fmt.Printf("圆的面积: %.2f\n", s.Area())
	fmt.Printf("圆的周长: %.2f\n", s.Perimeter())

	s = Rectangle{Width: 4, Height: 6}
	fmt.Printf("矩形的面积: %.2f\n", s.Area())
	fmt.Printf("矩形的周长: %.2f\n", s.Perimeter())

	fmt.Println("\n2. 接口值")
	describeShape(Circle{Radius: 3})
	describeShape(Rectangle{Width: 2, Height: 5})

	fmt.Println("\n3. 空接口")
	showInfo(42)
	showInfo("Hello")
	showInfo(true)
	showInfo(Circle{Radius: 2})

	fmt.Println("\n4. 类型断言")
	checkType(42)
	checkType("Go语言")
	checkType(true)
	checkType(Circle{Radius: 1})

	fmt.Println("\n5. 类型选择")
	typeSwitch(42)
	typeSwitch("Go语言")
	typeSwitch(true)
	typeSwitch(Circle{Radius: 1})
}

// Shape 形状接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle 圆形结构体
type Circle struct {
	Radius float64
}

// Area 计算圆的面积
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Perimeter 计算圆的周长
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// Rectangle 矩形结构体
type Rectangle struct {
	Width, Height float64
}

// Area 计算矩形的面积
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter 计算矩形的周长
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 接口值示例
func describeShape(s Shape) {
	fmt.Printf("形状: %T, 面积: %.2f\n", s, s.Area())
}

// 空接口示例
func showInfo(i interface{}) {
	fmt.Printf("值: %v, 类型: %T\n", i, i)
}

// 类型断言示例
func checkType(i interface{}) {
	if val, ok := i.(string); ok {
		fmt.Printf("这是字符串: %s\n", val)
		return
	}
	if val, ok := i.(int); ok {
		fmt.Printf("这是整数: %d\n", val)
		return
	}
	fmt.Printf("未知类型: %T\n", i)
}

// 类型选择示例
func typeSwitch(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("整数: %d\n", v)
	case string:
		fmt.Printf("字符串: %s\n", v)
	case bool:
		fmt.Printf("布尔值: %t\n", v)
	default:
		fmt.Printf("其他类型: %T\n", v)
	}
}
