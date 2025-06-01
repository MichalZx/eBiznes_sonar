package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name       string  `gorm:"not null"`
	Price      float64 `gorm:"not null"`
	CategoryId uint
	Category   Category `gorm:"foreignKey:CategoryId"`
	Carts      []Cart   `gorm:"many2many:cart_products;"`
}
