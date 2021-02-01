package config

import (
	"testing"
)

func TestLoad(t *testing.T) {
	config := New()
	c, err := config.Load()
	if err != nil {
		t.Error(err)
	}
	t.Errorf("%#v", c)
}

func TestLog(t *testing.T) {
	config, err := New().Load()
	if err != nil {
		t.Error(err)
	}
	t.Errorf("%#v", config.GetLog())
}
