package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Environment string

const (
	EnvTest       Environment = "test"
	EnvLocal      Environment = "local"
	EnvProduction Environment = "production"
)

type Config struct {
	Environment Environment `required:"true" envconfig:"ENVIRONMENT"`
	Development bool        `required:"true" envconfig:"DEVELOPMENT"`

	Postgres Postgres

	JwtSecretKey string `required:"true" envconfig:"JWT_SECRET"`
}

type Postgres struct {
	Host         string `required:"true" envconfig:"DB_HOST" 			default:"localhost"`
	User         string `required:"true" envconfig:"DB_USER"			default:"postgres"`
	Password     string `required:"true" envconfig:"DB_PASSWORD"		default:"postgres"`
	DatabaseName string `required:"true" envconfig:"DB_NAME" 			default:"digitalbank"`
	Port         string `required:"true" envconfig:"DB_PORT"			default:"5432"`
}

func Load() (Config, error) {
	const operation = "Config.New"

	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("%s -> %w", operation, err)
	}
	return cfg, nil
}
