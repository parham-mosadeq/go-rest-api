package main

import (
	"fmt"
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")
	server := gin.Default()
	server.GET("/events", getAllEvents)
	server.POST("/events", createEvent)
	server.Run(":8080")
}

func getAllEvents(c *gin.Context) {
	events := models.GetAllEvents()
	c.JSON(200, events)
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	event.ID = 1
	event.UserID = 1
	c.JSON(http.StatusCreated, gin.H{"message": "event created"})
}
