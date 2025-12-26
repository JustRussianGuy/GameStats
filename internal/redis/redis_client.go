package redis

import (
	"context"
	"fmt"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

var RDB *goredis.Client

func InitRedis(host string, port int, db int) {
	RDB = goredis.NewClient(&goredis.Options{
		Addr: fmt.Sprintf("%s:%d", host, port),
		DB:   db,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pong, err := RDB.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("failed to connect to Redis: %v", err))
	}

	fmt.Println("Redis connected:", pong)
	fmt.Println("Redis addr:", fmt.Sprintf("%s:%d", host, port))
}

func InvalidateByPattern(ctx context.Context, pattern string) error {
	iter := RDB.Scan(ctx, 0, pattern, 0).Iterator()
	for iter.Next(ctx) {
		if err := RDB.Del(ctx, iter.Val()).Err(); err != nil {
			return err
		}
	}
	return iter.Err()
}

