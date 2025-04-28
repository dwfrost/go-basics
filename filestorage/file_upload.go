package filestorage

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// DemonstrateFileUpload 展示文件上传处理
func DemonstrateFileUpload() {
	// 创建上传目录
	uploadDir := createUploadDir()
	defer os.RemoveAll(uploadDir)

	fmt.Printf("创建上传目录: %s\n", uploadDir)

	// 模拟文件上传
	fmt.Println("\n1. 模拟HTTP文件上传")
	simulateFileUpload(uploadDir)

	// 文件上传安全性
	fmt.Println("\n2. 文件上传安全性考虑")
	fmt.Println("- 验证文件类型和大小")
	fmt.Println("- 生成安全的文件名")
	fmt.Println("- 限制上传目录")
	fmt.Println("- 设置适当的文件权限")
	fmt.Println("- 扫描上传文件是否包含恶意内容")

	// 实现文件类型验证
	fmt.Println("\n3. 文件类型验证")
	validateFileTypes()

	// 实现文件大小限制
	fmt.Println("\n4. 文件大小限制")
	limitFileSize()

	// 安全的文件名处理
	fmt.Println("\n5. 安全的文件名处理")
	handleSecureFilenames()
}

// 创建上传目录
func createUploadDir() string {
	uploadDir, err := os.MkdirTemp("", "uploads-*")
	if err != nil {
		fmt.Printf("创建上传目录失败: %v\n", err)
		return ""
	}
	return uploadDir
}

// 模拟HTTP文件上传
func simulateFileUpload(uploadDir string) {
	// 创建一个模拟的multipart表单
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	// 添加文本字段
	_ = writer.WriteField("description", "这是一个测试文件")

	// 创建文件部分
	part, err := writer.CreateFormFile("file", "test.txt")
	if err != nil {
		fmt.Printf("创建表单文件失败: %v\n", err)
		return
	}

	// 写入文件内容
	_, err = part.Write([]byte("这是上传的文件内容"))
	if err != nil {
		fmt.Printf("写入文件内容失败: %v\n", err)
		return
	}

	// 关闭multipart writer
	writer.Close()

	// 创建请求
	req, err := http.NewRequest("POST", "http://example.com/upload", &b)
	if err != nil {
		fmt.Printf("创建请求失败: %v\n", err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 模拟处理上传请求
	fmt.Println("模拟处理上传请求:")
	handleFileUpload(req, uploadDir)
}

// 处理文件上传请求
func handleFileUpload(r *http.Request, uploadDir string) {
	// 在实际应用中，这个函数会是一个HTTP处理器

	// 解析multipart表单
	fmt.Println("- 解析multipart表单")

	// 获取表单字段
	fmt.Println("- 获取表单字段: description")

	// 获取上传的文件
	fmt.Println("- 获取上传的文件: test.txt")

	// 保存文件
	filePath := filepath.Join(uploadDir, "test.txt")
	fmt.Printf("- 保存文件到: %s\n", filePath)

	// 创建示例文件
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("创建文件失败: %v\n", err)
		return
	}
	defer file.Close()

	// 写入示例内容
	_, err = file.WriteString("这是上传的文件内容")
	if err != nil {
		fmt.Printf("写入文件失败: %v\n", err)
		return
	}

	fmt.Println("- 文件上传成功")
}

// 验证文件类型
func validateFileTypes() {
	allowedTypes := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".pdf":  true,
		".txt":  true,
	}

	testFiles := []string{
		"document.pdf",
		"image.jpg",
		"script.js",
		"data.txt",
		"malicious.php",
		"image.png.php", // 尝试绕过检查
	}

	for _, filename := range testFiles {
		ext := strings.ToLower(filepath.Ext(filename))
		if allowedTypes[ext] {
			fmt.Printf("- 文件 %s 类型有效 (%s)\n", filename, ext)
		} else {
			fmt.Printf("- 文件 %s 类型无效 (%s)\n", filename, ext)
		}
	}
}

// 限制文件大小
func limitFileSize() {
	maxSize := int64(10 * 1024 * 1024) // 10MB

	testSizes := []struct {
		name string
		size int64
	}{
		{"small.jpg", 500 * 1024},       // 500KB
		{"medium.pdf", 5 * 1024 * 1024}, // 5MB
		{"large.zip", 20 * 1024 * 1024}, // 20MB
	}

	for _, test := range testSizes {
		if test.size <= maxSize {
			fmt.Printf("- 文件 %s 大小有效 (%.2f MB)\n", test.name, float64(test.size)/(1024*1024))
		} else {
			fmt.Printf("- 文件 %s 大小超限 (%.2f MB > %.2f MB)\n",
				test.name, float64(test.size)/(1024*1024), float64(maxSize)/(1024*1024))
		}
	}
}

// 安全的文件名处理
func handleSecureFilenames() {
	unsafeNames := []string{
		"../../../etc/passwd",
		"malicious.php;.jpg",
		"image.jpg/../../config.php",
		"image with spaces.jpg",
		"image#with#special#chars.jpg",
	}

	for _, name := range unsafeNames {
		safe := generateSecureFilename(name)
		fmt.Printf("- 原始文件名: %s\n  安全文件名: %s\n", name, safe)
	}
}

// 生成安全的文件名
func generateSecureFilename(filename string) string {
	// 获取基本文件名（不含路径）
	filename = filepath.Base(filename)

	// 获取扩展名
	ext := filepath.Ext(filename)

	// 获取不含扩展名的文件名
	nameWithoutExt := strings.TrimSuffix(filename, ext)

	// 替换特殊字符
	nameWithoutExt = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '-' || r == '_' {
			return r
		}
		return '_'
	}, nameWithoutExt)

	// 添加时间戳前缀
	timestamp := fmt.Sprintf("%d", os.Getpid())

	// 组合安全的文件名
	return timestamp + "_" + nameWithoutExt + ext
}
