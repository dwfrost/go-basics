package database

import (
	"fmt"
	"log"
	"time"

	"github.com/glebarez/sqlite" // 使用纯Go实现的SQLite驱动
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GormUser 用于GORM示例的用户模型
type GormUser struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"size:100;not null;unique"`
	Email     string `gorm:"size:100;not null"`
	Age       int    `gorm:"default:18"`
	Active    bool   `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Posts     []Post         `gorm:"foreignKey:UserID"`
}

// Post 用于GORM示例的文章模型
type Post struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"size:200;not null"`
	Content   string `gorm:"type:text"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

// DemonstrateORM 展示GORM的使用
func DemonstrateORM() {
	fmt.Println("=== ORM数据库操作示例 ===")

	// 配置GORM - 使用纯Go实现的SQLite驱动
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("无法连接到数据库: %v", err)
	}

	// 自动迁移
	db.AutoMigrate(&GormUser{}, &Post{})
	fmt.Println("1. 成功创建表")

	// 创建用户
	createGormUsers(db)
	fmt.Println("\n2. 成功创建用户")

	// 查询用户
	queryGormUsers(db)

	// 更新用户
	updateGormUser(db)
	fmt.Println("\n4. 成功更新用户")

	// 删除用户
	deleteGormUser(db)
	fmt.Println("\n5. 成功删除用户")

	// 关联查询
	createPostsForUser(db)
	fmt.Println("\n6. 成功创建关联")

	// 查询关联
	queryUserWithPosts(db)
}

// createGormUsers 创建GORM用户
func createGormUsers(db *gorm.DB) {
	users := []GormUser{
		{Username: "gorm_user1", Email: "user1@example.com", Age: 25},
		{Username: "gorm_user2", Email: "user2@example.com", Age: 30},
		{Username: "gorm_user3", Email: "user3@example.com", Age: 35},
	}

	for _, user := range users {
		result := db.Create(&user)
		if result.Error != nil {
			log.Printf("创建用户失败: %v", result.Error)
			continue
		}
		fmt.Printf("创建用户: ID=%d, 用户名=%s\n", user.ID, user.Username)
	}
}

// queryGormUsers 查询GORM用户
func queryGormUsers(db *gorm.DB) {
	fmt.Println("\n3. 查询用户")

	// 查询所有用户
	var users []GormUser
	db.Find(&users)
	fmt.Printf("找到 %d 个用户:\n", len(users))
	for _, user := range users {
		fmt.Printf("- ID=%d, 用户名=%s, 邮箱=%s, 年龄=%d\n",
			user.ID, user.Username, user.Email, user.Age)
	}

	// 条件查询
	var user GormUser
	db.Where("username = ?", "gorm_user1").First(&user)
	fmt.Printf("\n条件查询: ID=%d, 用户名=%s, 邮箱=%s\n",
		user.ID, user.Username, user.Email)

	// 排序
	var olderUsers []GormUser
	db.Where("age > ?", 25).Order("age desc").Find(&olderUsers)
	fmt.Printf("\n年龄大于25的用户 (降序):\n")
	for _, u := range olderUsers {
		fmt.Printf("- ID=%d, 用户名=%s, 年龄=%d\n", u.ID, u.Username, u.Age)
	}
}

// updateGormUser 更新GORM用户
func updateGormUser(db *gorm.DB) {
	fmt.Println("\n4. 更新用户")
	// 方法1: 保存整个对象
	var user GormUser
	db.First(&user, 1)
	user.Email = "updated1@example.com"
	user.Age = 40
	db.Save(&user)
	fmt.Printf("更新用户1: ID=%d, 新邮箱=%s, 新年龄=%d\n",
		user.ID, user.Email, user.Age)

	// 方法2: 更新特定字段
	db.Model(&GormUser{}).Where("id = ?", 2).
		Updates(map[string]interface{}{"email": "updated2@example.com", "age": 45})
	var user2 GormUser
	db.First(&user2, 2)
	fmt.Printf("更新用户2: ID=%d, 新邮箱=%s, 新年龄=%d\n",
		user2.ID, user2.Email, user2.Age)
}

// deleteGormUser 删除GORM用户
func deleteGormUser(db *gorm.DB) {
	fmt.Println("\n5. 删除用户")
	// 软删除 (由于设置了gorm.DeletedAt)
	db.Delete(&GormUser{}, 3)
	fmt.Println("软删除用户ID=3")

	// 验证软删除
	var count int64
	db.Model(&GormUser{}).Count(&count)
	fmt.Printf("剩余用户数: %d\n", count)

	// 包括已删除的记录
	var totalCount int64
	db.Unscoped().Model(&GormUser{}).Count(&totalCount)
	fmt.Printf("包括已删除的总用户数: %d\n", totalCount)

	// 永久删除
	// db.Unscoped().Delete(&GormUser{}, 3)
}

// createPostsForUser 为用户创建文章
func createPostsForUser(db *gorm.DB) {
	fmt.Println("\n6. 创建文章")
	// 为用户1创建文章
	posts := []Post{
		{Title: "第一篇文章", Content: "这是用户1的第一篇文章", UserID: 1},
		{Title: "第二篇文章", Content: "这是用户1的第二篇文章", UserID: 1},
	}

	for _, post := range posts {
		db.Create(&post)
	}
	fmt.Printf("为用户ID=1创建了2篇文章\n")

	// 为用户2创建文章
	db.Create(&Post{
		Title:   "用户2的文章",
		Content: "这是用户2的文章",
		UserID:  2,
	})
	fmt.Printf("为用户ID=2创建了1篇文章\n")
}

// queryUserWithPosts 查询用户及其文章
func queryUserWithPosts(db *gorm.DB) {
	fmt.Println("\n7. 关联查询")

	// 预加载文章
	var users []GormUser
	db.Preload("Posts").Find(&users)

	for _, user := range users {
		fmt.Printf("用户: ID=%d, 用户名=%s\n", user.ID, user.Username)
		if len(user.Posts) > 0 {
			fmt.Printf("  文章数: %d\n", len(user.Posts))
			for _, post := range user.Posts {
				fmt.Printf("  - 标题: %s\n", post.Title)
			}
		} else {
			fmt.Println("  没有文章")
		}
	}

	// 联接查询
	type Result struct {
		UserID    uint
		Username  string
		PostCount int
	}

	var results []Result
	db.Model(&GormUser{}).
		Select("gorm_users.id as user_id, gorm_users.username, count(posts.id) as post_count").
		Joins("left join posts on posts.user_id = gorm_users.id").
		Group("gorm_users.id").
		Scan(&results)

	fmt.Println("\n用户文章统计:")
	for _, r := range results {
		fmt.Printf("- 用户ID=%d, 用户名=%s, 文章数=%d\n",
			r.UserID, r.Username, r.PostCount)
	}
}
