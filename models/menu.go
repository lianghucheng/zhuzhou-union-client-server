package models

import "github.com/jinzhu/gorm"

type Menu struct {
	gorm.Model
	Name     string
	Higher   *Category
	HigherID uint
	Category *Category
}
