package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/richardktran/MyBlogBE/database"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	if IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// Binding routes in the routes folder to the router instance
	// using the BindRoutes function from the app package
	LoadAPIRouter(router, "v1")

	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Run("localhost:3000")

	defer database.CloseDB()
}
