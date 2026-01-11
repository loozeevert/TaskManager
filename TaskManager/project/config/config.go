package config

import (
	"os"
	"strconv"
)

type Config struct {
	HTTPAddr     string
	ServerPort   string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	DBSSLMode    string
	Environment  string
	ReadTimeout  int
	WriteTimeout int
}

func NewConfig() *Config {
	return &Config{
		HTTPAddr:     getEnv("HTTP_ADDR", "0.0.0.0:8888"),
		ServerPort:   getEnv("SERVER_PORT", "8888"),
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       getEnv("DB_PORT", "5432"),
		DBUser:       getEnv("DB_USER", "postgres"),
		DBPassword:   getEnv("DB_PASSWORD", ""),
		DBName:       getEnv("DB_NAME", "taskmanager"),
		DBSSLMode:    getEnv("DB_SSL_MODE", "disable"),
		Environment:  getEnv("ENVIRONMENT", "development"),
		ReadTimeout:  getEnvAsInt("READ_TIMEOUT", 10),
		WriteTimeout: getEnvAsInt("WRITE_TIMEOUT", 10),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}