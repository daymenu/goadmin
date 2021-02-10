package global

import (
	"context"
	"testing"
)

func TestAppLog(t *testing.T) {
	log, err := initAppLog()
	if err != nil {
		t.Error(err)
	}
	log.Info("hahaha")
	defer log.Sync()
}

func TestDB(t *testing.T) {
	client, err := initAdminEntClient()

	if err != nil {
		t.Error(err)
	}

	admin, err := client.Admin.Query().First(context.Background())
	if err != nil {
		t.Error(err)
	}

	t.Error(admin)
}
