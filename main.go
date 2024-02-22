package main

import (
	"log"

	"github.com/richardktran/MyBlogBE/pkg/database"
	"github.com/richardktran/MyBlogBE/pkg/env"
	"github.com/richardktran/MyBlogBE/pkg/router"
)

func init() {
	env.Setup()
}

func main() {
	router := router.GetRouter()

	if router == nil {
		log.Fatal("Failed to initialize router")
	}

	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Run("localhost:3000")

	defer database.CloseDB()
}
