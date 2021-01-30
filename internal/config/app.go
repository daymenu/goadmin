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

// Config define config
type Config struct {
	DBs DBs `yaml:"db"`
}
