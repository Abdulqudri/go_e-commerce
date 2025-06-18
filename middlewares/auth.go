package middlewares

import (
	"net/http"
	"strings"

	"github.com/Abdulqudri/myapi/database"
	"github.com/Abdulqudri/myapi/models"
	"github.com/Abdulqudri/myapi/utils"
	"github.com/gin-gonic/gin"
)

func RequiredAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := utils.ValidateToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	userID := uint(claims["user_id"].(float64))
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	c.Set("user", user)	
	c.Next()

}