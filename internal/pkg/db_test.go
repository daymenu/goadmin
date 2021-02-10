package pkg

import (
	"testing"

	"github.com/daymenu/goadmin/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

func TestGetGoAdminDB(t *testing.T) {
	dbconf := &config.DB{
		Driver:   "mysql",
		Host:     "127.0.0.1",
		Port:     "3306",
		UserName: "root",
		Password: "123456",
		DbName:   "mysql",
		Charset:  "utf8mb4",
	}
	client, err := EntClient(&AdminConfigDB{dbconf})

	if err != nil {
		t.Error(err)
	}

	if err := client; err != nil {
		t.Error(err)
	}
}
