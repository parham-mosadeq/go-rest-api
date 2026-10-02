package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")
	server := gin.Default()
	server.Run(":8080")
}
