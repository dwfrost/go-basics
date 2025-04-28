package filestorage

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// DemonstrateLargeFileProcessing 展示大文件处理
func DemonstrateLargeFileProcessing() {
	// 创建临时目录
	tempDir := createTempDir()
	defer os.RemoveAll(tempDir)

	// 创建大文件
	largeFilePath := createLargeFile(tempDir)
	if largeFilePath == "" {
		return
	}

	// 分块读取
	fmt.Println("\n1. 分块读取大文件")
	readLargeFileInChunks(largeFilePath)

	// 并行处理
	fmt.Println("\n2. 并行处理大文件")
	processLargeFileInParallel(largeFilePath)

	// 流式处理
	fmt.Println("\n3. 流式处理大文件")
	streamProcessLargeFile(largeFilePath)

	// 文件分片
	fmt.Println("\n4. 文件分片")
	splitLargeFile(largeFilePath, tempDir)
}

// 创建大文件
func createLargeFile(dir string) string {
	filePath := filepath.Join(dir, "large_file.dat")

	// 创建一个10MB的示例文件
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("创建大文件失败: %v\n", err)
		return ""
	}
	defer file.Close()

	// 使用缓冲写入提高性能
	writer := bufio.NewWriter(file)

	// 写入10MB的数据
	chunkSize := 64 * 1024 // 64KB的数据块
	data := make([]byte, chunkSize)

	// 填充数据
	for i := 0; i < len(data); i++ {
		data[i] = byte(i % 256)
	}

	totalSize := 10 * 1024 * 1024 // 10MB
	chunks := totalSize / chunkSize

	fmt.Printf("创建大文件: %s\n", filePath)
	fmt.Printf("- 总大小: %d 字节 (%d MB)\n", totalSize, totalSize/1024/1024)
	fmt.Printf("- 块大小: %d 字节 (%d KB)\n", chunkSize, chunkSize/1024)
	fmt.Printf("- 块数量: %d\n", chunks)

	start := time.Now()

	for i := 0; i < chunks; i++ {
		_, err := writer.Write(data)
		if err != nil {
			fmt.Printf("写入文件失败: %v\n", err)
			return ""
		}
	}

	// 确保所有数据都写入文件
	err = writer.Flush()
	if err != nil {
		fmt.Printf("刷新缓冲区失败: %v\n", err)
		return ""
	}

	elapsed := time.Since(start)
	fmt.Printf("- 写入耗时: %v\n", elapsed)

	return filePath
}

// 分块读取大文件
func readLargeFileInChunks(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("打开文件失败: %v\n", err)
		return
	}
	defer file.Close()

	// 获取文件大小
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("获取文件信息失败: %v\n", err)
		return
	}

	fileSize := fileInfo.Size()
	fmt.Printf("文件大小: %d 字节\n", fileSize)

	// 设置缓冲区大小
	bufferSize := 1024 * 1024 // 1MB
	buffer := make([]byte, bufferSize)

	totalBytesRead := int64(0)
	totalChunks := 0

	start := time.Now()

	// 分块读取
	for {
		bytesRead, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("读取文件失败: %v\n", err)
			return
		}

		totalBytesRead += int64(bytesRead)
		totalChunks++

		// 处理数据块 (这里只是计算前10个字节的和作为示例)
		sum := 0
		for i := 0; i < 10 && i < bytesRead; i++ {
			sum += int(buffer[i])
		}

		// 打印进度
		progress := float64(totalBytesRead) / float64(fileSize) * 100
		fmt.Printf("\r- 进度: %.1f%% (块 %d, 已读取 %d 字节)", progress, totalChunks, totalBytesRead)
	}

	elapsed := time.Since(start)
	fmt.Printf("\n- 读取完成: 总共读取 %d 块, %d 字节\n", totalChunks, totalBytesRead)
	fmt.Printf("- 读取耗时: %v\n", elapsed)
}

