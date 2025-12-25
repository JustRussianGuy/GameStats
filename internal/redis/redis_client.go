package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func InitRedis(host string, port int, db int) {
	RDB = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", host, port),
		DB:   db,
	})

	// Проверим подключение
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pong, err := RDB.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("failed to connect to Redis: %v", err))
	}
	fmt.Println("Redis connected:", pong)
}

// Функция для теста записи/чтения
func TestRedisSetGet() {
	ctx := context.Background()
	err := RDB.Set(ctx, "test_key", "ok", 10*time.Second).Err()
	if err != nil {
		fmt.Println("Redis SET error:", err)
		return
	}

	val, err := RDB.Get(ctx, "test_key").Result()
	if err != nil {
		fmt.Println("Redis GET error:", err)
	} else {
		fmt.Println("Redis GET test_key =", val)
	}
}
