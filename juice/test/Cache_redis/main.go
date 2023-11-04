package main

import (
	"fmt"
	"time"

	"juice/test/config"

	"github.com/go-redis/redis"
)

// 创建redis客户端
func newClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Cache_address,  // redis地址
		Password: config.Cache_password, // 密码
		DB:       config.Cache_database, // 使用默认数据库
	})
	return client
}

func main() {

	// 创建客户端
	client := newClient()
	defer client.Close()

	// 设置key
	err := client.Set("name", "john", 10*time.Second).Err()
	if err != nil {
		panic(err)
	}

	// 获取key
	val, err := client.Get("name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name", val)

}
