package controllers

import (
	"net/http"

	"github.com/Abdulqudri/myapi/database"
	"github.com/Abdulqudri/myapi/models"
	"github.com/Abdulqudri/myapi/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var input struct {
		Name    string `json:"name" binding:"required"`
		Email   string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{Name: input.Name, Email: input.Email, Password: string(hash)}
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user": user})
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	accessToken, _ := utils.GenerateAccessToken(user.ID)
	refreshToken, _ := utils.GenerateRefreshToken(user.ID)
	
	user.RefreshToken = refreshToken
	database.DB.Save(&user)

	utils.SetSecureCookie(c.Writer, "refresh_token", refreshToken, 7*24*60*60)


	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": accessToken})
}

func Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token not found"})
		return
	}
	claims, err := utils.ValidateToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}
	userID := uint(claims["user_id"].(float64))	

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil || user.RefreshToken != refreshToken {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
		return
	}

	accessToken, err := utils.GenerateAccessToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}

func Logout(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token not found"})
		return
	}

	claims, err := utils.ValidateToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	userID := uint(claims["user_id"].(float64))
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil || user.RefreshToken != refreshToken {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
		return
	}

	user.RefreshToken = ""
	database.DB.Save(&user)
	 
	utils.DeleteCookie(c.Writer, "refresh_token")
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
func GetMe(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}