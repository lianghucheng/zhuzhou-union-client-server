package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Prioty   int
}

func (userInfo User) DisplayName() string {
	return userInfo.Username
}
