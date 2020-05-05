package config

import "os"

type Config struct {
	DatabaseAddress string
	CartGrpcAddress string
	StripeSecret    string
	HttpPort        string
}

func New() *Config {
	return &Config{
		DatabaseAddress: ParseEnv("DB_ADDRESS", "db:5432"),
		CartGrpcAddress: ParseEnv("CART_GRPC_ADDRESS", "cart:8000"),
		StripeSecret:    ParseEnv("STRIPE_SECRET", ""),
		HttpPort:        ParseEnv("CHECKOUT_HTTP_PORT", "8083"),
	}
}

func ParseEnv(key string, def string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return def
}
