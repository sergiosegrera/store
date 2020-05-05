package config

import "os"

type Config struct {
	DatabaseAddress string
	HttpPort        string
}

func New() *Config {
	return &Config{
		DatabaseAddress: ParseEnv("DB_ADDRESS", "db:5432"),
		HttpPort:        ParseEnv("PRODUCT_MANAGER_HTTP_PORT", "8082"),
	}
}

func ParseEnv(key string, def string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return def
}
