package controllers

import (
	"net/http"
	"strconv"

	"github.com/Abdulqudri/myapi/dtos"
	"github.com/Abdulqudri/myapi/services"
	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var input dtos.CreateCategoryDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := services.CreateCategory(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Category Created Successfully", "category": category})
}

func GetCategories(c *gin.Context) {
	categories, err := services.GetCategories()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "sucessful", "categories": categories})

}

func GetCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	category, err := services.GetCategory(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"message": "successful", "category": category})
}
