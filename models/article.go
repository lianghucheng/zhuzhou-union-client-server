package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/media/oss"
)

type Article struct {
	gorm.Model
	Status            uint
	Title             string
	Source            string
	Author            string
	Category          Category
	CategoryID        uint
	Cover             oss.OSS
	Content           string `gorm:"longtext"`
	Editor            string
	ResponsibleEditor string
}
