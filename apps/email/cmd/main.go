package main

import (
	"go-micro.dev/v4/logger"

	"github.com/kitanoyoru/kita/apps/email/internal/app"
	"github.com/kitanoyoru/kita/apps/email/internal/config"
)

const (
	ServiceName    = "emailservice"
	ServiceVersion = "0.1.0"

	ServiceConfigPath = "/etc/kita/kita-emailservice.yaml"
)

func main() {
	if err := config.Load(ServiceConfigPath); err != nil {
		logger.Fatal(err)
	}

	srv := app.NewApp(ServiceName, ServiceVersion)

	if err := srv.Init(); err != nil {
		logger.Fatal(err)
	}

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
