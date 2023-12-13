package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Table Table
	Dishes []Dish
	TableID uint
	IsTaken bool
	IsDone bool
}