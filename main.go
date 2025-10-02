package main

import (
	"log"

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
	routers.RegisterPoductRoutes(r)

	// Start server
	log.Println("Running in env:", configs.GetEnv())

	port := configs.GetPort()
	
	r.Run(":" + port)
}
