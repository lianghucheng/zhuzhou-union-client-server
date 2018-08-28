package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Prioty   int	//1,root管理员   2,普通管理员	2,用户
	Name     string
	Sex      int //1男   2女
	Icon     string
	QQ       string
	Email    string
	Sign     string
}

func (userInfo User) DisplayName() string {
	return userInfo.Username
}
