package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")
	server := gin.Default()
	server.GET("/", func(c *gin.Context) {
		// * H means Map, which is a shortcut for map[string]interface{}
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	server.Run(":8080")
}
