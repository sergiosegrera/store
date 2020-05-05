package config

import "os"

type Config struct {
	DatabaseAddress string
	HttpPort        string
	GrpcPort        string
}

func New() *Config {
	return &Config{
		DatabaseAddress: ParseEnv("DB_ADDRESS", "db:5432"),
		HttpPort:        ParseEnv("CART_HTTP_PORT", "8082"),
		GrpcPort:        ParseEnv("CART_GRPC_PORT", "8000"),
	}
}

func ParseEnv(key string, def string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return def
}
