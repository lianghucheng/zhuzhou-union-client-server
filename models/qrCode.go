package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/media/oss"
)

type QrCode struct {
	gorm.Model
	CodeImage oss.OSS
}
