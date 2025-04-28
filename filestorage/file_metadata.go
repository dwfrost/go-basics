package filestorage

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"os"
	"path/filepath"
	"time"
)

// FileMetadata 文件元数据结构
type FileMetadata struct {
	Name         string    `json:"name"`
	Path         string    `json:"path"`
	Size         int64     `json:"size"`
	Mode         string    `json:"mode"`
	ModTime      time.Time `json:"mod_time"`
	IsDir        bool      `json:"is_dir"`
	MimeType     string    `json:"mime_type,omitempty"`
	MD5Hash      string    `json:"md5_hash,omitempty"`
	SHA256Hash   string    `json:"sha256_hash,omitempty"`
	CreationTime time.Time `json:"creation_time,omitempty"`
}

// DemonstrateFileMetadata 展示文件元数据处理
func DemonstrateFileMetadata() {
	// 创建临时目录
	tempDir := createTempDir()
	defer os.RemoveAll(tempDir)

	// 创建示例文件
	filePath := createMetadataTestFile(tempDir)
	if filePath == "" {
		return
	}

	// 基本元数据
	fmt.Println("\n1. 基本文件元数据")
	metadata := getBasicMetadata(filePath)
	printMetadata(metadata)

	// 高级元数据
	fmt.Println("\n2. 高级文件元数据")
	enrichMetadata(metadata, filePath)
	printMetadata(metadata)

	// 元数据存储
	fmt.Println("\n3. 元数据存储")
	storeMetadata(metadata, tempDir)

	// 元数据索引
	fmt.Println("\n4. 元数据索引和搜索")
	demonstrateMetadataIndex(tempDir)

	// 元数据应用场景
	fmt.Println("\n5. 元数据应用场景")
	fmt.Println("- 文件去重")
	fmt.Println("- 文件分类")
	fmt.Println("- 文件搜索")
	fmt.Println("- 文件完整性验证")
	fmt.Println("- 文件版本控制")
}

// 创建元数据测试文件
func createMetadataTestFile(dir string) string {
	filePath := filepath.Join(dir, "metadata_test.txt")

	// 创建文件
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("创建文件失败: %v\n", err)
		return ""
	}
	defer file.Close()

	// 写入内容
	content := "这是一个用于测试元数据的文件。\n它包含多行文本内容。\n这是第三行。"
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("写入文件失败: %v\n", err)
		return ""
	}

	fmt.Printf("创建测试文件: %s\n", filePath)
	return filePath
}

// 获取基本元数据
func getBasicMetadata(filePath string) *FileMetadata {
	// 获取文件信息
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Printf("获取文件信息失败: %v\n", err)
		return nil
	}

	// 创建元数据结构
	metadata := &FileMetadata{
		Name:    fileInfo.Name(),
		Path:    filePath,
		Size:    fileInfo.Size(),
		Mode:    fileInfo.Mode().String(),
		ModTime: fileInfo.ModTime(),
		IsDir:   fileInfo.IsDir(),
	}

	return metadata
}

// 打印元数据
func printMetadata(metadata *FileMetadata) {
	if metadata == nil {
		fmt.Println("无元数据")
		return
	}

	fmt.Println("文件元数据:")
	fmt.Printf("- 名称: %s\n", metadata.Name)
	fmt.Printf("- 路径: %s\n", metadata.Path)
	fmt.Printf("- 大小: %d 字节\n", metadata.Size)
	fmt.Printf("- 权限: %s\n", metadata.Mode)
	fmt.Printf("- 修改时间: %s\n", metadata.ModTime)
	fmt.Printf("- 是否目录: %t\n", metadata.IsDir)

	if metadata.MimeType != "" {
		fmt.Printf("- MIME类型: %s\n", metadata.MimeType)
	}

	if metadata.MD5Hash != "" {
		fmt.Printf("- MD5哈希: %s\n", metadata.MD5Hash)
	}

	if metadata.SHA256Hash != "" {
		fmt.Printf("- SHA256哈希: %s\n", metadata.SHA256Hash)
	}

	if !metadata.CreationTime.IsZero() {
		fmt.Printf("- 创建时间: %s\n", metadata.CreationTime)
	}
}

// 丰富元数据
func enrichMetadata(metadata *FileMetadata, filePath string) {
	if metadata == nil {
		return
	}

	// 获取MIME类型
	ext := filepath.Ext(filePath)
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}
	metadata.MimeType = mimeType

	// 计算MD5哈希
	md5hash, err := calculateMD5(filePath)
	if err == nil {
		metadata.MD5Hash = md5hash
	}

	// 计算SHA256哈希
	sha256hash, err := calculateSHA256(filePath)
	if err == nil {
		metadata.SHA256Hash = sha256hash
	}

	// 设置创建时间（在某些系统上可能与修改时间相同）
	metadata.CreationTime = time.Now()
}

