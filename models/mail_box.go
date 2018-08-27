package models

import "github.com/jinzhu/gorm"

type MailBox struct {
	gorm.Model
	Title   string
	Content string `gorm:"type:longtext"`
	Ip      string
}
