package admin

import (
	"time"

	"github.com/daymenu/goadmin/internal/app/admin/controller"
	"github.com/daymenu/goadmin/internal/app/admin/router"
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

	engine.GET("/ping", controller.Ping)

	router.AddAdmin(engine)

	engine.Run(addr...)
}
