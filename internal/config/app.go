package config

// DB db config
type DB struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	Charset  string `yaml:"charset"`
}

// DBs database
type DBs struct {
	GoAdmin DB `yaml:"goadmin"`
}

// Log log
type Log struct {
	Filename   string `json:"filename" yaml:"filename"`
	MaxSize    int    `json:"maxsize" yaml:"maxsize"`
	MaxAge     int    `json:"maxage" yaml:"maxage"`
	MaxBackups int    `json:"maxbackups" yaml:"maxbackups"`
	LocalTime  bool   `json:"localtime" yaml:"localtime"`
	Compress   bool   `json:"compress" yaml:"compress"`
}

// Conf define config
type Conf struct {
	DBs DBs `yaml:"db"`
	Log Log `yaml:"log"`
}

// GoAdminDB 获取goadmin 数据库
func (c *Conf) GoAdminDB() *DB {
	return &c.DBs.GoAdmin
}

// GetLog get log config
func (c *Conf) GetLog() *Log {
	return &c.Log
}
