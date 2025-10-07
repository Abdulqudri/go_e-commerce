package services

import (
	"github.com/Abdulqudri/myapi/database"
	"github.com/Abdulqudri/myapi/dtos"
	"github.com/Abdulqudri/myapi/models"
)

func CreateCategory(input dtos.CreateCategoryDTO) (models.Category, error) {

	category := models.Category{
		ID:   input.ID,
		Name: input.Name,
	}

	err := database.DB.Create(&category).Error

	return category, err
}

func GetCategory(id uint) (*models.Category, error) {
	var category models.Category

	err := database.DB.Preload("Products").First(&category, id).Error

	return &category, err
}

func GetCategories() ([]models.Category, error) {
	var categories []models.Category

	err := database.DB.Find(&categories).Error

	return categories, err
}
