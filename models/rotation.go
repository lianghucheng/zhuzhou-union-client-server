package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/media/oss"
)

type Rotation struct {
	gorm.Model
	Url      oss.OSS
	Sequence int
	Link     string
}
