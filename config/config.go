package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Host string
	Port int32

	PostgresDatabase string
	PostgresUsername string
	PostgresPassword string
	PostgresHost     string
	PostgresPort     int32

	EmailPrivateKey string
	EmailPublicKey  string
}

func NewConfig() (cfg *Config, err error) {
	cfg = &Config{}

	if err = godotenv.Load("./.env"); err != nil {
		return cfg, err
	}

	cfg.Host = cast.ToString(getDefaultOrValue(os.Getenv("Host"), "127.0.0.1"))
	cfg.Port = cast.ToInt32(getDefaultOrValue(os.Getenv("Port"), 3000))

	cfg.PostgresDatabase = cast.ToString(getDefaultOrValue(os.Getenv("PostgresDatabase"), "postgres"))
	cfg.PostgresUsername = cast.ToString(getDefaultOrValue(os.Getenv("PostgresUsername"), "postgres"))
	cfg.PostgresPassword = cast.ToString(getDefaultOrValue(os.Getenv("PostgresPassword"), "postgres"))
	cfg.PostgresHost = cast.ToString(getDefaultOrValue(os.Getenv("PostgresHost"), "0.0.0.0"))
	cfg.PostgresPort = cast.ToInt32(getDefaultOrValue(os.Getenv("PostgresPort"), 5432))

	cfg.EmailPrivateKey = cast.ToString(getDefaultOrValue(os.Getenv("EmailPrivateKey"), "super_secret_key"))
	cfg.EmailPublicKey = cast.ToString(getDefaultOrValue(os.Getenv("EmailPublicKey"), "super_public_key"))

	return cfg, err
}

func getDefaultOrValue(val, def interface{}) interface{} {
	if val == "" {
		return def
	}
	return val
}
