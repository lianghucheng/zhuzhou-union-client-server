package models

import "github.com/jinzhu/gorm"

type MailBox struct {
	gorm.Model
	Title   string
	User    *User
	UserID  uint
	Content string `gorm:"type:longtext"`
	Ip      string
}
