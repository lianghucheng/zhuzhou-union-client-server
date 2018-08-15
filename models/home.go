package models

import "github.com/jinzhu/gorm"

type Home struct {
	gorm.Model
	Name           string //首页分类名称
	Category       *Category
	CategoryID     uint
	Position       int      //每个分类的位置
	IndexArticle   *Article //单个分类置顶的文章
	IndexArticleID uint
	Url            string
	Layout         int //模块
	SubCategorys   []*Category
}
