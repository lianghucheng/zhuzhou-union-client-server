package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/media/oss"
)

type Article struct {
	gorm.Model
	Status            uint
	User              *User
	UserID            uint
	Title             string
	Source            string
	Author            string
	Category          *Category
	CategoryID        uint
	Cover             oss.OSS
	Content           string `gorm:"type:longtext"` //type:rich_editor
	Editor            string
	ResponsibleEditor string
}
