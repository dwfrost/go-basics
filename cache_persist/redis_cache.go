package cache_persist

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisCache 结构体
type RedisCache struct {
	client *redis.Client
	ctx    context.Context
}

// NewRedisCache 创建一个新的Redis缓存实例
func NewRedisCache(addr string, password string, db int) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisCache{
		client: client,
		ctx:    context.Background(),
	}
}

// Set 设置缓存
func (c *RedisCache) Set(key string, value string, ttl time.Duration) error {
	return c.client.Set(c.ctx, key, value, ttl).Err()
}

// Get 获取缓存
func (c *RedisCache) Get(key string) (string, error) {
	return c.client.Get(c.ctx, key).Result()
}

// Delete 删除缓存
func (c *RedisCache) Delete(key string) error {
	return c.client.Del(c.ctx, key).Err()
}

// Clear 清空所有缓存
func (c *RedisCache) Clear() error {
	return c.client.FlushDB(c.ctx).Err()
}

// Close 关闭Redis连接
func (c *RedisCache) Close() error {
	return c.client.Close()
}

func main() {
	// 创建Redis缓存实例
	cache := NewRedisCache("localhost:6379", "", 0)
	defer cache.Close()

	// 设置缓存
	err := cache.Set("test_key", "test_value", 5*time.Second)
	if err != nil {
		fmt.Printf("设置缓存失败: %v\n", err)
		return
	}

	// 获取缓存
	value, err := cache.Get("test_key")
	if err != nil {
		fmt.Printf("获取缓存失败: %v\n", err)
		return
	}
	fmt.Printf("获取到的缓存值: %s\n", value)

	// 等待缓存过期
	time.Sleep(6 * time.Second)

	// 再次尝试获取已过期的缓存
	value, err = cache.Get("test_key")
	if err == redis.Nil {
		fmt.Println("缓存已过期")
	} else if err != nil {
		fmt.Printf("获取缓存失败: %v\n", err)
	} else {
		fmt.Printf("获取到的缓存值: %s\n", value)
	}
}
