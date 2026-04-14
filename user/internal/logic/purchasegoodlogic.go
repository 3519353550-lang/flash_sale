package logic

import (
	"context"
	"errors"
	"fmt"
	"time"
	"zgw/ks/flash_sale/user/configs"
	"zgw/ks/flash_sale/user/model"
	"zgw/ks/flash_sale/user/pkg"

	"zgw/ks/flash_sale/user/internal/svc"
	"zgw/ks/flash_sale/user/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type PurchaseGoodLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPurchaseGoodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PurchaseGoodLogic {
	return &PurchaseGoodLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PurchaseGoodLogic) PurchaseGood(in *users.PurchaseGoodRequest) (*users.PurchaseGoodResponse, error) {
	// todo: add your logic here and delete this line
	var goods model.Goods
	var image model.Image
	var TotalAmount float64
	var orderItem []model.OrderItem
	if in.UserId == 0 {
		return nil, errors.New("用户ID未登录")
	}
	orderNo := pkg.GenerateOrderNo()
	for _, good := range in.PurchaseGoods {
		cacheKey := fmt.Sprintf("purchase:%v", good.GoodId)

		lock, _ := pkg.TryLock(configs.Rdb, cacheKey, 1*time.Second)
		if !lock {
			return nil, errors.New("商品已被购买完")
		}

		defer pkg.Unlock(configs.Rdb, cacheKey)

		if err := configs.DB.Where("id = ?", good.GoodId).First(&goods).Error; err != nil {
			return nil, errors.New("商品不存在")
		}

		if err := configs.DB.Where("goods_id = ?", goods.ID).First(&image).Error; err != nil {
			return nil, errors.New("商品图片不存在")
		}
		orderItem = append(orderItem, model.OrderItem{
			//Model:        gorm.Model{},
			//OrderID:      ,
			ProductID:    good.GoodId,
			ProductName:  goods.Name,
			ProductImage: image.Url,
			Price:        goods.Price,
			Num:          good.Quantity,
			TotalPrice:   goods.Price * float64(good.Quantity),
		})
		TotalAmount += goods.Price * float64(good.Quantity)

		if !pkg.GrabStock(int64(goods.ID), int(good.Quantity)) {
			return nil, errors.New("商品库存不足")
		}
	}

	order := model.Orders{
		OrderNo:     orderNo,
		UserID:      in.UserId,
		MerchantID:  goods.UserId,
		TotalAmount: TotalAmount,
		Status:      0,
	}
	tx := configs.DB.Begin()
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("创建订单失败")
	}
	for i, _ := range orderItem {
		orderItem[i].OrderID = int64(order.ID)
	}
	// 订单详细 item
	if err := tx.Create(&orderItem); err != nil {
		tx.Rollback()
		return nil, errors.New("创建订单详细item失败")
	}
	tx.Commit()

	return &users.PurchaseGoodResponse{
		OrderNo:    orderNo,
		TotalPrice: float32(TotalAmount),
	}, nil
}
