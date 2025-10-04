package routers

import (
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	RegisterRoutes(r)
	RegisterPoductRoutes(r)
	RegisterCategoryRoutes(r)
}
