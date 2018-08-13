package models

import "github.com/jinzhu/gorm"

type Home struct {
	gorm.Model
	Category       *Category
	CategoryID     uint
	Position       int //每个分类的位置
	IndexArticle   *Article //单个分类置顶的文章
	IndexArticleID uint
}
