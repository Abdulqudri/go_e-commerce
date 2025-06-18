package main

import (
	"log"
	"os"

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
	log.Println("Running in env:", os.Getenv("GIN_ENV"))

	port := configs.GetPort()
	
	r.Run(":" + port)
}
