package config

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"

	"gopkg.in/yaml.v2"
)

// Config config struct
type Config struct {
	profile string
	dir     string
	conf    *Conf
}

// New 创建一个 config 的对象
func New() *Config {
	return &Config{
		dir: defaultDir(),
	}
}

// SetProfile set different profile
func (c *Config) SetProfile(profile string) *Config {
	c.profile = profile
	return c
}

func defaultDir() string {
	_, f, _, _ := runtime.Caller(0)
	path := filepath.Dir(filepath.Dir(filepath.Dir(f)))
	return filepath.Join(path, "config")
}

// Path path
func (c *Config) Path() string {
	pre := filepath.Join(c.dir, "app")
	if strings.TrimSpace(c.profile) == "" {
		return pre + ".yml"
	}
	return pre + "-" + c.profile + ".yml"
}

// Load 加载配置文件
func (c *Config) Load() (*Conf, error) {
	fb, err := ioutil.ReadFile(c.Path())
	if err != nil {
		return nil, err
	}
	conf := &Conf{}
	parseConfig(fb, conf)
	return conf, nil
}

func parseConfig(data []byte, c *Conf) (*Conf, error) {
	err := yaml.Unmarshal(data, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
