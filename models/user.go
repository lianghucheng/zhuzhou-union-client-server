package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Prioty   int
	Name string
	Sex int //0女   1男
	Icon string
	QQ string
	Email string
	Sign string
}

func (userInfo User) DisplayName() string {
	return userInfo.Username
}
