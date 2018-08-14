package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name     string
	Sequence uint
	Higher   *Category
	HigherID uint
	Category int
	Special  int // 用于在文章列表页显示的分类模块判断
}
