package controllers

import (
	"net/http"
	"strconv"

	"github.com/Abdulqudri/myapi/dtos"
	"github.com/Abdulqudri/myapi/services"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context){
	var input dtos.CreateProductDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	product, err := services.CreateProduct(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return 
	}

	c.JSON(http.StatusCreated, gin.H{"message": "product created successfully", "product": product})


}

func UpdateProduct(c *gin.Context) {
	var input dtos.UpdateProductDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	id, _ := strconv.Atoi(c.Param("id"))

	product, err := services.UpdateProduct(uint(id), input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error updating product"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated succcessfully", "product":product})
}

func GetProducts(c *gin.Context) {
	products, err := services.GetProducts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	c.JSON(http.StatusOK, gin.H{"message": "successful", "products": products})
}

func GetProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := services.GetProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "successful", "product": product})
}