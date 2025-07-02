package services

import (
	"errors"

	"github.com/Abdulqudri/myapi/database"
	"github.com/Abdulqudri/myapi/dtos"
	"github.com/Abdulqudri/myapi/models"
	"github.com/Abdulqudri/myapi/utils"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(input dtos.CreateUserDTO) (models.User, error) {

	hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hash),
	}

	err := database.DB.Create(&user).Error

	return user, err
}

func LoginUser(input dtos.LoginDTO) (string, string, error) {
	var user models.User
	// Step 1: Find user by email
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return "", "", errors.New("invalid email or password")
	}

	// Step 2: Compare passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return "", "", errors.New("invalid email or password")
	}

	accessToken, err := utils.GenerateAccessToken(user.ID)
	if err != nil {
		return "", "", errors.New("failed to generate access token")
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		return "", "", errors.New("failed to generate refresh token")
	}

	// Step 4: Store refresh token in DB
	user.RefreshToken = refreshToken
	database.DB.Save(&user)

	return refreshToken, accessToken, nil

}

func Refresh(refreshToken string) (string, error) {
	claims, err := utils.ValidateToken(refreshToken)
	if err != nil {
		return "", errors.New(" Invalid refresh token")
	}
	userID := uint(claims["user_id"].(float64))

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil || user.RefreshToken != refreshToken {
		return "", errors.New(" Invalid session")
	}

	accessToken, err := utils.GenerateAccessToken(user.ID)
	if err != nil {
		return "", errors.New(" Failed to generate access token")
	}

	return accessToken, nil
}

func Logout(refreshToken string) error{
	claims, err := utils.ValidateToken(refreshToken)
	if err != nil {
		return errors.New(" Invalid refresh token")
	}

	userID := uint(claims["user_id"].(float64))
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil || user.RefreshToken != refreshToken {
		return errors.New(" Invalid session") 
	}

	user.RefreshToken = ""
	database.DB.Save(&user)

	return nil
}