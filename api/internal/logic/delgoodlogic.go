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

type DelGoodLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelGoodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelGoodLogic {
	return &DelGoodLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelGoodLogic) DelGood(req *types.DelGoodRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	data, err := l.svcCtx.UserRpc.DelGood(l.ctx, &users.DelGoodRequest{
		UserId:  req.UserId,
		GoodsId: req.GoodsId,
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
