package filestorage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

// DemonstrateFileDownload 展示文件下载处理
func DemonstrateFileDownload() {
	// 创建下载目录
	downloadDir := createDownloadDir()
	defer os.RemoveAll(downloadDir)

	fmt.Printf("创建下载目录: %s\n", downloadDir)

	// 创建示例文件
	filePath := createSampleFile(downloadDir)
	if filePath == "" {
		return
	}

	// 模拟HTTP文件下载
	fmt.Println("\n1. 模拟HTTP文件下载")
	simulateFileDownload(filePath)

	// 断点续传
	fmt.Println("\n2. 断点续传")
	simulateResumeDownload(filePath)

	// 下载进度
	fmt.Println("\n3. 下载进度跟踪")
	trackDownloadProgress(filePath)

	// 文件下载安全性
	fmt.Println("\n4. 文件下载安全性考虑")
	fmt.Println("- 验证用户权限")
	fmt.Println("- 防止目录遍历攻击")
	fmt.Println("- 设置适当的Content-Disposition")
	fmt.Println("- 限制下载速率和并发数")
	fmt.Println("- 使用HTTPS传输")
}

// 创建下载目录
func createDownloadDir() string {
	downloadDir, err := os.MkdirTemp("", "downloads-*")
	if err != nil {
		fmt.Printf("创建下载目录失败: %v\n", err)
		return ""
	}
	return downloadDir
}

// 创建示例文件
func createSampleFile(dir string) string {
	filePath := filepath.Join(dir, "sample.txt")

	// 创建一个1MB的示例文件
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("创建示例文件失败: %v\n", err)
		return ""
	}
	defer file.Close()

	// 写入1MB的数据
	data := make([]byte, 1024)  // 1KB的数据块
	for i := 0; i < 1024; i++ { // 写入1024次，总共1MB
		for j := 0; j < len(data); j++ {
			data[j] = byte((i + j) % 256) // 生成一些变化的数据
		}
		_, err := file.Write(data)
		if err != nil {
			fmt.Printf("写入文件失败: %v\n", err)
			return ""
		}
	}

	fmt.Printf("创建示例文件: %s (1MB)\n", filePath)
	return filePath
}

// 模拟HTTP文件下载
func simulateFileDownload(filePath string) {
	// 在实际应用中，这个函数会是一个HTTP处理器
	fmt.Println("模拟处理下载请求:")

	// 检查文件是否存在
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Printf("- 文件不存在: %v\n", err)
		return
	}

	// 获取文件名
	filename := filepath.Base(filePath)
	fmt.Printf("- 文件名: %s\n", filename)

	// 设置响应头
	fmt.Println("- 设置响应头:")
	fmt.Printf("  Content-Disposition: attachment; filename=\"%s\"\n", filename)
	fmt.Printf("  Content-Type: application/octet-stream\n")
	fmt.Printf("  Content-Length: %d\n", fileInfo.Size())

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("- 打开文件失败: %v\n", err)
		return
	}
	defer file.Close()

	// 模拟发送文件
	fmt.Println("- 发送文件内容 (模拟)")
	fmt.Println("- 下载完成")
}

// 模拟断点续传
func simulateResumeDownload(filePath string) {
	fmt.Println("模拟断点续传:")

	// 获取文件信息
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Printf("- 获取文件信息失败: %v\n", err)
		return
	}

	// 模拟客户端已下载的字节数
	startByte := int64(512 * 1024) // 假设已下载了512KB

	// 检查范围是否有效
	if startByte >= fileInfo.Size() {
		fmt.Println("- 无效的范围请求: 起始位置超过文件大小")
		return
	}

	// 设置响应头
	fmt.Println("- 设置响应头:")
	fmt.Printf("  Content-Range: bytes %d-%d/%d\n", startByte, fileInfo.Size()-1, fileInfo.Size())
	fmt.Printf("  Content-Length: %d\n", fileInfo.Size()-startByte)
	fmt.Printf("  Content-Type: application/octet-stream\n")
	fmt.Printf("  Accept-Ranges: bytes\n")

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("- 打开文件失败: %v\n", err)
		return
	}
	defer file.Close()

	// 移动到起始位置
	_, err = file.Seek(startByte, io.SeekStart)
	if err != nil {
		fmt.Printf("- 设置文件位置失败: %v\n", err)
		return
	}

	// 模拟发送剩余部分
	remainingBytes := fileInfo.Size() - startByte
	fmt.Printf("- 从位置 %d 开始发送 %d 字节 (模拟)\n", startByte, remainingBytes)
	fmt.Println("- 续传完成")
}

// 跟踪下载进度
func trackDownloadProgress(filePath string) {
	fmt.Println("模拟下载进度跟踪:")

	// 获取文件信息
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Printf("- 获取文件信息失败: %v\n", err)
		return
	}

	totalSize := fileInfo.Size()
	fmt.Printf("- 文件总大小: %d 字节\n", totalSize)

	// 模拟下载进度
	downloadedBytes := int64(0)
	chunkSize := totalSize / 10 // 分10次下载

	for i := 0; i < 10; i++ {
		// 模拟下载延迟
		time.Sleep(100 * time.Millisecond)

		// 更新已下载字节数
		downloadedBytes += chunkSize
		if downloadedBytes > totalSize {
			downloadedBytes = totalSize
		}

		// 计算进度百分比
		progress := float64(downloadedBytes) / float64(totalSize) * 100

		// 打印进度条
		fmt.Printf("\r- 下载进度: [")
		for j := 0; j < 20; j++ {
			if float64(j) < progress/5 {
				fmt.Print("=")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Printf("] %.1f%% (%d/%d 字节)", progress, downloadedBytes, totalSize)
	}

	fmt.Println("\n- 下载完成")
}
