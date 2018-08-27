package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name        string
	Sequence    int
	Higher      *Category
	HigherID    uint
	Category    int //0封面,1列表,2集合,3单文章,4可投稿 按配置文件递增
	Special     int
	SubCategory []*Category `gorm:"foreignkey:HigherID"`
}
