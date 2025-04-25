package stdlib

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// DemonstrateOS 展示os包的使用
func DemonstrateOS() {
	// 获取环境变量
	fmt.Println("3.1 环境变量")
	fmt.Printf("HOME环境变量: %s\n", os.Getenv("HOME"))
	fmt.Printf("PATH环境变量: %s\n", os.Getenv("PATH"))

	// 设置环境变量
	os.Setenv("MY_VAR", "my_value")
	fmt.Printf("设置的环境变量MY_VAR: %s\n", os.Getenv("MY_VAR"))

	// 获取当前工作目录
	fmt.Println("\n3.2 文件系统操作")
	dir, _ := os.Getwd()
	fmt.Printf("当前工作目录: %s\n", dir)

	// 创建临时文件 (注释掉实际文件操作)
	fmt.Println("\n临时文件和目录操作示例 (代码已注释)")
	/*
		// 创建临时文件
		tempFile, err := ioutil.TempFile("", "example")
		if err == nil {
			fmt.Printf("创建临时文件: %s\n", tempFile.Name())
			tempFile.Write([]byte("临时文件内容"))
			tempFile.Close()
			// 删除临时文件
			os.Remove(tempFile.Name())
		}

		// 创建临时目录
		tempDir, err := ioutil.TempDir("", "example")
		if err == nil {
			fmt.Printf("创建临时目录: %s\n", tempDir)
			// 删除临时目录
			os.RemoveAll(tempDir)
		}
	*/

	// 文件信息
	fmt.Println("\n3.3 文件信息")
	info, err := os.Stat("go.mod")
	if err == nil {
		fmt.Printf("文件名: %s\n", info.Name())
		fmt.Printf("大小: %d 字节\n", info.Size())
		fmt.Printf("权限: %s\n", info.Mode())
		fmt.Printf("修改时间: %s\n", info.ModTime())
		fmt.Printf("是目录: %t\n", info.IsDir())
	} else {
		fmt.Printf("获取文件信息失败: %v\n", err)
	}

	// 遍历目录
	fmt.Println("\n3.4 遍历目录")
	count := 0
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if count < 5 { // 只显示前5个文件
			fmt.Println(path)
			count++
		}
		return nil
	})
	fmt.Println("... 等等")

	// 执行外部命令
	fmt.Println("\n3.5 执行外部命令")
	cmd := exec.Command("go", "version")
	output, err := cmd.Output()
	if err == nil {
		fmt.Printf("命令输出: %s", output)
	} else {
		fmt.Printf("执行命令失败: %v\n", err)
	}

	// 进程信息
	fmt.Println("\n3.6 进程信息")
	fmt.Printf("进程ID: %d\n", os.Getpid())
	fmt.Printf("父进程ID: %d\n", os.Getppid())
	fmt.Printf("用户ID: %d\n", os.Getuid())
}
