package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestJoinPath(t *testing.T) {
	testFile := "internal/config/app.go"
	path := JoinPath(testFile)
	f, err := os.Open(path)
	if err != nil {
		t.Errorf("open file is failed :%s", path)
	}
	defer f.Close()
	if !strings.HasSuffix(f.Name(), testFile) {
		t.Errorf("%s is not contains %s\n", f.Name(), testFile)
	}
}

func TestLoad(t *testing.T) {
	fileName := filepath.Join("config", "app.yml")
	path := JoinPath(fileName)
	c, err := Load(path)
	if err != nil {
		t.Errorf("read config is failed")
	}
	fmt.Println(c.DBs.GoAdmin)
	t.Errorf("%#v", c)
}
