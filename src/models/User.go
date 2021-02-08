package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	Password     string        `json:"password"`
	Verified     bool          `gorm:"default:false"`
	Country      string        `json:"country"`
	Purchases    []Purchase    `gorm:"foreignKey:buyer"`
	Role         string        `gorm:"default:user"`
	Commentaries []Commentary  `gorm:"foreignKey:creator;constraint:OnDelete:CASCADE"`
	Cart         []CartProduct `gorm:"foreignKey:user"`
}
