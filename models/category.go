package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name     string
	Sequence uint
	Higher   *Category
	HigherID uint
	Category int
}
