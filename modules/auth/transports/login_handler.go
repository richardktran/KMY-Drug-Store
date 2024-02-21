package transports

import (
	"github.com/gin-gonic/gin"
	"github.com/richardktran/MyBlogBE/database"
)

func Login(c *gin.Context) {

	// Test db connection

	db := database.GetDB()
	if db == nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
