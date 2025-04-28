package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/glebarez/sqlite"     // 使用纯Go实现的SQLite驱动
	_ "github.com/go-sql-driver/mysql" // 导入MySQL驱动
)

// User 用户结构体
type User struct {
	ID        int
	Username  string
	Email     string
	CreatedAt time.Time
}

// DemonstrateSQL 展示database/sql包的使用
func DemonstrateSQL() {
	fmt.Println("=== SQL数据库操作示例 ===")

	// 使用SQLite内存数据库进行演示
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatalf("无法打开数据库: %v", err)
	}
	defer db.Close()

	// 设置连接池参数
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	// 检查连接
	if err := db.Ping(); err != nil {
		log.Fatalf("无法连接到数据库: %v", err)
	}
	fmt.Println("1. 成功连接到数据库")

	// 创建表
	createTable(db)
	fmt.Println("\n2. 成功创建表")

	// 插入数据
	insertUsers(db)
	fmt.Println("\n3. 成功插入数据")

	// 查询单条数据
	queryUser(db, 1)

	// 查询多条数据
	queryAllUsers(db)

	// 更新数据
	updateUser(db, 1, "updated_user", "updated@example.com")
	fmt.Println("\n6. 成功更新数据")
	queryUser(db, 1) // 验证更新

	// 删除数据
	deleteUser(db, 2)
	fmt.Println("\n7. 成功删除数据")
	queryAllUsers(db) // 验证删除

	// 事务示例
	transactionExample(db)
	fmt.Println("\n8. 事务示例完成")
	queryAllUsers(db) // 验证事务

	// 预处理语句示例
	preparedStatementExample(db)
	fmt.Println("\n9. 预处理语句示例完成")
}

// createTable 创建用户表
func createTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		email TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("创建表失败: %v", err)
	}
}

// insertUsers 插入示例用户数据
func insertUsers(db *sql.DB) {
	// 插入多个用户
	users := []struct {
		username string
		email    string
	}{
		{"user1", "user1@example.com"},
		{"user2", "user2@example.com"},
		{"user3", "user3@example.com"},
	}

	for _, user := range users {
		result, err := db.Exec(
			"INSERT INTO users (username, email) VALUES (?, ?)",
			user.username, user.email,
		)
		if err != nil {
			log.Printf("插入用户失败: %v", err)
			continue
		}

		id, _ := result.LastInsertId()
		fmt.Printf("插入用户 ID: %d, 用户名: %s\n", id, user.username)
	}
}

// queryUser 查询单个用户
func queryUser(db *sql.DB, id int) {
	fmt.Printf("\n4. 查询用户 ID: %d\n", id)

	// 查询单行
	var user User
	err := db.QueryRow(
		"SELECT id, username, email, created_at FROM users WHERE id = ?", id,
	).Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("未找到ID为%d的用户\n", id)
		} else {
			log.Printf("查询失败: %v", err)
		}
		return
	}

	fmt.Printf("用户: ID=%d, 用户名=%s, 邮箱=%s, 创建时间=%v\n",
		user.ID, user.Username, user.Email, user.CreatedAt)
}

// queryAllUsers 查询所有用户
func queryAllUsers(db *sql.DB) {
	fmt.Println("\n5. 查询所有用户")

	// 查询多行
	rows, err := db.Query("SELECT id, username, email, created_at FROM users")
	if err != nil {
		log.Printf("查询失败: %v", err)
		return
	}
	defer rows.Close() // 重要：关闭rows以释放连接

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
		if err != nil {
			log.Printf("扫描行失败: %v", err)
			continue
		}
		users = append(users, user)
	}

	// 检查迭代过程中是否有错误
	if err := rows.Err(); err != nil {
		log.Printf("行迭代错误: %v", err)
	}

	// 打印结果
	fmt.Printf("找到 %d 个用户:\n", len(users))
	for _, user := range users {
		fmt.Printf("- ID=%d, 用户名=%s, 邮箱=%s\n",
			user.ID, user.Username, user.Email)
	}
}

// updateUser 更新用户信息
func updateUser(db *sql.DB, id int, newUsername, newEmail string) {
	result, err := db.Exec(
		"UPDATE users SET username = ?, email = ? WHERE id = ?",
		newUsername, newEmail, id,
	)
	if err != nil {
		log.Printf("更新用户失败: %v", err)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("更新了 %d 行数据\n", rowsAffected)
}

// deleteUser 删除用户
func deleteUser(db *sql.DB, id int) {
	result, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		log.Printf("删除用户失败: %v", err)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("删除了 %d 行数据\n", rowsAffected)
}

// transactionExample 事务示例
func transactionExample(db *sql.DB) {
	// 开始事务
	tx, err := db.Begin()
	if err != nil {
		log.Printf("开始事务失败: %v", err)
		return
	}

	// 使用defer和recover处理panic
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("事务panic并回滚: %v", r)
		}
	}()

	// 执行多个操作
	_, err = tx.Exec("INSERT INTO users (username, email) VALUES (?, ?)",
		"tx_user1", "tx1@example.com")
	if err != nil {
		tx.Rollback()
		log.Printf("事务中插入失败: %v", err)
		return
	}

	_, err = tx.Exec("INSERT INTO users (username, email) VALUES (?, ?)",
		"tx_user2", "tx2@example.com")
	if err != nil {
		tx.Rollback()
		log.Printf("事务中插入失败: %v", err)
		return
	}

	// 提交事务
	if err := tx.Commit(); err != nil {
		log.Printf("提交事务失败: %v", err)
		return
	}

	fmt.Println("事务成功提交")
}

// preparedStatementExample 预处理语句示例
func preparedStatementExample(db *sql.DB) {
	// 准备语句
	stmt, err := db.Prepare("INSERT INTO users (username, email) VALUES (?, ?)")
	if err != nil {
		log.Printf("准备语句失败: %v", err)
		return
	}
	defer stmt.Close()

	// 使用预处理语句执行多次
	users := []struct {
		username string
		email    string
	}{
		{"prep_user1", "prep1@example.com"},
		{"prep_user2", "prep2@example.com"},
		{"prep_user3", "prep3@example.com"},
	}

	for _, user := range users {
		_, err := stmt.Exec(user.username, user.email)
		if err != nil {
			log.Printf("执行预处理语句失败: %v", err)
			continue
		}
	}

	fmt.Println("使用预处理语句插入了3个用户")
}
