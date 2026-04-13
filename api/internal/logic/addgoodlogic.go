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

type AddGoodLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddGoodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGoodLogic {
	return &AddGoodLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddGoodLogic) AddGood(req *types.AddGoodRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	var images []*users.Image
	for _, image := range req.Images {
		images = append(images, &users.Image{
			Url:     image.Url,
			GoodsId: image.GoodsId,
		})
	}

	data, err := l.svcCtx.UserRpc.AddGood(l.ctx, &users.AddGoodRequest{
		UserId:      req.UserId,
		Name:        req.Name,
		Price:       float32(req.Price),
		Status:      req.Status,
		Description: req.Description,
		StockId:     req.StockId,
		Types:       req.Types,
		Image:       images,
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
