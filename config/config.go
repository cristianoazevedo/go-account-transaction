package config

import "os"

//Config definition of the configuration struct
type Config struct {
	App      *AppConfig
	DBConfig *DBConfig
}

//DBConfig struct that contains the database configuration
type DBConfig struct {
	Driver string
	Name   string
	Host   string
	User   string
	Pwd    string
}

//AppConfig struct that contains the configuration of the application
type AppConfig struct {
	Name string
	Host string
}

//GetConfig create a new configuration
func GetConfig() *Config {
	return &Config{
		App: &AppConfig{
			Name: os.Getenv("APP_NAME"),
			Host: os.Getenv("APP_HOST"),
		},
		DBConfig: &DBConfig{
			Driver: os.Getenv("DB_DRIVER_NAME"),
			Name:   os.Getenv("DB_NAME"),
			Host:   os.Getenv("DB_HOST"),
			User:   os.Getenv("DB_USER"),
			Pwd:    os.Getenv("DB_PASS"),
		},
	}
}
