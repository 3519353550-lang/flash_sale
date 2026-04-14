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

type SearchGoodsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchGoodsLogic {
	return &SearchGoodsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchGoodsLogic) SearchGoods(req *types.SearchGoodsRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	data, err := l.svcCtx.UserRpc.SearchGoods(l.ctx, &users.SearchGoodsRequest{
		Name:     req.Name,
		Types:    req.Types,
		MaxPrice: float32(req.MaxPrice),
		MinPrice: float32(req.MinPrice),
		Sales:    req.Sales,
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
