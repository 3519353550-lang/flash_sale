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

type AddGoodLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddGoodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGoodLogic {
	return &AddGoodLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddGoodLogic) AddGood(in *users.AddGoodRequest) (*users.AddGoodResponse, error) {
	// todo: add your logic here and delete this line
	if in.UserId == 0 ||
		in.Status == 0 ||
		in.Price == 0.0 ||
		in.StockId == 0 ||
		in.Description == "" {
		return nil, errors.New("参数不能为空")
	}
	var m []*model.Image
	goods := model.Goods{
		//Model:       gorm.Model{},
		UserId:      in.UserId,
		Name:        in.Name,
		Price:       float64(in.Price),
		Status:      in.Status,
		Description: in.Description,
		StockId:     in.StockId,
		Types:       in.Types,
	}

	if err := goods.AddGood(configs.DB); err != nil {
		return nil, errors.New("添加商品失败")
	}
	if len(in.Image) > 5 {
		return nil, errors.New("商品图片数量不能超过5张")
	}
	for _, image := range in.Image {
		m = append(m, &model.Image{
			Url:     image.Url,
			GoodsId: int64(goods.ID),
		})
	}
	if err := configs.DB.Create(m).Error; err != nil {
		return nil, errors.New("添加商品图片失败")
	}
	GoodsMap := map[string]interface{}{
		"id":          goods.ID,
		"name":        goods.Name,
		"price":       goods.Price,
		"status":      goods.Status,
		"description": goods.Description,
		"stock_id":    goods.StockId,
		"types":       goods.Types,
	}
	_, err := configs.Esc.Index().
		Index("goods").
		BodyJson(GoodsMap).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	return &users.AddGoodResponse{Success: true}, nil
}
