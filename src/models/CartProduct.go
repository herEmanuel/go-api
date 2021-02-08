package models

import "gorm.io/gorm"

type CartProduct struct {
	gorm.Model
	User uint
	ProductID uint
	Product Product `gorm:"constraint:OnDelete:CASCADE"`
	Amount int
}