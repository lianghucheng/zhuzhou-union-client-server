package models

import "github.com/jinzhu/gorm"

type Menu struct {
	gorm.Model
	Name       string
	Higher     *Menu
	HigherID   uint
	Category   Category
	CategoryID uint
	URL        string
	Menus      []Menu
}
