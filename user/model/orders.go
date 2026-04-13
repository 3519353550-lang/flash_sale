package model

import (
	"gorm.io/gorm"
	"zgw/ks/flash_sale/user/users"
)

type Orders struct {
	gorm.Model
	OrderNo     string  `gorm:"type:varchar(50);not null;unique;comment:订单号"`
	UserID      int64   `gorm:"type:bigint;not null;index;comment:买家ID"`
	MerchantID  int64   `gorm:"type:bigint;not null;index;comment:商户ID"`
	TotalAmount float64 `gorm:"type:decimal(10,2);not null;comment:订单总金额"`
	Status      int64   `gorm:"type:tinyint;default:1;index;comment:1-待支付 2-待发货 3-待收货 4-完成 5-取消"`
}

func (o *Orders) OrderList(db *gorm.DB, list []*users.OrderList, in *users.OrderListRequest) []*users.OrderList {

	tx := db.Table("orders o").
		Joins("LEFT JOIN order_items oi ON o.id = oi.order_id").
		Where("o.deleted_at IS NULL")
	tx = tx.Where("o.status = ?", 2)
	if in.OrderNo != "" {
		tx = tx.Where("o.order_no = ?", in.OrderNo)
	}
	tx.Select(
		"o.id",
		"o.order_no",
		"o.user_id",
		"o.merchant_id",
		"o.total_amount",
		"o.status",
		"oi.product_id",
		"oi.product_name",
		"oi.product_image",
		"oi.price",
		"oi.num",
		"oi.total_price",
	).Find(&list)

	return list
}

type OrderItem struct {
	gorm.Model
	OrderID      int64   `gorm:"type:bigint;not null;index;comment:订单ID"`
	ProductID    int64   `gorm:"type:bigint;not null;index;comment:商品ID"`
	ProductName  string  `gorm:"type:varchar(200);not null;comment:商品名称"`
	ProductImage string  `gorm:"type:varchar(500);not null;comment:商品图片"`
	Price        float64 `gorm:"type:decimal(10,2);not null;comment:单价"`
	Num          int64   `gorm:"type:int;not null;comment:购买数量"`
	TotalPrice   float64 `gorm:"type:decimal(10,2);not null;comment:小计金额"`
}
