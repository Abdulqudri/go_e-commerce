package services

import (
	"github.com/Abdulqudri/myapi/database"
	"github.com/Abdulqudri/myapi/dtos"
	"github.com/Abdulqudri/myapi/models"
)

func CreateProduct(input dtos.CreateProductDTO) (models.Product, error) {
	product := models.Product{
		ID: input.ID,
		Name: input.Name,
		Price: input.Price,
		Category: input.Category,
		Description: input.Description,
	}

	err := database.DB.Create(&product).Error

	return product, err

}

func UpdateProduct(id uint, input dtos.UpdateProductDTO) (*models.Product, error) {
	err := database.DB.Model(&models.Product{}).Where("id = ?", id).Updates(input).Error
	if err != nil {
		return nil, err
	}

	var updatedProduct models.Product

	err = database.DB.First(&updatedProduct, id).Error


	return &updatedProduct, err
}

func GetProducts() ([]models.Product, error){
	var products []models.Product

	err := database.DB.Find(&products).Error


	return products, err
}

func GetProduct(id uint) (*models.Product, error) {
	var product models.Product
	err := database.DB.First(&product, id).Error

	return &product, err
}