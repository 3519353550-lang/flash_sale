package model

import (
	"gorm.io/gorm"
	"zgw/ks/flash_sale/user/pkg"
	"zgw/ks/flash_sale/user/users"
)

type Goods struct {
	gorm.Model
	UserId      int64   `gorm:"type:int(11);not null;comment:用户ID"`
	Name        string  `gorm:"type:varchar(30);not null;comment:商品名称"`
	Price       float64 `gorm:"type:decimal(10,2);not null;comment:售价"`
	Status      int64   `gorm:"type:tinyint;not null;default:1;index;comment:1-上架 0-下架"`
	Description string  `gorm:"type:text;comment:商品描述"`
	StockId     int64   `gorm:"type:int;not null;default:0;comment:库存"`
	Types       int64   `gorm:"type:int(11);not null;default:0;comment:商品类型"`
	IsHot       int64   `gorm:"type:tinyint;default:0;comment:1=预热 0=普通"`
}
type Image struct {
	gorm.Model
	Url     string `gorm:"type:varchar(200);not null;comment:图片URL"`
	GoodsId int64  `gorm:"type:int(11);not null;comment:商品ID"`
}
type Stocks struct {
	gorm.Model
	GoodsId      int64   `gorm:"type:int(11);not null;comment:商品ID"`
	Stock        int64   `gorm:"type:int;not null;default:0;comment:库存"`
	Supplier     string  `gorm:"type:varchar(30);not null;comment:供货商名称"`                //供货商
	StockNum     int64   `gorm:"type:int;not null;default:0;comment:库存数量"`                //库存数量
	StockAddress string  `gorm:"type:varchar(200);not null;comment:库存物品位置"`             //库存物品位置
	StockUnit    string  `gorm:"type:varchar(30);not null;comment:商品单位"`                  //商品单位
	StockPrice   float64 `gorm:"type:decimal(10,2);not null;comment:售价"`                    //价格
	StockStatus  int64   `gorm:"type:tinyint;not null;default:1;index;comment:1-上架 0-下架"` //库存状态"`
}

func (s *Stocks) FindStockById(db *gorm.DB, id int64, id2 int64) error {
	return db.Where("goods_id = ? AND id = ?", id2, id).First(s).Error
}

func (s *Stocks) UpStockMessage(db *gorm.DB, in *users.UpStockMessageRequest) error {
	return db.Debug().Save(s).Error
}

func (g *Goods) AddGood(db *gorm.DB) error {
	return db.Create(g).Error
}

func (g *Goods) IsHots(db *gorm.DB, in *users.IsHotRequest) error {
	return db.Model(&Goods{}).Where("id = ? AND user_id = ?", in.GoodsId,
		in.UserId).Update("is_hot", 1).Error
}

func (g *Goods) GoodsList(db *gorm.DB, list []*users.GoodsList, in *users.GoodsListRequest) []*users.GoodsList {
	/*
		SELECT
		  `goods`.`user_id`,
		  `goods`.`name`,
		  `goods`.`price`,
		  `goods`.`status`,
		  `goods`.`description`,
		  `goods`.`stock_id`,
		  `goods`.`types`,
		  `goods`.`is_hot`,
		  `stocks`.`stock`,
		  `stocks`.`supplier`,
		  `stocks`.`stock_num`,
		  `stocks`.`stock_address`,
		  `stocks`.`stock_unit`,
		  `stocks`.`stock_price`,
		  `stocks`.`stock_status`
		FROM
		  `goods`
		  JOIN stocks ON stocks.id = goods.stock_id
		WHERE
		  `goods`.`deleted_at` IS NULL
		  LIMIT 2
	*/

	db.Model(&Goods{}).Select(
		"`goods`.`user_id`",
		"`goods`.`name`",
		"`goods`.`price`",
		"`goods`.`status`",
		"`goods`.`description`",
		"`goods`.`stock_id`",
		"`goods`.`types`",
		"`goods`.`is_hot`",
		"`stocks`.`stock`",
		"`stocks`.`supplier`",
		"`stocks`.`stock_num`",
		"`stocks`.`stock_address`",
		"`stocks`.`stock_unit`",
		"`stocks`.`stock_price`",
		"`stocks`.`stock_status`").
		Joins("JOIN stocks ON stocks.id = goods.stock_id").
		Scopes(pkg.Paginate(int(in.Page), int(in.Size))).
		Find(&list)
	return list
}

func (g *Goods) GoodsDetailed(db *gorm.DB, list *users.GoodsList, in *users.GoodsDetailedRequest) *users.GoodsList {
	db.Model(&Goods{}).Select(
		"`goods`.`user_id`",
		"`goods`.`name`",
		"`goods`.`price`",
		"`goods`.`status`",
		"`goods`.`description`",
		"`goods`.`stock_id`",
		"`goods`.`types`",
		"`goods`.`is_hot`",
		"`stocks`.`stock`",
		"`stocks`.`supplier`",
		"`stocks`.`stock_num`",
		"`stocks`.`stock_address`",
		"`stocks`.`stock_unit`",
		"`stocks`.`stock_price`",
		"`stocks`.`stock_status`").
		Joins("JOIN stocks ON stocks.id = goods.stock_id").
		Where("goods.id = ?", in.GoodId).
		Find(&list)
	return list
}
