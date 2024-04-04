package config

import (
	"fmt"
	"time"

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

	App    App
	Server Server

	//DATABASE
	Postgres Postgres

	JwtSecretKey string `required:"true" envconfig:"JWT_SECRET"`
}

type App struct {
	Name                    string        `required:"true" envconfig:"APP_NAME"`
	ID                      string        `required:"true" envconfig:"APP_ID"`
	GracefulShutdownTimeout time.Duration `required:"true" envconfig:"APP_GRACEFUL_SHUTDOWN_TIMEOUT"`
}

type Server struct {
	SwaggerHost  string        `required:"true" envconfig:"SERVER_SWAGGER_HOST"`
	Address      string        `required:"true" envconfig:"SERVER_ADDRESS"`
	ReadTimeout  time.Duration `required:"true" envconfig:"SERVER_READ_TIMEOUT"`
	WriteTimeout time.Duration `required:"true" envconfig:"SERVER_WRITE_TIMEOUT"`
}

type Postgres struct {
	Host         string `required:"true" envconfig:"DB_HOST" 			default:"localhost"`
	User         string `required:"true" envconfig:"DB_USER"			default:"postgres"`
	Password     string `required:"true" envconfig:"DB_PASSWORD"		default:"postgres"`
	DatabaseName string `required:"true" envconfig:"DB_NAME" 			default:"digitalbank"`
	Port         string `required:"true" envconfig:"DB_PORT"			default:"5432"`
}

func New() (Config, error) {
	const operation = "Config.New"

	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("%s -> %w", operation, err)
	}
	return cfg, nil
}
