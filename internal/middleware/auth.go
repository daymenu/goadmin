package middleware

import (
	"github.com/daymenu/goadmin/internal/global"
	"github.com/gin-gonic/gin"
)

// AdminAuth admin auth
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		global.AppLog.Info("AdminAuth")
		c.AbortWithStatus(403)
		c.Next()
	}
}
