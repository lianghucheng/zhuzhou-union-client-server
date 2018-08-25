package models

import "github.com/jinzhu/gorm"

type StaffShow struct {
	gorm.Model
	Title    string
	User     *User
	UserID   uint
	Status   int //是否审核
	Category int //1书法作品 2摄影作品 3视频作品
	Content  string `gorm:"type:longtext"`
}
