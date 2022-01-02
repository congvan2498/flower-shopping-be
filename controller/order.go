package controller

import (
	"togo/db"
	"togo/model"
)

func CreateOrder(input *model.Order) error  {
	if err := db.DB.Create(input).Error; err != nil {
		return err
	}

	return nil
}

func GetOrder(id uint)  *model.Order{
	Order := model.Order{}
	db.DB.First(&Order, "id = ?", id)

	return &Order
}

func GetOrders() ([]*model.Order,error)  {
	var Orders []*model.Order
	if err := db.DB.Find(&Orders).Limit(20).Error; err != nil {
		return nil,err
	}
	return Orders,nil
}

func CountOrder() int64 {
	var result int64
	db.DB.Table("orders").Count(&result)
	return result
}
