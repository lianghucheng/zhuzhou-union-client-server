package models

import "github.com/jinzhu/gorm"

type Active struct {
	gorm.Model
	Title    string
	URL      string
	Template string
}
