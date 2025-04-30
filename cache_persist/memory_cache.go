package cache_persist

import (
	"fmt"
	"sync"
	"time"
)

// MemoryCache 结构体
type MemoryCache struct {
	cache    map[string]*cacheItem
	mutex    sync.RWMutex
	stopChan chan struct{}
}

// cacheItem 表示缓存中的一个项目
type cacheItem struct {
	value      string
	expiration time.Time
}

// NewMemoryCache 创建一个新的内存缓存实例
func NewMemoryCache() *MemoryCache {
	cache := &MemoryCache{
		cache:    make(map[string]*cacheItem),
		stopChan: make(chan struct{}),
	}

	return cache
}

// StartCleaner 启动清理过期缓存的协程
func (c *MemoryCache) StartCleaner(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for {
			select {
			case <-ticker.C:
				c.cleanExpired()
			case <-c.stopChan:
				ticker.Stop()
				return
			}
		}
	}()
}

// StopCleaner 停止清理过期缓存的协程
func (c *MemoryCache) StopCleaner() {
	close(c.stopChan)
}

// Set 设置缓存
func (c *MemoryCache) Set(key string, value string, ttl time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	expiration := time.Now().Add(ttl)
	c.cache[key] = &cacheItem{
		value:      value,
		expiration: expiration,
	}
}

// Get 获取缓存
func (c *MemoryCache) Get(key string) (string, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	item, found := c.cache[key]
	if !found {
		return "", false
	}

	// 检查是否过期
	if time.Now().After(item.expiration) {
		return "", false
	}

	return item.value, true
}

// Delete 删除缓存
func (c *MemoryCache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.cache, key)
}

// Clear 清空所有缓存
func (c *MemoryCache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cache = make(map[string]*cacheItem)
}

// cleanExpired 清理过期的缓存项
func (c *MemoryCache) cleanExpired() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	now := time.Now()
	for key, item := range c.cache {
		if now.After(item.expiration) {
			delete(c.cache, key)
		}
	}
}

func DemonstrateMemoryCache() {
	// 创建内存缓存实例
	cache := NewMemoryCache()

	// 开始定期清理缓存，每秒执行一次
	cache.StartCleaner(time.Second)
	defer cache.StopCleaner()

	// 设置缓存
	cache.Set("test_key", "test_value", 5*time.Second)
	fmt.Println("缓存已设置: test_key=test_value (5秒过期)")

	// 获取缓存
	value, found := cache.Get("test_key")
	if found {
		fmt.Printf("获取到的缓存值: %s\n", value)
	} else {
		fmt.Println("缓存不存在")
	}

	// 等待缓存过期
	fmt.Println("等待6秒...")
	time.Sleep(6 * time.Second)

	// 再次尝试获取已过期的缓存
	value, found = cache.Get("test_key")
	if found {
		fmt.Printf("获取到的缓存值: %s\n", value)
	} else {
		fmt.Println("缓存已过期或不存在")
	}

	// 测试删除功能
	cache.Set("to_delete", "delete_me", 30*time.Second)
	fmt.Println("设置新缓存: to_delete=delete_me")

	cache.Delete("to_delete")
	fmt.Println("删除缓存项: to_delete")

	_, found = cache.Get("to_delete")
	if !found {
		fmt.Println("缓存项已成功删除")
	}
}
