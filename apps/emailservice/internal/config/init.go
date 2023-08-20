package config

import (
	"fmt"

	"github.com/pkg/errors"
	goconfig "go-micro.dev/v4/config"
	"go-micro.dev/v4/config/source/file"
)

var cfg *Config = new(Config)

func Address() string {
	return fmt.Sprintf(":%d", cfg.Port)
}

func Email() *EmailConfig {
	return &cfg.Email
}

func Cache() *CacheConfig {
	return &cfg.Cache
}

func Tracing() *TracingConfig {
	return &cfg.Tracing
}

func Load(configPath string) error {
	source := file.NewSource(
		file.WithPath(configPath),
	)

	configor, err := goconfig.NewConfig(goconfig.WithSource(source))
	if err != nil {
		return errors.Wrap(err, "configor.New")
	}

	if err := configor.Load(); err != nil {
		return errors.Wrap(err, "configor.Load")
	}

	if err := configor.Scan(cfg); err != nil {
		return errors.Wrap(err, "configor.Scan")
	}

	return nil
}
