package config

import (
	"os"

	"github.com/spf13/cast"
)

var PerPageSize = 100

// Config ...
type Config struct {
	Environment string // develop, staging, production

	PostgresHost     string
	PostgresPort     int
	PostgresPassword string   
	PostgresUser     string
	PostgresDB       string
	LogLevel         string

	BillingServiceHost string
	BillingServicePort int

	RPCPort string

	PasscodePool   string
	PasscodeLength int
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDB = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "sharq_user_service"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "postgres12"))

	c.BillingServiceHost = cast.ToString(getOrReturnDefault("BILLING_SERVICE_HOST", "localhost"))
	c.BillingServicePort = cast.ToInt(getOrReturnDefault("BILLING_SERVICE_PORT", 5004))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":8003"))

	c.PasscodePool = cast.ToString(getOrReturnDefault("PASSCODE_POOL", "0123456789"))
	c.PasscodeLength = cast.ToInt(getOrReturnDefault("PASSCODE_LENGTH", "6"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
