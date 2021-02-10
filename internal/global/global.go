package global

import (
	"github.com/daymenu/goadmin/internal/ent"
	"github.com/daymenu/goadmin/internal/pkg"
)

var (
	// AdminEntClient ent client
	AdminEntClient *ent.Client

	// AppLog app log
	AppLog *pkg.Logger
)

func init() {
	AdminEntClient, _ = initAdminEntClient()
	AppLog, _ = initAppLog()
}
