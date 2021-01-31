package pkg

import (
	"fmt"

	// mysql db driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/daymenu/goadmin/internal/config"
	"github.com/daymenu/goadmin/internal/ent"
)

// Hello hello
var Hello = "hello"

// EntAdminClient goadmin db
func EntAdminClient(c *config.DB) (*ent.Client, error) {
	dns := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		c.UserName,
		c.Password,
		c.Host,
		c.Port,
		c.DbName,
	)
	return ent.Open(c.Driver, dns)
}