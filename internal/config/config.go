package config

import (
	"io/ioutil"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v2"
)

// AppDir 获取goadmin root path
func AppDir() string {
	return filepath.Dir(filepath.Dir(curDir()))
}

func curDir() string {
	_, f, _, _ := runtime.Caller(0)
	path := filepath.Dir(f)
	return path
}

// Load 加载配置文件
func Load(path string) (*Config, error) {
	fb, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	c := &Config{}
	parseConfig(fb, c)
	return c, nil
}

func parseConfig(data []byte, c *Config) (*Config, error) {
	err := yaml.Unmarshal(data, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// JoinPath join path
func JoinPath(path string) string {
	return filepath.Join(AppDir(), path)
}