// 并行处理大文件
func processLargeFileInParallel(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("打开文件失败: %v\n", err)
		return
	}
	defer file.Close()

	// 获取文件大小
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("获取文件信息失败: %v\n", err)
		return
	}

	fileSize := fileInfo.Size()

	// 计算每个goroutine处理的数据大小
	numWorkers := 4
	chunkSize := fileSize / int64(numWorkers)

	fmt.Printf("并行处理文件:\n")
	fmt.Printf("- 工作线程数: %d\n", numWorkers)
	fmt.Printf("- 每个线程处理: %d 字节\n", chunkSize)

	var wg sync.WaitGroup
	results := make([]int64, numWorkers)

	start := time.Now()

	// 启动工作线程
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// 计算当前工作线程的起始位置和大小
			startPos := int64(id) * chunkSize
			endPos := startPos + chunkSize
			if id == numWorkers-1 {
				endPos = fileSize // 最后一个工作线程处理剩余的所有数据
			}

			// 打开文件并移动到起始位置
			workerFile, err := os.Open(filePath)
			if err != nil {
				fmt.Printf("工作线程 %d 打开文件失败: %v\n", id, err)
				return
			}
			defer workerFile.Close()

			_, err = workerFile.Seek(startPos, io.SeekStart)
			if err != nil {
				fmt.Printf("工作线程 %d 设置文件位置失败: %v\n", id, err)
				return
			}

			// 读取并处理数据
			buffer := make([]byte, 64*1024) // 64KB缓冲区
			bytesToRead := endPos - startPos
			bytesRead := int64(0)
			checksum := int64(0)

			for bytesRead < bytesToRead {
				// 计算本次应读取的字节数
				bufSize := bytesToRead - bytesRead
				if bufSize > int64(len(buffer)) {
					bufSize = int64(len(buffer))
				}

				// 读取数据
				n, err := workerFile.Read(buffer[:bufSize])
				if err != nil && err != io.EOF {
					fmt.Printf("工作线程 %d 读取失败: %v\n", id, err)
					return
				}
				if n == 0 {
					break
				}

				// 处理数据 (计算简单校验和)
				for i := 0; i < n; i++ {
					checksum += int64(buffer[i])
				}

				bytesRead += int64(n)
			}

			// 保存结果
			results[id] = checksum

			fmt.Printf("- 工作线程 %d 完成: 处理了 %d 字节, 校验和: %d\n", id, bytesRead, checksum)
		}(i)
	}

	// 等待所有工作线程完成
	wg.Wait()

	elapsed := time.Since(start)

	// 合并结果
	totalChecksum := int64(0)
	for _, checksum := range results {
		totalChecksum += checksum
	}

	fmt.Printf("- 所有工作线程完成, 总校验和: %d\n", totalChecksum)
	fmt.Printf("- 处理耗时: %v\n", elapsed)
}

// 流式处理大文件
func streamProcessLargeFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("打开文件失败: %v\n", err)
		return
	}
	defer file.Close()

	// 创建带缓冲的读取器
	reader := bufio.NewReader(file)

	// 创建MD5哈希计算器
	hash := md5.New()

	// 创建多写入器，同时写入到哈希计算器和/dev/null
	writer := io.MultiWriter(hash, io.Discard)

	fmt.Println("流式处理文件:")
	fmt.Println("- 计算文件MD5哈希值")

	start := time.Now()

	// 流式复制数据
	bytesProcessed, err := io.Copy(writer, reader)
	if err != nil {
		fmt.Printf("处理文件失败: %v\n", err)
		return
	}

	elapsed := time.Since(start)

	// 获取MD5哈希值
	md5sum := fmt.Sprintf("%x", hash.Sum(nil))

	fmt.Printf("- 处理完成: %d 字节\n", bytesProcessed)
	fmt.Printf("- MD5哈希值: %s\n", md5sum)
	fmt.Printf("- 处理耗时: %v\n", elapsed)
}

