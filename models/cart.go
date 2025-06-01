package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Name     string    `gorm:"not null"`
	Products []Product `gorm:"many2many:cart_products;"`
}
