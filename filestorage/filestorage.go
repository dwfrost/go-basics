package filestorage

import "fmt"

// DemonstrateFileStorage 展示文件存储相关功能
func DemonstrateFileStorage() {
	fmt.Println("=== 文件存储学习 ===")

	fmt.Println("\n1. 基本文件操作")
	DemonstrateBasicFileOperations()

	fmt.Println("\n2. 文件上传")
	DemonstrateFileUpload()

	fmt.Println("\n3. 文件下载")
	DemonstrateFileDownload()

	fmt.Println("\n4. 大文件处理")
	DemonstrateLargeFileProcessing()

	fmt.Println("\n5. 文件元数据")
	DemonstrateFileMetadata()
}
