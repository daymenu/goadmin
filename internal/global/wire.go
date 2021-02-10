//+build wireinject

package global

import (
	"github.com/daymenu/goadmin/internal/ent"
	"github.com/daymenu/goadmin/internal/pkg"
	"github.com/google/wire"
)

// initAdminEntClient 初始化数据库
func initAdminEntClient() (*ent.Client, error) {
	wire.Build(pkg.EntClient, pkg.NewAdminConfigDB)
	return &ent.Client{}, nil
}

// initAppLog 初始化日志
func initAppLog() (*pkg.Logger, error) {
	wire.Build(pkg.NewLog, pkg.DefalultLogOption)
	return &pkg.Logger{}, nil
}
