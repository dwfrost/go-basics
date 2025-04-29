package filestorage

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/tencentyun/cos-go-sdk-v5"
)

// DemonstrateTencentCOS 展示腾讯云对象存储的使用
func DemonstrateTencentCOS() {
	fmt.Println("=== 腾讯云对象存储(COS)示例 ===")

	// 创建COS客户端
	client, err := createCOSClient()
	if err != nil {
		fmt.Printf("创建COS客户端失败: %v\n", err)
		return
	}
	fmt.Println("client:", client)

	// 创建临时目录和测试文件
	tempDir, testFilePath := createTestFile()
	if tempDir != "" {
		defer os.RemoveAll(tempDir)
	}
	if testFilePath == "" {
		return
	}

	// 1. 上传文件
	fmt.Println("\n1. 文件上传")
	objectKey := "test/example.txt"
	err = uploadFileToCOS(client, testFilePath, objectKey)
	if err != nil {
		fmt.Printf("上传文件失败: %v\n", err)
		return
	}

	// 2. 下载文件
	fmt.Println("\n2. 文件下载")
	downloadPath := filepath.Join(tempDir, "downloaded_example.txt")
	err = downloadFileFromCOS(client, objectKey, downloadPath)
	if err != nil {
		fmt.Printf("下载文件失败: %v\n", err)
		return
	}

	// 3. 获取文件信息
	fmt.Println("\n3. 获取文件信息")
	err = getObjectInfo(client, objectKey)
	if err != nil {
		fmt.Printf("获取文件信息失败: %v\n", err)
	}

	// 4. 列出存储桶中的对象
	fmt.Println("\n4. 列出存储桶中的对象")
	err = listObjects(client, "test/")
	if err != nil {
		fmt.Printf("列出对象失败: %v\n", err)
	}

	// 5. 生成预签名URL
	fmt.Println("\n5. 生成预签名URL")
	err = generatePresignedURL(client, objectKey)
	if err != nil {
		fmt.Printf("生成预签名URL失败: %v\n", err)
	}

	// 6. 删除文件
	fmt.Println("\n6. 删除文件")
	err = deleteObject(client, objectKey)
	if err != nil {
		fmt.Printf("删除文件失败: %v\n", err)
	}
}

// 创建COS客户端
func createCOSClient() (*cos.Client, error) {
	godotenv.Load() // 加载.env文件
	// 从环境变量获取COS配置
	bucketURL := os.Getenv("COS_BUCKET_URL")
	secretID := os.Getenv("COS_SECRET_ID")
	secretKey := os.Getenv("COS_SECRET_KEY")

	// 检查必要的环境变量
	if bucketURL == "" || secretID == "" || secretKey == "" {
		return nil, fmt.Errorf("请设置COS_BUCKET_URL, COS_SECRET_ID和COS_SECRET_KEY环境变量")
	}

	// 解析存储桶URL
	u, err := url.Parse(bucketURL)
	if err != nil {
		return nil, fmt.Errorf("解析存储桶URL失败: %v", err)
	}

	// 创建基本传输对象
	b := &cos.BaseURL{BucketURL: u}

	// 创建COS客户端
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretID,
			SecretKey: secretKey,
		},
	})

	return client, nil
}

// 创建测试文件
func createTestFile() (string, string) {
	// 创建临时目录
	tempDir, err := os.MkdirTemp("", "cos-demo-*")
	if err != nil {
		fmt.Printf("创建临时目录失败: %v\n", err)
		return "", ""
	}

	// 创建测试文件
	testFilePath := filepath.Join(tempDir, "example.txt")
	content := "这是一个用于测试腾讯云对象存储的示例文件。\n创建时间: " + time.Now().Format(time.RFC3339)

	err = os.WriteFile(testFilePath, []byte(content), 0644)
	if err != nil {
		fmt.Printf("创建测试文件失败: %v\n", err)
		os.RemoveAll(tempDir)
		return "", ""
	}

	fmt.Printf("创建测试文件: %s\n", testFilePath)
	return tempDir, testFilePath
}

