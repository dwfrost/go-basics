package filestorage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// DemonstrateBasicFileOperations 展示基本文件操作
func DemonstrateBasicFileOperations() {
	// 创建临时目录
	tempDir := createTempDir()
	defer os.RemoveAll(tempDir) // 函数结束时清理

	fmt.Printf("创建临时目录: %s\n", tempDir)

	// 1. 创建文件
	filePath := filepath.Join(tempDir, "example.txt")
	content := "这是一个示例文件内容。\n这是第二行。"

	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		fmt.Printf("创建文件失败: %v\n", err)
		return
	}
	fmt.Printf("创建文件: %s\n", filePath)

	// 2. 读取文件
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("读取文件失败: %v\n", err)
		return
	}
	fmt.Printf("读取文件内容:\n%s\n", string(data))

	// 3. 追加内容到文件
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打开文件失败: %v\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("\n这是追加的内容。")
	if err != nil {
		fmt.Printf("追加内容失败: %v\n", err)
		return
	}
	fmt.Println("追加内容到文件")

	// 4. 再次读取文件
	data, _ = os.ReadFile(filePath)
	fmt.Printf("更新后的文件内容:\n%s\n", string(data))

	// 5. 复制文件
	destPath := filepath.Join(tempDir, "copy.txt")
	err = copyFile(filePath, destPath)
	if err != nil {
		fmt.Printf("复制文件失败: %v\n", err)
		return
	}
	fmt.Printf("复制文件到: %s\n", destPath)

	// 6. 重命名文件
	newPath := filepath.Join(tempDir, "renamed.txt")
	err = os.Rename(destPath, newPath)
	if err != nil {
		fmt.Printf("重命名文件失败: %v\n", err)
		return
	}
	fmt.Printf("重命名文件为: %s\n", newPath)

	// 7. 删除文件
	err = os.Remove(newPath)
	if err != nil {
		fmt.Printf("删除文件失败: %v\n", err)
		return
	}
	fmt.Printf("删除文件: %s\n", newPath)

	// 8. 创建目录
	dirPath := filepath.Join(tempDir, "subdir")
	err = os.Mkdir(dirPath, 0755)
	if err != nil {
		fmt.Printf("创建目录失败: %v\n", err)
		return
	}
	fmt.Printf("创建目录: %s\n", dirPath)

	// 9. 遍历目录
	fmt.Println("遍历目录内容:")
	err = filepath.Walk(tempDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Printf("- %s (%s)\n", path, fileTypeString(info))
		return nil
	})
	if err != nil {
		fmt.Printf("遍历目录失败: %v\n", err)
	}
}

// 创建临时目录
func createTempDir() string {
	tempDir, err := os.MkdirTemp("", "fileops-*")
	if err != nil {
		fmt.Printf("创建临时目录失败: %v\n", err)
		return ""
	}
	return tempDir
}

// 复制文件
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

// 返回文件类型字符串
func fileTypeString(info os.FileInfo) string {
	if info.IsDir() {
		return "目录"
	}
	return "文件"
}
