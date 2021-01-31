package pkg

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/daymenu/goadmin/internal/config"
)

func TestGetGoAdminDB(t *testing.T) {
	client, err := EntAdminClient(&config.DB{
		Driver:   "mysql",
		Host:     "127.0.0.1",
		Port:     "3306",
		UserName: "root",
		Password: "123456",
		DbName:   "mysql",
		Charset:  "utf8mb4",
	})

	if err != nil {
		t.Error(err)
	}

	if err := client; err != nil {
		t.Error(err)
	}
}
