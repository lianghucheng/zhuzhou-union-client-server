package models

import "github.com/jinzhu/gorm"

type MailBox struct {
	gorm.Model
	Author  string
	Title   string
	Contact string
	Content string `gorm:"type:longtext"`
}
