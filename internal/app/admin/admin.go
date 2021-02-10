package admin

import (
	"time"

	"github.com/daymenu/goadmin/internal/app/admin/controller"
	"github.com/daymenu/goadmin/internal/app/admin/router"
	"github.com/daymenu/goadmin/internal/global"

	// 初始化全局变量
	_ "github.com/daymenu/goadmin/internal/global"
	"github.com/daymenu/goadmin/internal/middleware"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

// Run run a server
func Run(addr ...string) {
	engine := gin.New()
	defer func() {
		global.AdminEntClient.Close()
		global.AppLog.Sync()
	}()
	engine.Use(requestid.New())

	engine.Use(middleware.Ginzap(global.AppLog.Logger, time.RFC3339, true))

	engine.Use(middleware.RecoveryWithZap(global.AppLog.Logger, true))

	engine.GET("/ping", controller.Ping)

	router.AddAdmin(engine)

	engine.Run(addr...)
}
