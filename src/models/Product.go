package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name         string
	Url          string `gorm:"default:none"`
	Stock        int
	Price        float32 `gorm:"scale:2"`
	Rating       float32 `gorm:"precision:2;scale:1"`
	RatingAmount int16
	Commentaries []Commentary `gorm:"foreignKey:FromProduct;constraint:OnDelete:CASCADE"`
	Categories   []Category `gorm:"many2many:category_products"`
}
