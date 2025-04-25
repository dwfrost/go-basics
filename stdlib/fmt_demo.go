package stdlib

import (
	"fmt"
	"os"
)

// DemonstrateFmt 展示fmt包的使用
func DemonstrateFmt() {
	// 基本打印函数
	fmt.Println("1.1 基本打印函数")
	fmt.Print("这是Print: 不会自动换行 ")
	fmt.Println("这是Println: 会自动换行")
	fmt.Printf("这是Printf: 支持格式化, 例如: %d, %s, %.2f\n", 10, "字符串", 3.1415926)

	// 格式化字符串
	fmt.Println("\n1.2 格式化字符串")
	s := fmt.Sprintf("Sprintf: 格式化并返回字符串: %d, %s", 42, "Go")
	fmt.Println(s)

	// 常用格式化动词
	fmt.Println("\n1.3 常用格式化动词")
	fmt.Printf("%%v 默认格式: %v\n", []int{1, 2, 3})
	fmt.Printf("%%+v 包含字段名: %+v\n", struct {
		Name string
		Age  int
	}{"张三", 30})
	fmt.Printf("%%#v Go语法表示: %#v\n", []string{"a", "b"})
	fmt.Printf("%%T 类型: %T\n", 3.14)
	fmt.Printf("%%d 十进制整数: %d\n", 42)
	fmt.Printf("%%b 二进制: %b\n", 42)
	fmt.Printf("%%x 十六进制: %x\n", 42)
	fmt.Printf("%%f 浮点数: %f\n", 3.1415926)
	fmt.Printf("%%9.2f 宽度9小数点2位: %9.2f\n", 3.1415926)
	fmt.Printf("%%q 带引号字符串: %q\n", "Go语言")

	// 读取输入 (注释掉以避免阻塞)
	/*
		fmt.Println("\n1.4 读取输入 (示例代码已注释)")
		var name string
		var age int
		fmt.Print("请输入姓名和年龄: ")
		fmt.Scanf("%s %d", &name, &age)
		fmt.Printf("你好，%s! 你的年龄是%d岁。\n", name, age)
	*/

	// 输出到不同的目标
	fmt.Println("\n1.5 输出到不同的目标")
	fmt.Fprintln(os.Stdout, "输出到标准输出")
	// 输出到文件示例
	/*
		file, err := os.Create("output.txt")
		if err == nil {
			fmt.Fprintln(file, "输出到文件")
			file.Close()
		}
	*/
}
