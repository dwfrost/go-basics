package main

import (
	"fmt"
	// "go-basics/constants"
	// "go-basics/controlflow"
	// "go-basics/concurrency"
	"go-basics/datatypes"
	// "go-basics/functions"
	// "go-basics/interfaces"
	// "go-basics/stdlib"
	// "go-basics/structs"
	// "go-basics/variables"
	// "go-basics/server"
	// "go-basics/database" // 新增数据库模块
	"go-basics/filestorage"
)

func main() {
	fmt.Println("=== Go语言基础语法学习 ===")

	// fmt.Println("\n=== 1. 变量 ===")
	// variables.DemonstrateVariables()

	// fmt.Println("\n=== 2. 常量 ===")
	// constants.DemonstrateConstants()

	fmt.Println("\n=== 3. 数据类型 ===")
	datatypes.DemonstrateDataTypes()

	// fmt.Println("\n=== 4. 控制结构 ===")
	// controlflow.DemonstrateControlFlow()

	// fmt.Println("\n=== 5. 函数 ===")
	// functions.DemonstrateFunctions()

	// fmt.Println("\n=== 6. 结构体和方法 ===")
	// structs.DemonstrateStructs()

	// fmt.Println("\n=== 7. 接口 ===")
	// interfaces.DemonstrateInterfaces()

	// fmt.Println("\n=== 8. 并发编程 ===")
	// fmt.Println("\n--- 8.1 Goroutines ---")
	// concurrency.DemonstrateGoroutines()

	// fmt.Println("\n--- 8.2 Channels ---")
	// concurrency.DemonstrateChannels()

	// fmt.Println("\n--- 8.3 实际应用 ---")
	// concurrency.DemonstratePracticalConcurrency()

	// fmt.Println("\n--- 8.4 Context包 ---")
	// concurrency.DemonstrateContext()

	// fmt.Println("\n=== 9. 标准库 ===")
	// stdlib.DemonstrateStdLib()

	// fmt.Println("\n=== 10. Go Http服务器 ===")
	// server.DemonstrateServer()

	// fmt.Println("\n=== 11. 数据库操作 ===")
	// database.DemonstrateDatabase()

	fmt.Println("\n=== 12. 文件存储 ===")
	filestorage.DemonstrateFileStorage()
}