// 计算MD5哈希
func calculateMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

// 计算SHA256哈希
func calculateSHA256(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

// 存储元数据
func storeMetadata(metadata *FileMetadata, dir string) {
	if metadata == nil {
		return
	}

	// 创建元数据JSON
	data, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		fmt.Printf("序列化元数据失败: %v\n", err)
		return
	}

	// 写入文件
	metadataPath := filepath.Join(dir, metadata.Name+".metadata.json")
	err = os.WriteFile(metadataPath, data, 0644)
	if err != nil {
		fmt.Printf("写入元数据文件失败: %v\n", err)
		return
	}

	fmt.Printf("元数据已存储到: %s\n", metadataPath)

	// 读取并显示存储的元数据
	storedData, err := os.ReadFile(metadataPath)
	if err != nil {
		fmt.Printf("读取元数据文件失败: %v\n", err)
		return
	}

	fmt.Println("存储的元数据内容:")
	fmt.Println(string(storedData))
}

// 演示元数据索引
func demonstrateMetadataIndex(dir string) {
	// 创建多个测试文件
	files := createMultipleTestFiles(dir, 5)

	// 创建元数据索引
	index := createMetadataIndex(files)

	// 显示索引
	fmt.Println("元数据索引:")
	for key, value := range index {
		fmt.Printf("- %s: %d 个文件\n", key, len(value))
	}

	// 模拟搜索
	fmt.Println("\n搜索示例:")

	// 按大小搜索
	sizeKey := "size:100-200"
	if files, ok := index[sizeKey]; ok {
		fmt.Printf("- 大小在100-200字节之间的文件: %d 个\n", len(files))
		for _, file := range files {
			fmt.Printf("  - %s (%d 字节)\n", file.Name, file.Size)
		}
	}

	// 按类型搜索
	typeKey := "type:text"
	if files, ok := index[typeKey]; ok {
		fmt.Printf("- 文本类型的文件: %d 个\n", len(files))
		for _, file := range files {
			fmt.Printf("  - %s (%s)\n", file.Name, file.MimeType)
		}
	}
}

// 创建多个测试文件
func createMultipleTestFiles(dir string, count int) []*FileMetadata {
	var files []*FileMetadata

	for i := 1; i <= count; i++ {
		// 创建不同类型的文件
		var filePath string
		var content string

		switch i % 3 {
		case 0:
			filePath = filepath.Join(dir, fmt.Sprintf("text_%d.txt", i))
			content = fmt.Sprintf("这是第 %d 个文本文件。", i)
		case 1:
			filePath = filepath.Join(dir, fmt.Sprintf("data_%d.dat", i))
			content = fmt.Sprintf("DATA-%d-BINARY-CONTENT", i)
		case 2:
			filePath = filepath.Join(dir, fmt.Sprintf("config_%d.json", i))
			content = fmt.Sprintf(`{"id": %d, "name": "配置文件%d"}`, i, i)
		}

		// 写入文件
		err := os.WriteFile(filePath, []byte(content), 0644)
		if err != nil {
			fmt.Printf("创建文件失败: %v\n", err)
			continue
		}

		// 获取元数据
		metadata := getBasicMetadata(filePath)
		if metadata != nil {
			enrichMetadata(metadata, filePath)
			files = append(files, metadata)
		}
	}

	fmt.Printf("创建了 %d 个测试文件\n", len(files))
	return files
}

// 创建元数据索引
func createMetadataIndex(files []*FileMetadata) map[string][]*FileMetadata {
	index := make(map[string][]*FileMetadata)

	for _, file := range files {
		// 按大小范围索引
		var sizeKey string
		switch {
		case file.Size < 100:
			sizeKey = "size:0-100"
		case file.Size < 200:
			sizeKey = "size:100-200"
		default:
			sizeKey = "size:200+"
		}
		index[sizeKey] = append(index[sizeKey], file)

		// 按文件类型索引
		var typeKey string
		if file.MimeType != "" {
			if len(file.MimeType) >= 5 && file.MimeType[:5] == "text/" {
				typeKey = "type:text"
			} else if len(file.MimeType) >= 12 && file.MimeType[:12] == "application/" {
				typeKey = "type:application"
			} else {
				typeKey = "type:other"
			}
			index[typeKey] = append(index[typeKey], file)
		}

		// 按修改时间索引
		timeKey := fmt.Sprintf("time:%d", file.ModTime.Year())
		index[timeKey] = append(index[timeKey], file)
	}

	return index
}
