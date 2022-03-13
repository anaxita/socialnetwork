// Package bootstrap stores basic common entities such as config and database connection
package bootstrap

import (
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

const defaultEnvPath = "./deployment/.env"

// Config is a common structure for all config types.
type Config struct {
	HTTPPort   string `envconfig:"HTTP_PORT" validate:"required"`
	DBScheme   string `envconfig:"DB_SCHEME" validate:"required"`
	DBHost     string `envconfig:"DB_HOST" validate:"required"`
	DBPort     string `envconfig:"DB_PORT" validate:"required"`
	DBName     string `envconfig:"DB_NAME" validate:"required"`
	DBUsername string `envconfig:"DB_USERNAME" validate:"required"`
	DBPassword string `envconfig:"DB_PASSWORD" validate:"required"`
}

// NewConfig loads configuration from the environment variables, optionally loading them from the file.
func NewConfig() (*Config, error) {
	err := godotenv.Load(defaultEnvPath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	var cfg Config

	err = envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}

	err = validator.New().Struct(cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
