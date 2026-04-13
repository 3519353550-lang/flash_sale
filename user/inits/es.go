package inits

import (
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
}