// 上传文件到COS
func uploadFileToCOS(client *cos.Client, filePath, objectKey string) error {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("打开文件失败: %v", err)
	}
	defer file.Close()

	// 获取文件信息
	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("获取文件信息失败: %v", err)
	}

	// 上传选项
	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: getContentType(filePath),
		},
	}

	// 执行上传
	ctx := context.Background()
	_, err = client.Object.Put(ctx, objectKey, file, opt)
	if err != nil {
		return fmt.Errorf("上传文件失败: %v", err)
	}

	fmt.Printf("成功上传文件 %s (大小: %d 字节) 到 %s\n",
		filePath, stat.Size(), objectKey)
	return nil
}

// 从COS下载文件
func downloadFileFromCOS(client *cos.Client, objectKey, downloadPath string) error {
	// 创建下载文件
	file, err := os.Create(downloadPath)
	if err != nil {
		return fmt.Errorf("创建下载文件失败: %v", err)
	}
	defer file.Close()

	// 下载选项
	opt := &cos.ObjectGetOptions{}

	// 执行下载
	ctx := context.Background()
	resp, err := client.Object.Get(ctx, objectKey, opt)
	if err != nil {
		return fmt.Errorf("下载文件失败: %v", err)
	}
	defer resp.Body.Close()

	// 将响应内容写入文件
	written, err := file.ReadFrom(resp.Body)
	if err != nil {
		return fmt.Errorf("写入文件失败: %v", err)
	}

	fmt.Printf("成功下载文件 %s 到 %s (大小: %d 字节)\n",
		objectKey, downloadPath, written)
	return nil
}

// 获取对象信息
func getObjectInfo(client *cos.Client, objectKey string) error {
	ctx := context.Background()
	resp, err := client.Object.Head(ctx, objectKey, nil)
	if err != nil {
		return fmt.Errorf("获取对象信息失败: %v", err)
	}

	fmt.Println("对象信息:")
	fmt.Printf("- 内容类型: %s\n", resp.Header.Get("Content-Type"))
	fmt.Printf("- 内容长度: %s 字节\n", resp.Header.Get("Content-Length"))
	fmt.Printf("- ETag: %s\n", resp.Header.Get("ETag"))
	fmt.Printf("- 最后修改时间: %s\n", resp.Header.Get("Last-Modified"))

	return nil
}

// 列出存储桶中的对象
func listObjects(client *cos.Client, prefix string) error {
	ctx := context.Background()
	opt := &cos.BucketGetOptions{
		Prefix:  prefix,
		MaxKeys: 100,
	}

	result, _, err := client.Bucket.Get(ctx, opt)
	if err != nil {
		return fmt.Errorf("列出对象失败: %v", err)
	}

	fmt.Printf("存储桶中的对象 (前缀: %s):\n", prefix)
	if len(result.Contents) == 0 {
		fmt.Println("- 没有找到对象")
		return nil
	}

	for _, object := range result.Contents {
		fmt.Printf("- %s (大小: %d 字节, 最后修改: %s)\n",
			object.Key, object.Size, object.LastModified)
	}

	return nil
}

// 生成预签名URL
func generatePresignedURL(client *cos.Client, objectKey string) error {
	ctx := context.Background()

	// 设置URL过期时间为1小时
	presignedURL, err := client.Object.GetPresignedURL(ctx, http.MethodGet, objectKey,
		os.Getenv("COS_SECRET_ID"), os.Getenv("COS_SECRET_KEY"), time.Hour, nil)
	if err != nil {
		return fmt.Errorf("生成预签名URL失败: %v", err)
	}

	fmt.Printf("预签名URL (有效期1小时): %s\n", presignedURL.String())
	return nil
}

// 删除对象
func deleteObject(client *cos.Client, objectKey string) error {
	ctx := context.Background()
	_, err := client.Object.Delete(ctx, objectKey)
	if err != nil {
		return fmt.Errorf("删除对象失败: %v", err)
	}

	fmt.Printf("成功删除对象: %s\n", objectKey)
	return nil
}

// 获取文件的Content-Type
func getContentType(filePath string) string {
	ext := strings.ToLower(filepath.Ext(filePath))
	switch ext {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".pdf":
		return "application/pdf"
	case ".txt":
		return "text/plain"
	case ".html":
		return "text/html"
	case ".json":
		return "application/json"
	case ".xml":
		return "application/xml"
	case ".mp3":
		return "audio/mpeg"
	case ".mp4":
		return "video/mp4"
	default:
		return "application/octet-stream"
	}
}
