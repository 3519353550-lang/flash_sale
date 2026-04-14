package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"zgw/ks/flash_sale/user/configs"
	"zgw/ks/flash_sale/user/model"

	"zgw/ks/flash_sale/user/internal/svc"
	"zgw/ks/flash_sale/user/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsDetailedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGoodsDetailedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsDetailedLogic {
	return &GoodsDetailedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GoodsDetailedLogic) GoodsDetailed(in *users.GoodsDetailedRequest) (*users.GoodsDetailedResponse, error) {
	// todo: add your logic here and delete this line
	key := fmt.Sprintf("goods_detailed_%v", in.GoodId)

	var goods model.Goods

	var list *users.GoodsList
	result, _ := configs.Rdb.Get(l.ctx, key).Result()

	if result == "" {
		list = goods.GoodsDetailed(configs.DB, list, in)

		marshal, _ := json.Marshal(list)

		configs.Rdb.Set(l.ctx, key, marshal, time.Minute*60)

	} else {
		json.Unmarshal([]byte(result), &list)
	}

	return &users.GoodsDetailedResponse{
		Goods: list,
	}, nil
}
