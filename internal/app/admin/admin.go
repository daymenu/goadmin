package admin

import "github.com/gin-gonic/gin"

// RunAdmin run a server
func RunAdmin(addr ...string) {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(addr...) // 监听并在 0.0.0.0:8080 上启动服务
}
