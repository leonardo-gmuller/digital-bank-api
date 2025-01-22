package config

import (
	"fmt"
	"time"

	"github.com/go-chi/jwtauth"
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

	// DATABASE
	Postgres Postgres

	JWT JWT
}

type App struct {
	Name                    string        `required:"true" envconfig:"APP_NAME"`
	ID                      string        `required:"true" envconfig:"APP_ID"`
	GracefulShutdownTimeout time.Duration `required:"true" envconfig:"APP_GRACEFUL_SHUTDOWN_TIMEOUT"`
}

type Server struct {
	Address      string        `required:"true" envconfig:"SERVER_ADDRESS"`
	ReadTimeout  time.Duration `required:"true" envconfig:"SERVER_READ_TIMEOUT"`
	WriteTimeout time.Duration `required:"true" envconfig:"SERVER_WRITE_TIMEOUT"`
}

type Postgres struct {
	Host         string `required:"true" envconfig:"DB_HOST"`
	User         string `required:"true" envconfig:"DB_USER"`
	Password     string `required:"true" envconfig:"DB_PASSWORD"`
	DatabaseName string `required:"true" envconfig:"DB_NAME"`
	Port         string `required:"true" envconfig:"DB_PORT"`
}

type JWT struct {
	Secret      string `required:"true" envconfig:"JWT_SECRET"`
	SecretAdmin string `required:"true" envconfig:"JWT_SECRET_ADMIN"`
	ExpiresIn   int    `required:"true" envconfig:"JWT_EXPIRES_IN"`
	TokenAuth   *jwtauth.JWTAuth
}

func New() (Config, error) {
	const operation = "Config.New"

	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("%s -> %w", operation, err)
	}

	cfg.JWT.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWT.Secret), nil)

	return cfg, nil
}
