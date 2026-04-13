package configs

import (
	"github.com/go-redis/redis/v8"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

var (
	Conf *Configs
	DB   *gorm.DB
	Rdb  *redis.Client
	Esc  *elastic.Client
)
