package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/media/oss"
)

type Article struct {
	gorm.Model
	Status            uint	//0未审核   1已审核
	User              *User
	UserID            uint
	Title             string
	Summary           string
	Source            string
	Author            string
	Category          *Category
	CategoryID        uint
	Cover             oss.OSS
	Content           string `gorm:"type:longtext"` //type:rich_editor
	Editor            string
	ResponsibleEditor string
	ReadNum           uint   //阅读数
	Url               string //媒体聚焦链接或者微信图文链接
	IsSpecial         int    //是否为文章 用作特殊页面的渲染
	VideoIndex        oss.OSS
	IsIndexUp         int
}
