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

type IsHotLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIsHotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsHotLogic {
	return &IsHotLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsHotLogic) IsHot(req *types.IsHotRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	data, err := l.svcCtx.UserRpc.IsHot(l.ctx, &users.IsHotRequest{
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
