package models

import "gorm.io/gorm"

type Purchase struct {
	gorm.Model

	Products         []Product `gorm:"many2many:purchase_products;"`
	Buyer            uint
	TotalPrice       float32 `gorm:"scale:2"`
	ConfirmedPayment bool
	Delivered        bool `gorm:"default:false"`
}
