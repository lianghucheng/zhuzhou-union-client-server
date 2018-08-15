package models

import "github.com/jinzhu/gorm"

type Home struct {
	gorm.Model
	Name           string //首页分类名称
	Category       *Category
	CategoryID     uint
	Position       int //每个分类的位置
	IndexArticleID uint
	Url            string
	Layout         int //模块
	SubCategorys   []*Category
}
