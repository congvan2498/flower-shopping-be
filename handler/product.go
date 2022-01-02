package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"togo/controller"
	"togo/form"
	"togo/model"
)

func (a *APIEnv) CreateProduct(c *gin.Context) {
	var input form.Product
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createProductRequest := &model.Product{
		Code:        input.Code,
		Name:        input.Name,
		Type:        input.Type,
		Description: input.Description,
	}

	controller.CreateProduct(createProductRequest)

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func (a *APIEnv) GetProduct(c *gin.Context) {
	log.Println(c.Request.URL.Query())
	queryParam := c.Request.URL.Query()

	if queryParam == nil || queryParam["product_id"] == nil {
		products,err := controller.GetProducts()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		total := controller.CountProduct()
		c.JSON(http.StatusOK, gin.H{"total": total,"data" : products})
		return
	}

	productId, err := strconv.Atoi(queryParam["product_id"][0])

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := controller.GetProduct(uint(productId))
	c.JSON(http.StatusOK, gin.H{"data": product})
}
