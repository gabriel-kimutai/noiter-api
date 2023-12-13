package models

import "gorm.io/gorm"

type Dish struct {
	gorm.Model
	Name string
	Description string
	Price float32
	Rating float32
	OrderID uint
	CategoryID uint
}