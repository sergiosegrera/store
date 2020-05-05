package config

import "os"

type Config struct {
	HttpPort     string
	GrpcPort     string
	JwtSecretKey []byte
	Password     []byte
}

func New() *Config {
	return &Config{
		HttpPort:     ParseEnv("AUTH_HTTP_PORT", "8085"),
		GrpcPort:     ParseEnv("AUTH_GRPC_PORT", "8001"),
		JwtSecretKey: []byte(ParseEnv("JWT_KEY", "verysecretmuchwow")),
		Password:     []byte(ParseEnv("ADMIN_PASSWORD", "verysecretmuchwow")),
	}
}

func ParseEnv(key string, def string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return def
}
