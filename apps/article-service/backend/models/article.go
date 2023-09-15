package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Name  string
	Price uint
}

func (a *Article) AfterCreate(tx *gorm.DB) (err error) {
	return nil
}

func (a *Article) AfterUpdate(tx *gorm.DB) (err error) {
	return nil
}
