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

type IsHotLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsHotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsHotLogic {
	return &IsHotLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsHotLogic) IsHot(in *users.IsHotRequest) (*users.IsHotResponse, error) {
	// todo: add your logic here and delete this line

	var goods model.Goods

	if err := goods.IsHots(configs.DB, in); err != nil {
		return nil, errors.New("修改失败，请稍后重试")
	}

	return &users.IsHotResponse{
		Success: true,
	}, nil
}
