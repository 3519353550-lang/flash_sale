package logic

import (
	"context"
	"zgw/ks/flash_sale/user/configs"
	"zgw/ks/flash_sale/user/model"

	"zgw/ks/flash_sale/user/internal/svc"
	"zgw/ks/flash_sale/user/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGoodsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsListLogic {
	return &GoodsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GoodsListLogic) GoodsList(in *users.GoodsListRequest) (*users.GoodsListResponse, error) {
	// todo: add your logic here and delete this line

	var goods model.Goods

	var list []*users.GoodsList

	list = goods.GoodsList(configs.DB, list, in)

	return &users.GoodsListResponse{
		UserList: list,
	}, nil
}
