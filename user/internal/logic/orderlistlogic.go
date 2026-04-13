package logic

import (
	"context"
	"zgw/ks/flash_sale/user/configs"
	"zgw/ks/flash_sale/user/model"

	"zgw/ks/flash_sale/user/internal/svc"
	"zgw/ks/flash_sale/user/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderListLogic {
	return &OrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderListLogic) OrderList(in *users.OrderListRequest) (*users.OrderListResponse, error) {
	// todo: add your logic here and delete this line

	var order model.Orders

	var list []*users.OrderList

	list = order.OrderList(configs.DB, list, in)

	return &users.OrderListResponse{
		OrderList: list,
	}, nil
}
