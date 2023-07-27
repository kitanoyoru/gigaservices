package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Hostname string `required:"false" default:"localhost:8000" envconfig:"hostname"`

	PsqlUrl  string `required:"true" default:"postgresql://kitanoyoru:kitanoyoru@postgres:5432/kitanoyoru?schema=public" envconfig:"PSQL_URL"`
	RedisUrl string `required:"true" default:"redis://localhost:6379/0" envconfig:"REDIS_URL"`
	RmqUrl   string `required:"true" default:"forgot_default_rmq_port:(" envconfig:"RMQ_URL"`

	JwtAccessSecret  string `required:"true" default:"TopSecret" envconfig:"JWT_ACCESS_SECRET"`
	JwtRefreshSecret string `required:"true" default:"TopSecret" envconfig:"JWT_REFRESH_SECRET"`
}

func NewConfig() (*Config, error) {
	var cfg Config

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(wd, "/workspace") {
		wd = "/workspace"
	} else {
		wd = filepath.Join(wd, "..", "..")
	}

	envPath := filepath.Join(wd, ".env")
	_ = godotenv.Load(envPath)

	if err = envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
