package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {

	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No .env file found or error loading it.")
	} else {
		log.Println("✅ .env loaded.")
	}

}

func GetDSN() string {
	return os.Getenv("DB_USER") + ":" +
		os.Getenv("DB_PASS") + "@tcp(" +
		os.Getenv("DB_HOST") + ":" +
		os.Getenv("DB_PORT") + ")/" +
		os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
}
func GetTokenSecret() string {

	secret := os.Getenv("TOKEN_SECRET")
	if secret == "" {
		log.Fatal("TOKEN_SECRET is not set in the environment variables.")
	}
	return secret
}
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("PORT is not set, defaulting to 8080.")
		return "8080"
	}
	return port
}

func GetEnv() string {
	ginEnv := os.Getenv("GIN_ENV")
	if ginEnv == "" {
		log.Println("Env is not set, defaulting to development.")
		return "development"
	}
	return ginEnv
}
