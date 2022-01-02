package controller

import (
	"togo/db"
	"togo/model"
)

func CreateProduct(input *model.Product)  {
	db.DB.Create(input)
}

func GetProduct(id uint)  *model.Product{
	product := model.Product{}
	db.DB.First(&product, "id = ?", id)

	return &product
}

func GetProducts() ([]*model.Product,error)  {
	var products []*model.Product
	if err := db.DB.Find(&products).Limit(20).Error; err != nil {
		return nil,err
	}
	return products,nil
}

func CountProduct() int64 {
	var result int64
	db.DB.Table("products").Count(&result)
	return result
}
