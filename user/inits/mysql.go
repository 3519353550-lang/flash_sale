package inits

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
	"zgw/ks/flash_sale/user/configs"
	"zgw/ks/flash_sale/user/model"
)

var (
	once sync.Once
	err  error
)

func InitMysql() {
	CM := configs.Conf.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		CM.User,
		CM.Password,
		CM.Host,
		CM.Port,
		CM.Database)
	once.Do(func() {
		configs.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	})
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	fmt.Println("数据库连接成功")

	if err := configs.DB.AutoMigrate(&model.User{},
		&model.Goods{},
		&model.Image{},
		&model.Stocks{},
		&model.Orders{},
		&model.OrderItem{}); err != nil {
		fmt.Println("数据表迁移失败")
		return
	}
	fmt.Println("数据表迁移成功")
	sqlDB, _ := configs.DB.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了可以重新使用连接的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

}
