package router

import (
	"github.com/daymenu/goadmin/internal/app/admin/controller"
	"github.com/daymenu/goadmin/internal/middleware"
	"github.com/gin-gonic/gin"
)

// AddAdmin add admin router
func AddAdmin(e *gin.Engine) {
	admin := e.Group("admin", middleware.AdminAuth())
	admin.GET("index", controller.AdminIndex)
}
