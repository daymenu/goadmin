package pkg

import (
	"fmt"

	// mysql db driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/daymenu/goadmin/internal/config"
	"github.com/daymenu/goadmin/internal/ent"
)

// EntClient ent db
func EntClient(c *AdminConfigDB) (*ent.Client, error) {
	dns := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		c.UserName,
		c.Password,
		c.Host,
		c.Port,
		c.DbName,
	)
	return ent.Open(c.Driver, dns)
}

// AdminConfigDB admin db config
type AdminConfigDB struct {
	*config.DB
}

// NewAdminConfigDB new a config db
func NewAdminConfigDB() *AdminConfigDB {
	dbConf, err := config.New().Load()
	if err != nil {
		panic("read config failed!")
	}
	return &AdminConfigDB{&dbConf.DBs.GoAdmin}
}
