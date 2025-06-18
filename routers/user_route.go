package routers

import (
	"github.com/Abdulqudri/myapi/controllers"
	"github.com/Abdulqudri/myapi/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
		auth.GET("/me", middlewares.RequiredAuth, controllers.GetMe)
	}

	user := r.Group("/users")

	user.POST("/", middlewares.RequiredAuth, controllers.CreateUser)
	user.GET("/", middlewares.RequiredAuth, controllers.GetUsers)
	user.GET("/:id", middlewares.RequiredAuth, controllers.GetUser)
	user.PUT("/:id", middlewares.RequiredAuth, controllers.UpdateUser)
	user.DELETE("/:id", middlewares.RequiredAuth, controllers.DeleteUser)
}