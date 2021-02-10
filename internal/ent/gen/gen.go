package main

import (
	"context"
	"log"

	"github.com/daymenu/goadmin/internal/ent/migrate"
	"github.com/daymenu/goadmin/internal/pkg"
)

func main() {
	client, err := pkg.EntClient(pkg.NewAdminConfigDB())
	if err != nil {
		log.Fatal("create ent client is failed:", err)
	}
	defer client.Close()

	ctx := context.Background()
	err = client.Debug().Schema.Create(ctx, migrate.WithDropIndex(true))
	if err != nil {
		log.Fatal(err)
	}

}
