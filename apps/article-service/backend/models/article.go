package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Name  string
	Price uint
}
