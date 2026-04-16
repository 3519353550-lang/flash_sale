package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/olivere/elastic/v7"
	"zgw/ks/flash_sale/user/configs"

	"zgw/ks/flash_sale/user/internal/svc"
	"zgw/ks/flash_sale/user/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchGoodsLogic {
	return &SearchGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchGoodsLogic) SearchGoods(in *users.SearchGoodsRequest) (*users.SearchGoodsResponse, error) {
	// todo: add your logic here and delete this line
	logx.Infof("SearchGoods, in: %v", in)
	boolQuery := elastic.NewBoolQuery()

	// 添加名称搜索
	if in.Name != "" {
		boolQuery.Must(elastic.NewMatchQuery("name", in.Name))
	}

	// 添加类型过滤
	if in.Types > 0 {
		boolQuery.Filter(elastic.NewTermQuery("types", in.Types))
	}

	// 添加销量过滤
	if in.Sales > 0 {
		boolQuery.Filter(elastic.NewTermQuery("sales", in.Sales))
	}

	// 添加价格范围过滤
	if in.MaxPrice > 0 || in.MinPrice > 0 {
		rangeQuery := elastic.NewRangeQuery("price")
		if in.MaxPrice > 0 {
			rangeQuery.Lte(in.MaxPrice)
		}
		if in.MinPrice > 0 {
			rangeQuery.Gte(in.MinPrice)
		}
		boolQuery.Filter(rangeQuery)
	}

	// 过滤掉状态为0的商品
	boolQuery.Filter(elastic.NewTermQuery("status", 1))

	// 添加排序
	sort := elastic.NewFieldSort("is_hot").Desc() // 1 排最前

	// 添加高亮
	highlight := elastic.NewHighlight()
	highlight.Field("name")
	highlight.Field("description")
	highlight.PreTags("<span style='color:red'>")
	highlight.PostTags("</span>")

	// 构建查询
	query := configs.Esc.
		Search().
		Index("goods").
		Query(boolQuery).
		SortBy(sort).
		Highlight(highlight).
		Size(100) // 限制返回数量

	// 执行查询
	res, err := query.Do(context.Background())
	if err != nil {
		logx.Error("Elasticsearch查询失败:", err)
		return nil, errors.New("查询失败：" + err.Error())
	}
	var list []*users.SearchGoods

	for _, hit := range res.Hits.Hits {
		var goods users.SearchGoods

		// ✅ 正确：使用 json.Unmarshal 解析 Source
		err := json.Unmarshal(hit.Source, &goods)
		if err != nil {
			logx.Error("解析商品错误:", err)
			continue
		}

		// 高亮赋值
		if name, ok := hit.Highlight["name"]; ok && len(name) > 0 {
			goods.Name = name[0]
		}

		list = append(list, &goods)
	}

	return &users.SearchGoodsResponse{
		SearchGoods: list,
	}, nil
}
