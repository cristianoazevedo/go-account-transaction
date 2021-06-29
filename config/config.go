package config

import "os"

type Config struct {
	DBConfig *DBConfig
}

type DBConfig struct {
	Name string
	Host string
	User string
	Pwd  string
}

func GetConfig() *Config {
	return &Config{
		DBConfig: &DBConfig{
			Name: os.Getenv("DB_NAME"),
			Host: os.Getenv("DB_HOST"),
			User: os.Getenv("DB_USER"),
			Pwd:  os.Getenv("DB_PASS"),
		},
	}
}
