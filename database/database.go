package database

import "fmt"

// DemonstrateDatabase 展示数据库操作
func DemonstrateDatabase() {
	fmt.Println("=== 数据库操作学习 ===")

	// fmt.Println("\n1. SQL数据库操作")
	// DemonstrateSQL()

	// fmt.Println("\n2. ORM数据库操作")
	// DemonstrateORM()

	fmt.Println("\n3. 连接池和事务处理")
	DemonstratePoolAndTransaction()

	// 未来可以添加其他数据库类型
	// fmt.Println("\n4. NoSQL数据库操作")
	// DemonstrateNoSQL()
}
