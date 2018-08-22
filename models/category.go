package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name        string
	Sequence    int
	Higher      *Category
	HigherID    uint
	Category    int
	Special     int
	SubCategory []*Category `gorm:"foreignkey:HigherID"`
}
