package router

import (
	"github.com/daymenu/goadmin/internal/app/admin/controller"
	"github.com/gin-gonic/gin"
)

// AddAdmin add admin router
func AddAdmin(e *gin.Engine) {
	e.GET("admin/index", controller.AdminIndex)
}
