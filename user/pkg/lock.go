package pkg

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

// 加锁（SET NX EX）
func TryLock(rdb *redis.Client, lockKey string, expire time.Duration) (bool, error) {
	return rdb.SetNX(context.Background(), lockKey, "1", expire).Result()

}

// 解锁（Lua 原子脚本，防止误删）
func Unlock(rdb *redis.Client, lockKey string) error {
	script := `
	if redis.call("get", KEYS[1]) == ARGV[1] then
		return redis.call("del", KEYS[1])
	else
		return 0
	end
	`

	res, err := rdb.Eval(context.Background(), script, []string{lockKey}, "1").Result()
	if err != nil {
		return err
	}
	if i, ok := res.(int64); !ok || i == 0 {
		return errors.New("解锁失败")
	}
	return nil
}
