package pkg

import (
	"crypto/md5"
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"gorm.io/gorm"
	"regexp"
	"strconv"
	"time"
	"zgw/ks/flash_sale/user/configs"
)

func IsMobile(mobile string) bool {
	matchString, err := regexp.MatchString("^[1][3456789]\\d{9}$", mobile)
	if err != nil {
		return false
	}
	return matchString
}

func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if page <= 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

/*
	定时任务
	30分钟未支付订单，自动释放库存
*/

func ScheduledTask() {
	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()
	for {
		<-ticker.C

		fmt.Println("定时任务开始")

		if configs.DB == nil {
			fmt.Println("数据库未连接")
			continue
		}

		rows, err := configs.DB.Raw(`SELECT user_id FROM orders WHERE status = 1  AND created_at < DATE_SUB(NOW(), INTERVAL 30 MINUTE)`).Rows()
		if err != nil {
			fmt.Println("查询错误")
			return
		}
		count := 0
		for rows.Next() {
			var goodsId int64

			if err := rows.Scan(&goodsId); err != nil {
				return
			}
			ReleaseStock(goodsId)
			count++
		}
		rows.Close()

		tx := configs.DB.Exec(`UPDATE orders SET status  = 5 WHERE status = 1 AND created_at < DATE_SUB(NOW(), INTERVAL 30 MINUTE)`)
		if tx.Error != nil {
			fmt.Println("更新订单状态失败")
			return
		} else {
			fmt.Println("更新订单数量：", count, "状态成功", tx.RowsAffected)
		}
	}

}

func AliPay(orderNo string, TotalAmount float64) {
	pay := configs.Conf.AilPay
	var privateKey = pay.PrivateKey // 必须，上一步中使用 RSA签名验签工具 生成的私钥
	client, err := alipay.New(pay.AppId, privateKey, false)
	if err != nil {

		return
	}

	var p = alipay.TradePagePay{}
	p.NotifyURL = pay.NotifyUrl
	p.ReturnURL = pay.ReturnUrl
	p.Subject = "标题"
	p.OutTradeNo = orderNo
	p.TotalAmount = strconv.FormatFloat(TotalAmount, 'f', 2, 64)
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, err := client.TradePagePay(p)
	if err != nil {
		return
	}

	// 这个 payURL 即是用于打开支付宝支付页面的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	var payURL = url.String()
	fmt.Println(payURL)
}
