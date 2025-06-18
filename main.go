package main

import (
	"github.com/Abdulqudri/myapi/configs"
	"github.com/Abdulqudri/myapi/database"
	"github.com/Abdulqudri/myapi/routers"
	"github.com/gin-gonic/gin"
)

func main() {

	configs.LoadEnv()
	database.ConnectDatabase()
	r := gin.Default() 

	routers.RegisterRoutes(r)

	// Start server
	r.Run(":8080") // listen on localhost:8080
}
