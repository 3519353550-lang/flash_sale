// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"zgw/ks/flash_sale/user/users"

	"zgw/ks/flash_sale/api/internal/svc"
	"zgw/ks/flash_sale/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsDetailedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsDetailedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsDetailedLogic {
	return &GoodsDetailedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsDetailedLogic) GoodsDetailed(req *types.GoodsDetailedRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	data, err := l.svcCtx.UserRpc.GoodsDetailed(l.ctx, &users.GoodsDetailedRequest{
		GoodId: req.GoodsId,
	})
	if err != nil {
		return &types.Response{
			Data: nil,
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	return &types.Response{
		Data: data,
		Code: 200,
		Msg:  "success",
	}, nil
}
