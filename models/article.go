package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/media/oss"
)

type Article struct {
	gorm.Model
	Status            uint	//0已审核   1未审核
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
	ReadNum           uint //阅读数
	Url               string//媒体聚焦链接或者微信图文链接
}
