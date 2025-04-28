package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/glebarez/sqlite" // 纯Go实现的SQLite驱动
	"gorm.io/gorm"
)

// DemonstratePoolAndTransaction 展示连接池和事务处理
func DemonstratePoolAndTransaction() {
	fmt.Println("=== 数据库连接池和事务处理示例 ===")

	// // 原生SQL连接池示例
	// fmt.Println("\n1. 原生SQL连接池")
	// demonstrateSQLPool()

	// // 原生SQL事务示例
	// fmt.Println("\n2. 原生SQL事务处理")
	// demonstrateSQLTransaction()

	// GORM连接池示例
	fmt.Println("\n3. GORM连接池")
	demonstrateGORMPool()

	// GORM事务示例
	fmt.Println("\n4. GORM事务处理")
	demonstrateGORMTransaction()
}

// demonstrateSQLPool 展示原生SQL连接池
func demonstrateSQLPool() {
	// 打开数据库连接
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatalf("无法打开数据库: %v", err)
	}
	defer db.Close()

	// 配置连接池
	db.SetMaxOpenConns(25)                  // 最大打开连接数
	db.SetMaxIdleConns(10)                  // 最大空闲连接数
	db.SetConnMaxLifetime(5 * time.Minute)  // 连接最大生命周期
	db.SetConnMaxIdleTime(10 * time.Minute) // 空闲连接最大生命周期

	// 创建测试表
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS pool_test (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		data TEXT
	)`)
	if err != nil {
		log.Fatalf("创建表失败: %v", err)
	}

	// 展示连接池状态
	printPoolStats(db)

	// 模拟高并发请求
	var wg sync.WaitGroup
	start := time.Now()

	// 启动20个并发goroutine
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// 模拟数据库操作
			_, err := db.Exec("INSERT INTO pool_test (data) VALUES (?)",
				fmt.Sprintf("数据 %d", id))
			if err != nil {
				log.Printf("插入失败: %v", err)
			}
			// 模拟一些处理时间
			time.Sleep(100 * time.Millisecond)
		}(i)
	}

	// 等待所有goroutine完成
	wg.Wait()
	elapsed := time.Since(start)

	// 查询结果
	var count int
	db.QueryRow("SELECT COUNT(*) FROM pool_test").Scan(&count)
	fmt.Printf("插入了 %d 条记录\n", count)
	fmt.Printf("总耗时: %v\n", elapsed)

	// 再次展示连接池状态
	printPoolStats(db)

	// 连接池最佳实践
	fmt.Println("\n连接池最佳实践:")
	fmt.Println("1. 根据应用负载和数据库服务器能力设置合适的连接数")
	fmt.Println("2. 监控连接池状态，及时调整参数")
	fmt.Println("3. 设置合理的连接超时和生命周期")
	fmt.Println("4. 使用context控制查询超时")
	fmt.Println("5. 总是关闭查询结果集(rows.Close())以释放连接")
}

// printPoolStats 打印连接池状态
func printPoolStats(db *sql.DB) {
	stats := db.Stats()
	fmt.Println("\n连接池状态:")
	fmt.Printf("打开的连接数: %d\n", stats.OpenConnections)
	fmt.Printf("使用中的连接数: %d\n", stats.InUse)
	fmt.Printf("空闲连接数: %d\n", stats.Idle)
	fmt.Printf("等待的连接请求数: %d\n", stats.WaitCount)
	fmt.Printf("最大等待时间: %v\n", stats.WaitDuration)
}

// demonstrateSQLTransaction 展示原生SQL事务处理
func demonstrateSQLTransaction() {
	// 打开数据库连接
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatalf("无法打开数据库: %v", err)
	}
	defer db.Close()

	// 创建测试表
	_, err = db.Exec(`CREATE TABLE accounts (
		id INTEGER PRIMARY KEY,
		name TEXT,
		balance REAL
	)`)
	if err != nil {
		log.Fatalf("创建表失败: %v", err)
	}

	// 插入初始数据
	_, err = db.Exec(`INSERT INTO accounts (id, name, balance) VALUES 
		(1, '张三', 1000), 
		(2, '李四', 500)`)
	if err != nil {
		log.Fatalf("插入数据失败: %v", err)
	}

	// 打印初始余额
	printBalances(db)

	// 1. 基本事务 - 转账成功
	fmt.Println("\n1.1 基本事务 - 转账成功示例")
	transferMoney(db, 1, 2, 200, false)
	printBalances(db)

	// 2. 事务回滚 - 转账失败
	fmt.Println("\n1.2 事务回滚 - 转账失败示例")
	transferMoney(db, 1, 2, 2000, false) // 余额不足
	printBalances(db)

	// 3. 带有上下文的事务
	fmt.Println("\n1.3 带有上下文的事务")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	transferMoneyWithContext(ctx, db, 2, 1, 100)
	printBalances(db)

	// 4. 嵌套事务 (SQLite不支持真正的嵌套事务，这里模拟实现)
	fmt.Println("\n1.4 嵌套事务 (模拟)")
	nestedTransactionDemo(db)
	printBalances(db)

	// 5. 事务隔离级别
	fmt.Println("\n1.5 事务隔离级别")
	fmt.Println("SQLite默认隔离级别: SERIALIZABLE (最高)")
	fmt.Println("常见隔离级别从低到高:")
	fmt.Println("- READ UNCOMMITTED: 可能读取到未提交的数据(脏读)")
	fmt.Println("- READ COMMITTED: 只能读取已提交的数据")
	fmt.Println("- REPEATABLE READ: 确保事务内多次读取结果一致")
	fmt.Println("- SERIALIZABLE: 最高级别，完全隔离")

	// 6. 事务最佳实践
	fmt.Println("\n1.6 事务最佳实践:")
	fmt.Println("1. 事务应尽可能短")
	fmt.Println("2. 避免在事务中进行耗时操作")
	fmt.Println("3. 正确处理错误和回滚")
	fmt.Println("4. 使用适当的隔离级别")
	fmt.Println("5. 考虑使用乐观锁或悲观锁处理并发")
	fmt.Println("6. 使用defer确保事务正确结束")
}

// printBalances 打印所有账户余额
func printBalances(db *sql.DB) {
	fmt.Println("当前账户余额:")
	rows, err := db.Query("SELECT id, name, balance FROM accounts")
	if err != nil {
		log.Printf("查询失败: %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var balance float64
		rows.Scan(&id, &name, &balance)
		fmt.Printf("账户 %d (%s): %.2f\n", id, name, balance)
	}
}

// transferMoney 转账函数
func transferMoney(db *sql.DB, fromID, toID int, amount float64, slowMode bool) {
	// 开始事务
	tx, err := db.Begin()
	if err != nil {
		log.Printf("开始事务失败: %v", err)
		return
	}

	// 使用defer和recover确保事务正确结束
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("事务panic并回滚: %v", r)
		}
	}()

	// 检查余额是否足够
	var balance float64
	err = tx.QueryRow("SELECT balance FROM accounts WHERE id = ?", fromID).Scan(&balance)
	if err != nil {
		tx.Rollback()
		log.Printf("查询余额失败: %v", err)
		return
	}

	if balance < amount {
		tx.Rollback()
		log.Printf("余额不足: 当前%.2f, 需要%.2f", balance, amount)
		return
	}

	// 模拟慢速操作，用于演示长事务的问题
	if slowMode {
		fmt.Println("事务中执行慢速操作...")
		time.Sleep(2 * time.Second)
	}

	// 扣除发送方余额
	_, err = tx.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, fromID)
	if err != nil {
		tx.Rollback()
		log.Printf("扣款失败: %v", err)
		return
	}

	// 增加接收方余额
	_, err = tx.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, toID)
	if err != nil {
		tx.Rollback()
		log.Printf("加款失败: %v", err)
		return
	}

	// 提交事务
	err = tx.Commit()
	if err != nil {
		log.Printf("提交事务失败: %v", err)
		return
	}

	fmt.Printf("成功从账户%d转账%.2f到账户%d\n", fromID, amount, toID)
}

// transferMoneyWithContext 带上下文的转账函数
func transferMoneyWithContext(ctx context.Context, db *sql.DB, fromID, toID int, amount float64) {
	// 开始事务
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("开始事务失败: %v", err)
		return
	}

	// 使用defer确保事务正确结束
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("事务panic并回滚: %v", r)
		}
	}()

	// 检查上下文是否已取消
	select {
	case <-ctx.Done():
		tx.Rollback()
		log.Printf("事务上下文已取消: %v", ctx.Err())
		return
	default:
		// 继续执行
	}

	// 检查余额
	var balance float64
	err = tx.QueryRowContext(ctx, "SELECT balance FROM accounts WHERE id = ?", fromID).Scan(&balance)
	if err != nil {
		tx.Rollback()
		log.Printf("查询余额失败: %v", err)
		return
	}

	if balance < amount {
		tx.Rollback()
		log.Printf("余额不足: 当前%.2f, 需要%.2f", balance, amount)
		return
	}

	// 扣除发送方余额
	_, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, fromID)
	if err != nil {
		tx.Rollback()
		log.Printf("扣款失败: %v", err)
		return
	}

	// 增加接收方余额
	_, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, toID)
	if err != nil {
		tx.Rollback()
		log.Printf("加款失败: %v", err)
		return
	}

	// 提交事务
	err = tx.Commit()
	if err != nil {
		log.Printf("提交事务失败: %v", err)
		return
	}

	fmt.Printf("成功从账户%d转账%.2f到账户%d (带上下文)\n", fromID, amount, toID)
}

// nestedTransactionDemo 嵌套事务示例
func nestedTransactionDemo(db *sql.DB) {
	// 开始外部事务
	outerTx, err := db.Begin()
	if err != nil {
		log.Printf("开始外部事务失败: %v", err)
		return
	}

	// 更新第一个账户
	_, err = outerTx.Exec("UPDATE accounts SET balance = balance - 50 WHERE id = 1")
	if err != nil {
		outerTx.Rollback()
		log.Printf("外部事务更新失败: %v", err)
		return
	}

	fmt.Println("外部事务: 从账户1扣除50")

	// 创建保存点 (SQLite支持)
	_, err = outerTx.Exec("SAVEPOINT nested_tx")
	if err != nil {
		outerTx.Rollback()
		log.Printf("创建保存点失败: %v", err)
		return
	}

	// 模拟内部事务操作
	_, err = outerTx.Exec("UPDATE accounts SET balance = balance + 30 WHERE id = 2")
	if err != nil {
		outerTx.Exec("ROLLBACK TO SAVEPOINT nested_tx")
		log.Printf("内部事务更新失败并回滚: %v", err)
		// 外部事务可以继续
	} else {
		fmt.Println("内部事务: 向账户2添加30")
		// 释放保存点
		outerTx.Exec("RELEASE SAVEPOINT nested_tx")
	}

	// 提交外部事务
	err = outerTx.Commit()
	if err != nil {
		log.Printf("提交外部事务失败: %v", err)
		return
	}

	fmt.Println("嵌套事务示例完成")
}

// demonstrateGORMPool 展示GORM连接池
func demonstrateGORMPool() {
	// 配置GORM
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalf("无法连接到数据库: %v", err)
	}

	// 获取底层的sql.DB对象
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("获取底层DB失败: %v", err)
	}

	// 配置连接池
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 创建测试模型
	type PoolTest struct {
		ID   uint `gorm:"primarykey"`
		Data string
	}

	// 自动迁移
	db.AutoMigrate(&PoolTest{})

	// 模拟并发请求
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// 创建记录
			db.Create(&PoolTest{Data: fmt.Sprintf("GORM数据 %d", id)})
			time.Sleep(100 * time.Millisecond)
		}(i)
	}

	wg.Wait()
	elapsed := time.Since(start)

	// 查询结果
	var count int64
	db.Model(&PoolTest{}).Count(&count)
	fmt.Printf("GORM插入了 %d 条记录\n", count)
	fmt.Printf("总耗时: %v\n", elapsed)

	// 打印连接池状态
	stats := sqlDB.Stats()
	fmt.Println("\nGORM连接池状态:")
	fmt.Printf("打开的连接数: %d\n", stats.OpenConnections)
	fmt.Printf("使用中的连接数: %d\n", stats.InUse)
	fmt.Printf("空闲连接数: %d\n", stats.Idle)

	fmt.Println("\nGORM连接池最佳实践:")
	fmt.Println("1. 通常只需要一个全局的gorm.DB实例")
	fmt.Println("2. 根据应用需求配置底层连接池")
	fmt.Println("3. 使用db.WithContext()传递上下文")
	fmt.Println("4. 监控连接池状态")
}

// demonstrateGORMTransaction 展示GORM事务处理
func demonstrateGORMTransaction() {
	// 配置GORM
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalf("无法连接到数据库: %v", err)
	}

	// 创建账户模型
	type Account struct {
		ID      uint `gorm:"primarykey"`
		Name    string
		Balance float64
	}

	// 自动迁移
	db.AutoMigrate(&Account{})

	// 插入初始数据
	db.Create(&Account{ID: 1, Name: "张三", Balance: 1000})
	db.Create(&Account{ID: 2, Name: "李四", Balance: 500})

	// 打印初始余额
	printGORMBalances(db)

	// 1. 基本事务
	fmt.Println("\n4.1 GORM基本事务")
	transferMoneyGORM(db, 1, 2, 200)
	printGORMBalances(db)

	// 2. 手动事务
	fmt.Println("\n4.2 GORM手动事务")
	manualGORMTransaction(db)
	printGORMBalances(db)

	// 3. 事务闭包
	fmt.Println("\n4.3 GORM事务闭包")
	db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行操作
		var account Account
		if err := tx.First(&account, 1).Error; err != nil {
			return err
		}

		// 更新账户
		account.Balance -= 100
		if err := tx.Save(&account).Error; err != nil {
			return err
		}

		// 更新另一个账户
		if err := tx.Model(&Account{}).Where("id = ?", 2).
			Update("balance", gorm.Expr("balance + ?", 100)).Error; err != nil {
			return err
		}

		fmt.Println("GORM事务闭包: 从账户1转账100到账户2")
		return nil // 返回nil提交事务
	})
	printGORMBalances(db)

	// 4. 嵌套事务
	fmt.Println("\n4.4 GORM嵌套事务")
	db.Transaction(func(tx *gorm.DB) error {
		// 外部事务
		tx.Model(&Account{}).Where("id = ?", 1).
			Update("balance", gorm.Expr("balance - ?", 50))
		fmt.Println("外部事务: 从账户1扣除50")

		// 嵌套事务
		return tx.Transaction(func(tx2 *gorm.DB) error {
			if err := tx2.Model(&Account{}).Where("id = ?", 2).
				Update("balance", gorm.Expr("balance + ?", 50)).Error; err != nil {
				return err
			}
			fmt.Println("内部事务: 向账户2添加50")
			return nil
		})
	})
	printGORMBalances(db)

	// 5. 事务最佳实践
	fmt.Println("\n4.5 GORM事务最佳实践:")
	fmt.Println("1. 优先使用事务闭包，自动处理提交和回滚")
	fmt.Println("2. 在事务中使用tx变量而非原始db")
	fmt.Println("3. 正确处理和返回错误")
	fmt.Println("4. 使用tx.WithContext()传递上下文")
	fmt.Println("5. 避免大事务，保持事务简短")
}

// printGORMBalances 打印GORM账户余额
func printGORMBalances(db *gorm.DB) {
	type Account struct {
		ID      uint
		Name    string
		Balance float64
	}

	var accounts []Account
	db.Find(&accounts)

	fmt.Println("当前GORM账户余额:")
	for _, acc := range accounts {
		fmt.Printf("账户 %d (%s): %.2f\n", acc.ID, acc.Name, acc.Balance)
	}
}

// transferMoneyGORM GORM转账函数
func transferMoneyGORM(db *gorm.DB, fromID, toID uint, amount float64) {
	type Account struct {
		ID      uint
		Name    string
		Balance float64
	}

	// 使用事务闭包
	err := db.Transaction(func(tx *gorm.DB) error {
		// 检查余额
		var fromAccount Account
		if err := tx.First(&fromAccount, fromID).Error; err != nil {
			return err
		}

		if fromAccount.Balance < amount {
			return fmt.Errorf("余额不足: 当前%.2f, 需要%.2f", fromAccount.Balance, amount)
		}

		// 扣除发送方余额
		if err := tx.Model(&Account{}).Where("id = ?", fromID).
			Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
			return err
		}

		// 增加接收方余额
		if err := tx.Model(&Account{}).Where("id = ?", toID).
			Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Printf("GORM转账失败: %v", err)
		return
	}

	fmt.Printf("GORM成功从账户%d转账%.2f到账户%d\n", fromID, amount, toID)
}

// manualGORMTransaction GORM手动事务示例
func manualGORMTransaction(db *gorm.DB) {
	type Account struct {
		ID      uint
		Name    string
		Balance float64
	}

	// 开始事务
	tx := db.Begin()
	if tx.Error != nil {
		log.Printf("开始事务失败: %v", tx.Error)
		return
	}

	// 使用defer确保事务结束
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("事务panic并回滚: %v", r)
		}
	}()

	// 执行事务操作
	if err := tx.Model(&Account{}).Where("id = ?", 1).
		Update("balance", gorm.Expr("balance - ?", 150)).Error; err != nil {
		tx.Rollback()
		log.Printf("更新账户1失败: %v", err)
		return
	}

	if err := tx.Model(&Account{}).Where("id = ?", 2).
		Update("balance", gorm.Expr("balance + ?", 150)).Error; err != nil {
		tx.Rollback()
		log.Printf("更新账户2失败: %v", err)
		return
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		log.Printf("提交事务失败: %v", err)
		return
	}

	fmt.Println("GORM手动事务: 从账户1转账150到账户2")
}
