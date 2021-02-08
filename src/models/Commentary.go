package models

import "gorm.io/gorm"

type Commentary struct {
	gorm.Model
	Title string
	Content string
	Rating float32 `gorm:"precision:2;scale:1"`
	FromProduct uint
	Creator uint
}
