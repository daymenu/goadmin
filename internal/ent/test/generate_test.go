package test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/daymenu/goadmin/internal/config"
	"github.com/daymenu/goadmin/internal/ent/migrate"
	"github.com/daymenu/goadmin/internal/pkg"
)

func TestGenerate(t *testing.T) {
	confPath := filepath.Join(config.AppDir(), "config", "app.yml")
	dbConf, err := config.Load(confPath)
	if err != nil {
		t.Error("read db config is failed:", err)
	}
	client, err := pkg.EntAdminClient(&dbConf.DBs.GoAdmin)
	if err != nil {
		t.Error("create ent client is failed:", err)
	}
	ctx := context.Background()
	err = client.Debug().Schema.Create(ctx, migrate.WithDropIndex(true))
	if err != nil {
		t.Error(err)
	}

}
