package pkg

import (
	"database/sql"
	"fmt"

	"github.com/daymenu/goadmin/internal/config"
)

// Hello hello
var Hello = "hello"

// GoAdminDB goadmin db
func GoAdminDB(c *config.DB) (*sql.DB, error) {
	dns := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		c.UserName,
		c.Password,
		c.Host,
		c.Port,
		c.DbName,
	)
	return sql.Open(c.Driver, dns)
}
