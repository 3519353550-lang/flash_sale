package pkg

import (
	"context"
	"fmt"
	"zgw/ks/flash_sale/user/configs"
)

func PreheatStock() {
	// 1. 检查数据库连接是否初始化
	if configs.DB == nil {
		fmt.Println("Database connection not initialized")
		return
	}

	// 2. 检查Redis连接是否初始化
	if configs.Rdb == nil {
		fmt.Println("Redis connection not initialized")
		return
	}
	rows, err := configs.DB.Raw("SELECT goods_id,stock_num FROM stocks").Rows()
	if err != nil {
		// 查询失败时直接返回，不进行后续操作
		fmt.Printf("Error querying seckill goods: %v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var goodsId int64 // 商品ID
		var stockNum int  // 商品库存数量
		err = rows.Scan(&goodsId, &stockNum)
		if err != nil {
			fmt.Printf("Error scanning row: %v\n", err)
			return
		}
		if stockNum < 0 {
			continue
		}
		key := fmt.Sprintf("seckill:stock:queue:%d", goodsId)
		configs.Rdb.Del(context.Background(), key)

		for i := 0; i < stockNum; i++ {
			configs.Rdb.LPush(context.Background(), key, "1")
		}

	}

}
func GrabStock(goodsId int64, quantity int) bool {
	var err error
	key := fmt.Sprintf("seckill:stock:queue:%d", goodsId)

	// 从队列右侧取一个（原子操作，不超卖）
	for i := 0; i < quantity; i++ {
		_, err = configs.Rdb.RPop(context.Background(), key).Result()
	}
	return err == nil
}

// 释放库存
func ReleaseStock(goodsId int64) {
	key := fmt.Sprintf("seckill:stock:queue:%d", goodsId)
	// 还回队列头部
	configs.Rdb.LPush(context.Background(), key, "1")
}
