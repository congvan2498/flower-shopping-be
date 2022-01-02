package model

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Id int
	Code string
	UserId int
	TotalAmount int
	Status string
	Address Address
	Product []Product  `gorm:"many2many:order_products;"`
}

type Address struct {
	gorm.Model
	OrderId uint
	Address string
	Name string
	WardName string
	DistrictName string
	ProvinceName string
	Code string
	Phone string
}

