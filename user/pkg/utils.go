package pkg

import (
	"crypto/md5"
	"fmt"
	"gorm.io/gorm"
	"regexp"
)

func IsMobile(mobile string) bool {
	matchString, err := regexp.MatchString("^[1][3456789]\\d{9}$", mobile)
	if err != nil {
		return false
	}
	return matchString
}

func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if page <= 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
