package config

import (
	"os"
	"strconv"
)

type Config struct {
	ServerPort int
	RedisAddr  string
}

func getEnv(key string, defaultValue any) any {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	switch defaultValue.(type) {
	case int:
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	case bool:
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	case float64:
		if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
			return floatValue
		}
	case string:
		return value
	}

	return defaultValue
}

func NewConfig() *Config {
	return &Config{
		ServerPort: getEnv("SERVER_PORT", 2609).(int),
		RedisAddr:  getEnv("REDIS_ADDR", "localhost:6379").(string),
	}
}
