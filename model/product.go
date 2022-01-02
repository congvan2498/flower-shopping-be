package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Id int
	Code string `gorm:"index"`
	Name string
	Type string
	ImageUrl string
	Description string
}
