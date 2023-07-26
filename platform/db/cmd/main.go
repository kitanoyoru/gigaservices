package main

import (
	"context"
	"os/signal"

	"github.com/kitanoyoru/gigaservices/platform/db/internal/app"
	"github.com/kitanoyoru/gigaservices/platform/db/internal/di"
	"github.com/kitanoyoru/gigaservices/platform/db/pkg/cfg"
	"github.com/samber/do"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
)

var config cfg.Config

func init() {
	config, err := cfg.NewConfig()
	if err != nil {
		log.Info("Failed to load config: %+v", err)
	}

	do.ProvideValue[*cfg.Config](di.Provider, config)

	if config.Debug {
		log.Info("Database Microservice is running on DEBUG mode")
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), unix.SIGTERM, unix.SIGINT)
	defer stop()

	errCh := make(chan error, 1)
	go func() {
		errCh <- app.Run(context.Background())
	}()

	select {
	case err := <-errCh:
		log.Fatalf("Error when runnning app: %+v", err)
	case <-ctx.Done():
		log.Info("Shutting down...")

	}
}
