package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/media/oss"
)

type ImageLinks struct {
	gorm.Model
	Url   string
	Image oss.OSS
}
