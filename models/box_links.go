package models

import "github.com/jinzhu/gorm"

type BoxLinks struct {
	gorm.Model
	Name     string
	Url      string
	Position int //3个下拉框的位置
}


