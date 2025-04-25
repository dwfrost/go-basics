package stdlib

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// DemonstrateStrings 展示strings和strconv包的使用
func DemonstrateStrings() {
	// strings包基本操作
	fmt.Println("7.1 strings包基本操作")
	s := "Hello, 世界! Go语言很棒。"

	fmt.Printf("原始字符串: %s\n", s)
	fmt.Printf("长度 (字节): %d\n", len(s))
	fmt.Printf("长度 (字符): %d\n", len([]rune(s)))
	fmt.Printf("包含'Go': %t\n", strings.Contains(s, "Go"))
	fmt.Printf("前缀'Hello': %t\n", strings.HasPrefix(s, "Hello"))
	fmt.Printf("后缀'棒。': %t\n", strings.HasSuffix(s, "棒。"))

	// 查找和替换
	fmt.Println("\n7.2 查找和替换")
	fmt.Printf("'世界'的位置: %d\n", strings.Index(s, "世界"))
	fmt.Printf("替换'世界'为'中国': %s\n", strings.Replace(s, "世界", "中国", 1))
	fmt.Printf("替换所有空格: %s\n", strings.ReplaceAll(s, " ", "_"))

	// 分割和连接
	fmt.Println("\n7.3 分割和连接")
	parts := strings.Split("a,b,c,d", ",")
	fmt.Printf("分割后: %v\n", parts)
	joined := strings.Join(parts, "-")
	fmt.Printf("连接后: %s\n", joined)

	// 大小写转换
	fmt.Println("\n7.4 大小写转换")
	fmt.Printf("转大写: %s\n", strings.ToUpper("Hello, World"))
	fmt.Printf("转小写: %s\n", strings.ToLower("Hello, World"))
	fmt.Printf("首字母大写: %s\n", strings.Title("hello world"))

	// 修剪
	fmt.Println("\n7.5 修剪")
	fmt.Printf("修剪空格: %q\n", strings.TrimSpace("  hello  "))
	fmt.Printf("修剪前缀: %q\n", strings.TrimPrefix("TestHello", "Test"))
	fmt.Printf("修剪后缀: %q\n", strings.TrimSuffix("HelloTest", "Test"))
	fmt.Printf("修剪指定字符: %q\n", strings.Trim("!!!Hello!!!", "!"))

	// 字符串构建器
	fmt.Println("\n7.6 字符串构建器")
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World")
	fmt.Printf("构建的字符串: %s\n", builder.String())
	fmt.Printf("构建器容量: %d\n", builder.Cap())

	// strconv包 - 字符串转换
	fmt.Println("\n7.7 strconv包 - 字符串转换")

	// 字符串转数字
	i, _ := strconv.Atoi("42")
	fmt.Printf("字符串转整数: %d (%T)\n", i, i)

	f, _ := strconv.ParseFloat("3.14159", 64)
	fmt.Printf("字符串转浮点数: %f (%T)\n", f, f)

	b, _ := strconv.ParseBool("true")
	fmt.Printf("字符串转布尔值: %t (%T)\n", b, b)

	// 数字转字符串
	fmt.Printf("整数转字符串: %s (%T)\n", strconv.Itoa(123), strconv.Itoa(123))
	fmt.Printf("浮点数转字符串: %s\n", strconv.FormatFloat(3.14159, 'f', 2, 64))
	fmt.Printf("布尔值转字符串: %s\n", strconv.FormatBool(true))

	// 进制转换
	fmt.Printf("十进制转二进制: %s\n", strconv.FormatInt(42, 2))
	fmt.Printf("十进制转十六进制: %s\n", strconv.FormatInt(42, 16))

	// 字符处理
	fmt.Println("\n7.8 字符处理")
	for _, r := range "Hello, 世界!" {
		fmt.Printf("%c: 是字母=%t, 是数字=%t, 是空白=%t\n",
			r, unicode.IsLetter(r), unicode.IsDigit(r), unicode.IsSpace(r))
	}
}
