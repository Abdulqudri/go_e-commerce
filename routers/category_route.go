package routers

import (
	"github.com/Abdulqudri/myapi/controllers"
	"github.com/Abdulqudri/myapi/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterCategoryRoutes(r *gin.Engine) {
	categories := r.Group("/categories")
	{
		categories.POST("/", middlewares.RequiredAuth, controllers.CreateCategory)
		categories.GET("/:id", middlewares.RequiredAuth, controllers.GetCategory)
		categories.GET("/", middlewares.RequiredAuth, controllers.GetCategories)
	}
}