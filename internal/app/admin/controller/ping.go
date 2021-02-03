package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping 提供 ping pong
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "pong",
	})
}
