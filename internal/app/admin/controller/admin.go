package controller

import "github.com/gin-gonic/gin"

// AdminIndex admin index
func AdminIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "admin",
	})
}
