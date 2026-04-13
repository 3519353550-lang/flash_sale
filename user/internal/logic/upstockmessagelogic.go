package logic

import (
	"context"
	"errors"
	"log"
	"zgw/ks/flash_sale/user/configs"
	"zgw/ks/flash_sale/user/model"

	"zgw/ks/flash_sale/user/internal/svc"
	"zgw/ks/flash_sale/user/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpStockMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpStockMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpStockMessageLogic {
	return &UpStockMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpStockMessageLogic) UpStockMessage(in *users.UpStockMessageRequest) (*users.UpStockMessageResponse, error) {
	// todo: add your logic here and delete this line
	log.Printf("调用UpStockMessage接口，参数：%v", in)
	var stock model.Stocks

	if err := stock.FindStockById(configs.DB, in.StockId, in.GoodsId); err != nil {
		return nil, errors.New("库存不存在")
	}
	stock.StockNum = in.StockNum
	err := stock.UpStockMessage(configs.DB, in)
	if err != nil {
		return nil, err
	}
	log.Printf("更新成功")
	return &users.UpStockMessageResponse{
		Success: true,
	}, nil
}
