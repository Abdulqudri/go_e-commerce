package routers

import (
	"github.com/Abdulqudri/myapi/controllers"
	"github.com/Abdulqudri/myapi/middlewares"
	"github.com/gin-gonic/gin"
)



func RegisterPoductRoutes(r *gin.Engine) {
	products := r.Group("/products")
	{
		products.POST("/", middlewares.RequiredAuth, controllers.CreateProduct)
		products.PUT("/:id", middlewares.RequiredAuth, controllers.UpdateProduct)
		products.GET("/", middlewares.RequiredAuth, controllers.GetProducts)
		products.GET("/:id", middlewares.RequiredAuth, controllers.GetProduct)

	}
}