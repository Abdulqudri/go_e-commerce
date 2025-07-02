package controllers

import (
	"net/http"

	"github.com/Abdulqudri/myapi/dtos"
	"github.com/Abdulqudri/myapi/services"
	"github.com/Abdulqudri/myapi/utils"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input dtos.CreateUserDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := services.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registration  successful", "user": user})

}

func Login(c *gin.Context) {
	var input dtos.LoginDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	refreshToken, accessToken, err := services.LoginUser(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	utils.SetSecureCookie(c.Writer, "refresh_token", refreshToken, 7*24*60*60)

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": accessToken})
}

func Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token not found"})
		return
	}
	accessToken, err := services.Refresh(refreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}

func Logout(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token not found"})
		return
	}

	if err := services.Logout(refreshToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

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
