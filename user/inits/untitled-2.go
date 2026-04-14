package inits

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"zgw/ks/flash_sale/user/configs"
)

func EsInit() {
	e := configs.Conf.Elastic
	configs.Esc, err = elastic.NewClient(elastic.SetURL(e.Host), elastic.SetSniff(false))
	if err != nil {
		fmt.Println("Es链接失败")
		return
	}
	fmt.Println("Es链接成功")

	createGoodsIndex()
}
func createGoodsIndex() {
	ctx := context.Background()
	// 检查索引是否存在
	exists, err := configs.Esc.IndexExists("goods").Do(ctx)
	if err != nil {
		fmt.Println("检查索引失败:", err)
		return
	}

	if !exists {
		// 创建索引
		mapping := `{
			"mappings": {
				"properties": {
					"name": {
						"type": "text",
						"analyzer": "standard"
					},
					"description": {
						"type": "text",
						"analyzer": "standard"
					},
					"price": {
						"type": "float"
					},
					"types": {
						"type": "long"
					},
					"sales": {
						"type": "long"
					},
					"status": {
						"type": "long"
					},
					"is_hot": {
						"type": "long"
					}
				}
			}
		}`

		_, err = configs.Esc.CreateIndex("goods").Body(mapping).Do(ctx)
		if err != nil {
			fmt.Println("创建索引失败:", err)
			return
		}
		fmt.Println("创建goods索引成功")
	} else {
		// 删除旧索引并重新创建
		_, err = configs.Esc.DeleteIndex("goods").Do(ctx)
		if err != nil {
			fmt.Println("删除旧索引失败:", err)
			return
		}
		fmt.Println("删除旧索引成功")

		// 创建新索引
		mapping := `{
			"mappings": {
				"properties": {
					"name": {
						"type": "text",
						"analyzer": "standard"
					},
					"description": {
						"type": "text",
						"analyzer": "standard"
					},
					"price": {
						"type": "float"
					},
					"types": {
						"type": "long"
					},
					"sales": {
						"type": "long"
					},
					"status": {
						"type": "long"
					},
					"is_hot": {
						"type": "long"
					}
				}
			}
		}`

		_, err = configs.Esc.CreateIndex("goods").Body(mapping).Do(ctx)
		if err != nil {
			fmt.Println("重新创建索引失败:", err)
			return
		}
		fmt.Println("重新创建goods索引成功")
	}
}
