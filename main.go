package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	fmt.Println("Application is running on ", "localhost:8080")

	router.Run(":8080")
}
