package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Prioty   int
	Article []Article `gorm:"foreignkey:UserID"`
}

func (userInfo User) DisplayName() string {
	return userInfo.Username
}
