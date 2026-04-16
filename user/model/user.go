package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(30);uniqueIndex;not null;comment:用户名"`
	Account  string `gorm:"type:varchar(30);not null;comment:账号"`
	Password string `gorm:"type:varchar(32);not null;comment:密码"`
	Mobile   string `gorm:"type:varchar(11);not null;comment:手机号"`
}

func (u *User) FindUserByAccount(db *gorm.DB, account string) error {
	return db.Where("account = ?", account).First(u).Error
}

func (u *User) FindUserByMobile(db *gorm.DB, mobile string) error {
	return db.Where("mobile = ?", mobile).First(u).Error
}
