package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	handle()
}

func handle() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	route := gin.Default()

	route.GET("/itwokks", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hey!!!"})
	})

	route.Run()
}
