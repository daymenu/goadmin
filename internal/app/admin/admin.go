package admin

import (
	"time"

	"github.com/daymenu/goadmin/internal/middleware"
	"github.com/daymenu/goadmin/internal/pkg"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

// RunAdmin run a server
func RunAdmin(addr ...string) {
	engine := gin.New()

	logger := pkg.NewLog()
	defer logger.Sync()

	engine.Use(requestid.New())

	engine.Use(middleware.Ginzap(logger.Logger, time.RFC3339, true))

	engine.Use(middleware.RecoveryWithZap(logger.Logger, true))

	engine.GET("/ping", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "pong" + requestid.Get(c),
		})
	})
	engine.Run(addr...) // 监听并在 0.0.0.0:8080 上启动服务
}
