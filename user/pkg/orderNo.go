package pkg

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	randObj = rand.New(rand.NewSource(time.Now().UnixNano()))
	mu      sync.Mutex
)

// GenerateOrderNo 生成订单号：时间戳 + 随机数
func GenerateOrderNo() string {
	// 14位日期时间：20260414225810
	now := time.Now().Format("20060102150405")

	mu.Lock()
	// 4位随机数 0001~9999
	seq := randObj.Intn(9999)
	mu.Unlock()

	return fmt.Sprintf("%s%04d", now, seq)
}
