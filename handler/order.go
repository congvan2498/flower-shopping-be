package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"togo/controller"
	"togo/form"
	"togo/middleware"
	"togo/model"
)

func (a *APIEnv) CreateOrder(c *gin.Context) {
	var input form.CreateOrderRequest
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserFromCtx(c)

	createOrderRequest := model.Order{
		Code:   input.Code,
		UserId: user.Id,
		Status: input.Status,
		Address: model.Address{
			Address:      input.Address.Address,
			WardName:     input.Address.WardName,
			DistrictName: input.Address.DistrictName,
			ProvinceName: input.Address.ProvinceName,
			Code:         input.Address.Code,
			Phone:        input.Address.Phone,
			Name:         input.Address.Name,
		},
		Product: []model.Product{},
	}

	if len(input.Product) >= 0 {
		for _, v := range input.Product {
			createOrderRequest.Product = append(createOrderRequest.Product, model.Product{
				Code:        v.Code,
				Name:        v.Name,
				Type:        v.Type,
				ImageUrl:    v.ImageUrl,
				Description: v.Description,
			})
		}
	}

	err := controller.CreateOrder(&createOrderRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "create order successfully"})
}

func (a *APIEnv) GetOrder(c *gin.Context) {
	queryParam := c.Request.URL.Query()

	if queryParam == nil || queryParam["order_id"] == nil {
		products, err := controller.GetOrders()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		total := controller.CountOrder()
		c.JSON(http.StatusOK, gin.H{"total": total, "data": products})
		return
	}

	orderId, err := strconv.Atoi(queryParam["order_id"][0])

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := controller.GetOrder(uint(orderId))
	c.JSON(http.StatusOK, gin.H{"data": order})
}