// 分片大文件
func splitLargeFile(filePath string, outputDir string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("打开文件失败: %v\n", err)
		return
	}
	defer file.Close()

	// 获取文件大小
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("获取文件信息失败: %v\n", err)
		return
	}

	fileSize := fileInfo.Size()

	// 设置分片大小和数量
	chunkSize := int64(2 * 1024 * 1024)                 // 2MB
	numChunks := (fileSize + chunkSize - 1) / chunkSize // 向上取整

	fmt.Printf("分片文件:\n")
	fmt.Printf("- 文件大小: %d 字节\n", fileSize)
	fmt.Printf("- 分片大小: %d 字节 (%d MB)\n", chunkSize, chunkSize/1024/1024)
	fmt.Printf("- 分片数量: %d\n", numChunks)

	// 创建分片目录
	chunksDir := filepath.Join(outputDir, "chunks")
	err = os.Mkdir(chunksDir, 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Printf("创建分片目录失败: %v\n", err)
		return
	}

	start := time.Now()

	// 分片处理
	buffer := make([]byte, 64*1024) // 64KB缓冲区

	for i := int64(0); i < numChunks; i++ {
		// 计算当前分片的起始位置和大小
		startPos := i * chunkSize
		currentChunkSize := chunkSize
		if startPos+currentChunkSize > fileSize {
			currentChunkSize = fileSize - startPos
		}

		// 创建分片文件
		chunkFileName := fmt.Sprintf("chunk_%03d.dat", i+1)
		chunkFilePath := filepath.Join(chunksDir, chunkFileName)

		chunkFile, err := os.Create(chunkFilePath)
		if err != nil {
			fmt.Printf("创建分片文件失败: %v\n", err)
			return
		}

		// 移动到起始位置
		_, err = file.Seek(startPos, io.SeekStart)
		if err != nil {
			fmt.Printf("设置文件位置失败: %v\n", err)
			chunkFile.Close()
			return
		}

		// 写入分片数据
		bytesRemaining := currentChunkSize
		for bytesRemaining > 0 {
			// 计算本次应读取的字节数
			bufSize := bytesRemaining
			if bufSize > int64(len(buffer)) {
				bufSize = int64(len(buffer))
			}

			// 读取数据
			n, err := file.Read(buffer[:bufSize])
			if err != nil && err != io.EOF {
				fmt.Printf("读取文件失败: %v\n", err)
				chunkFile.Close()
				return
			}
			if n == 0 {
				break
			}

			// 写入数据到分片文件
			_, err = chunkFile.Write(buffer[:n])
			if err != nil {
				fmt.Printf("写入分片文件失败: %v\n", err)
				chunkFile.Close()
				return
			}

			bytesRemaining -= int64(n)
		}

		chunkFile.Close()

		// 打印进度
		progress := float64(i+1) / float64(numChunks) * 100
		fmt.Printf("\r- 进度: %.1f%% (已创建 %d/%d 个分片)", progress, i+1, numChunks)
	}

	elapsed := time.Since(start)
	fmt.Printf("\n- 分片完成: 总共创建 %d 个分片\n", numChunks)
	fmt.Printf("- 分片耗时: %v\n", elapsed)

	// 验证分片
	fmt.Println("\n验证分片完整性:")
	validateChunks(chunksDir, fileSize)
}

// 验证分片完整性
func validateChunks(chunksDir string, originalSize int64) {
	// 获取所有分片文件
	files, err := os.ReadDir(chunksDir)
	if err != nil {
		fmt.Printf("读取分片目录失败: %v\n", err)
		return
	}

	// 计算所有分片的总大小
	var totalSize int64
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		info, err := file.Info()
		if err != nil {
			fmt.Printf("获取文件信息失败: %v\n", err)
			continue
		}

		totalSize += info.Size()
	}

	// 验证大小
	if totalSize == originalSize {
		fmt.Printf("- 验证成功: 分片总大小 %d 字节与原始文件大小一致\n", totalSize)
	} else {
		fmt.Printf("- 验证失败: 分片总大小 %d 字节与原始文件大小 %d 字节不一致\n", totalSize, originalSize)
	}
}
