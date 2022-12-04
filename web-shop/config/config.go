package config

import "os"

type DBConfig struct {
	DB_HOST   string
	DB_PORT   string
	DB_NAME   string
	DB_USER   string
	DB_PASSWD string
}

type Config struct {
	Host    string
	Port    string
	DB      DBConfig
	GinMode string
}

func NewConfig() *Config {
	return &Config{
		Host: getEnv("HOST", "localhost"),
		Port: getEnv("PORT", "80"),
		DB: DBConfig{
			DB_HOST:   getEnv("DB_HOST", "localhost"),
			DB_PORT:   getEnv("DB_PORT", "5432"),
			DB_NAME:   getEnv("DB_NAME", ""),
			DB_USER:   getEnv("DB_USER", "postgres"),
			DB_PASSWD: getEnv("DB_PASSWD", ""),
		},
		GinMode: getEnv("GIN_MODE", "debug"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultVal
}
