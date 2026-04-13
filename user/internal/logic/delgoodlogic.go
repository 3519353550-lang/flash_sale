package logic

import (
	"context"
	"errors"
	"zgw/ks/flash_sale/user/configs"
	"zgw/ks/flash_sale/user/model"

	"zgw/ks/flash_sale/user/internal/svc"
	"zgw/ks/flash_sale/user/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelGoodLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelGoodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelGoodLogic {
	return &DelGoodLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelGoodLogic) DelGood(in *users.DelGoodRequest) (*users.DelGoodResponse, error) {
	// todo: add your logic here and delete this line
	if in.UserId == 0 ||
		in.GoodsId == 0 {
		return nil, errors.New("参数不能为空")
	}

	if err := configs.DB.Where("user_id = ? and id = ?", in.UserId, in.GoodsId).Limit(1).Delete(&model.Goods{}).Error; err != nil {
		return nil, errors.New("删除商品失败")
	}
	return &users.DelGoodResponse{
		Success: true,
	}, nil
}
