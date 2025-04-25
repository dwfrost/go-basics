package stdlib

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// DemonstrateIO 展示io包的使用
func DemonstrateIO() {
	// Reader接口
	fmt.Println("2.1 Reader接口")
	reader := strings.NewReader("Hello, Reader!")
	buffer := make([]byte, 8)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		fmt.Printf("读取了 %d 字节: %s\n", n, buffer[:n])
	}

	// Writer接口
	fmt.Println("\n2.2 Writer接口")
	var buf bytes.Buffer
	buf.WriteString("Hello, ")
	buf.WriteString("Writer!")
	fmt.Printf("缓冲区内容: %s\n", buf.String())

	// 复制数据
	fmt.Println("\n2.3 复制数据")
	src := strings.NewReader("这是源数据")
	dst := &bytes.Buffer{}
	written, _ := io.Copy(dst, src)
	fmt.Printf("复制了 %d 字节: %s\n", written, dst.String())

	// 多Reader合并
	fmt.Println("\n2.4 多Reader合并")
	r1 := strings.NewReader("第一部分 ")
	r2 := strings.NewReader("第二部分 ")
	r3 := strings.NewReader("第三部分")
	readers := io.MultiReader(r1, r2, r3)
	data, _ := io.ReadAll(readers)
	fmt.Printf("合并后的数据: %s\n", data)

	// TeeReader - 读取的同时写入
	fmt.Println("\n2.5 TeeReader")
	teeReader := io.TeeReader(strings.NewReader("TeeReader示例"), os.Stdout)
	io.ReadAll(teeReader) // 读取的同时会输出到标准输出
	fmt.Println()         // 添加换行

	// LimitReader - 限制读取的数量
	fmt.Println("\n2.6 LimitReader")
	original := strings.NewReader("这是一个很长的字符串，但我们只读取前10个字节")
	limited := io.LimitReader(original, 10)
	limitedData, _ := io.ReadAll(limited)
	fmt.Printf("限制读取的数据: %s\n", limitedData)
}
