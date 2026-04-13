package inits

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"zgw/ks/flash_sale/user/configs"
)

var ctx = context.Background()

func InitRedis() {
	RC := configs.Conf.Redis
	Addr := fmt.Sprintf("%s:%d", RC.Host, RC.Port)
	configs.Rdb = redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: RC.Password, // no password set
		DB:       RC.Database, // use default DB
	})

	if err = configs.Rdb.Ping(ctx).Err(); err != nil {
		fmt.Println("Redis数据库连接失败", err)
		return
	}
	fmt.Println("Redis数据库连接成功")

}
