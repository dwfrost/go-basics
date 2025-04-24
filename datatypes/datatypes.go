package datatypes

import (
	"fmt"
	"math/cmplx"
	"strconv"
	"unsafe"
)

// DemonstrateDataTypes 展示Go语言中的各种数据类型
func DemonstrateDataTypes() {
	// 1. 基本数据类型
	demonstrateBasicTypes()

	// 2. 复合数据类型
	demonstrateCompositeTypes()

	// 3. 类型转换
	demonstrateTypeConversion()
}

// demonstrateBasicTypes 展示基本数据类型
func demonstrateBasicTypes() {
	fmt.Println("--- 基本数据类型 ---")

	// 整数类型
	var intVal int = 42
	var int8Val int8 = 127
	var int16Val int16 = 32767
	var int32Val int32 = 2147483647
	var int64Val int64 = 9223372036854775807

	var uintVal uint = 42
	var uint8Val uint8 = 255
	var uint16Val uint16 = 65535
	var uint32Val uint32 = 4294967295

	fmt.Println("整数类型:")
	fmt.Printf("  int: %d, 大小: %d字节\n", intVal, int(unsafe.Sizeof(intVal)))
	fmt.Printf("  int8: %d, 范围: -128 到 127\n", int8Val)
	fmt.Printf("  int16: %d, 范围: -32768 到 32767\n", int16Val)
	fmt.Printf("  int32: %d\n", int32Val)
	fmt.Printf("  int64: %d\n", int64Val)
	fmt.Printf("  uint: %d\n", uintVal)
	fmt.Printf("  uint8: %d\n", uint8Val)
	fmt.Printf("  uint16: %d\n", uint16Val)
	fmt.Printf("  uint32: %d\n", uint32Val)

	// 浮点类型
	var float32Val float32 = 3.14
	var float64Val float64 = 3.14159265358979

	fmt.Println("\n浮点类型:")
	fmt.Printf("  float32: %f\n", float32Val)
	fmt.Printf("  float64: %f\n", float64Val)

	// 复数类型
	var complexVal complex128 = cmplx.Sqrt(-5 + 12i)
	fmt.Println("\n复数类型:")
	fmt.Printf("  complex128: %v\n", complexVal)

	// 布尔类型
	var boolTrue bool = true
	var boolFalse bool = false

	fmt.Println("\n布尔类型:")
	fmt.Printf("  true: %t\n", boolTrue)
	fmt.Printf("  false: %t\n", boolFalse)

	// 字符串类型
	var str string = "你好，Go语言!"

	fmt.Println("\n字符串类型:")
	fmt.Printf("  string: %s, 长度: %d\n", str, len(str))

	// 字符类型
	var char rune = '中'
	var byte1 byte = 'A'

	fmt.Println("\n字符类型:")
	fmt.Printf("  rune: %c (%d)\n", char, char)
	fmt.Printf("  byte: %c (%d)\n", byte1, byte1)
}

// demonstrateCompositeTypes 展示复合数据类型
func demonstrateCompositeTypes() {
	fmt.Println("\n--- 复合数据类型 ---")

	// 数组
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("数组:")
	fmt.Printf("  %v, 长度: %d\n", arr, len(arr))

	// 简短声明数组
	arr2 := [3]string{"苹果", "香蕉", "橙子"}
	fmt.Printf("  %v\n", arr2)

	// 自动计算长度
	arr3 := [...]int{10, 20, 30, 40}
	fmt.Printf("  %v, 长度: %d\n", arr3, len(arr3))

	// 切片
	fmt.Println("\n切片:")
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("  %v, 长度: %d, 容量: %d\n", slice1, len(slice1), cap(slice1))

	// 使用make创建切片
	slice2 := make([]int, 3, 5) // 长度为3，容量为5
	fmt.Printf("  %v, 长度: %d, 容量: %d\n", slice2, len(slice2), cap(slice2))

	// 切片操作
	fmt.Println("  切片操作:")
	fmt.Printf("    slice1[1:3]: %v\n", slice1[1:3])

	// 追加元素
	slice2 = append(slice2, 100, 200)
	fmt.Printf("    追加后: %v\n", slice2)

	// 映射(Map)
	fmt.Println("\n映射:")
	m := map[string]int{
		"小明": 18,
		"小红": 20,
		"小李": 19,
	}
	fmt.Printf("  %v\n", m)

	// 使用make创建map
	scores := make(map[string]int)
	scores["数学"] = 90
	scores["语文"] = 85
	scores["英语"] = 95
	fmt.Printf("  %v\n", scores)

	// 结构体
	fmt.Println("\n结构体:")
	type Person struct {
		Name string
		Age  int
		City string
	}

	p1 := Person{
		Name: "张三",
		Age:  30,
		City: "北京",
	}
	fmt.Printf("  %+v\n", p1)

	// 匿名结构体
	student := struct {
		Name  string
		Grade int
	}{
		Name:  "李四",
		Grade: 3,
	}
	fmt.Printf("  %+v\n", student)
}

// demonstrateTypeConversion 展示类型转换
func demonstrateTypeConversion() {
	fmt.Println("\n--- 类型转换 ---")

	// 整数转换
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Println("整数转换:")
	fmt.Printf("  int -> float64 -> uint: %d -> %f -> %d\n", i, f, u)

	// 字符串与数字转换
	var i2 int = 65
	var s string = string(i2) // 将ASCII码转为字符

	fmt.Println("\n字符串与数字:")
	fmt.Printf("  int -> string: %d -> %s\n", i2, s)

	// 使用strconv包进行字符串转换

	str := "123"
	num, _ := strconv.Atoi(str)
	fmt.Printf("  string -> int: %s -> %d\n", str, num)

	str2 := strconv.Itoa(456)
	fmt.Printf("  int -> string: %d -> %s\n", 456, str2)

	// 浮点数转字符串
	str3 := strconv.FormatFloat(3.1415, 'f', 2, 64)
	fmt.Printf("  float -> string: %f -> %s\n", 3.1415, str3)
}
