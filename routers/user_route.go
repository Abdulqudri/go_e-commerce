package routers

import (
	"github.com/Abdulqudri/myapi/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	user := r.Group("/users")

	user.POST("/", controllers.CreateUser)
	user.GET("/", controllers.GetUsers)
	user.GET("/:id", controllers.GetUser)
	user.PUT("/:id", controllers.UpdateUser)
	user.DELETE("/:id", controllers.DeleteUser)
}