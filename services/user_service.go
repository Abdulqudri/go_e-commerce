package services

import (
	"github.com/Abdulqudri/myapi/database"
	"github.com/Abdulqudri/myapi/models"
)


func GetUser(id uint) (*models.User, error) {
	var user models.User
	err := database.DB.First(&user, id).Error
	return &user, err
}

func GetUsers() ([]models.User, error) {
	var users []models.User
	err := database.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func UpdateUser(id uint, user *models.User) (*models.User, error) {
	err := database.DB.Model(&models.User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"name": user.Name,
	}).Error

	if err != nil {
		return nil, err
	}
	var updatedUser models.User
	err = database.DB.First(&updatedUser, id).Error
	if err != nil {
		return nil, err
	}
	return &updatedUser, nil
}

func DeleteUser(id uint) error {
	return database.DB.Delete(&models.User{}, id).Error
}
