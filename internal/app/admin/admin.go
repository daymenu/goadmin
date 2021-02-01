package admin

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RunAdmin run a server
func RunAdmin(addr ...string) {
	engine := gin.New()

	logger, _ := zap.NewDevelopment()

	engine.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	engine.Use(ginzap.RecoveryWithZap(logger, true))

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	engine.Run(addr...) // 监听并在 0.0.0.0:8080 上启动服务
}
