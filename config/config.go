package config

import "os"

type Config struct {
	App      *AppConfig
	DBConfig *DBConfig
}

type DBConfig struct {
	Name string
	Host string
	User string
	Pwd  string
}

type AppConfig struct {
	Name string
	Host string
}

func GetConfig() *Config {
	return &Config{
		App: &AppConfig{
			Name: os.Getenv("APP_NAME"),
			Host: os.Getenv("APP_HOST"),
		},
		DBConfig: &DBConfig{
			Name: os.Getenv("DB_NAME"),
			Host: os.Getenv("DB_HOST"),
			User: os.Getenv("DB_USER"),
			Pwd:  os.Getenv("DB_PASS"),
		},
	}
}
