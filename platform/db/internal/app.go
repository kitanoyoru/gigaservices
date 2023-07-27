package internal

import (
	"context"

	pkggrpc "github.com/kitanoyoru/gigaservices/pkg/grpc"
	"github.com/kitanoyoru/gigaservices/platform/db/internal/database"
	grpcimpl "github.com/kitanoyoru/gigaservices/platform/db/internal/delivery/grpc"
	"github.com/kitanoyoru/gigaservices/platform/db/internal/di"
	"github.com/kitanoyoru/gigaservices/platform/db/pkg/cfg"
	"github.com/kitanoyoru/gigaservices/platform/db/pkg/proto"
	"github.com/samber/do"
	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
)

var config cfg.Config

func init() {
	config, err := cfg.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %+v", err)
	}

	do.ProvideValue[*cfg.Config](di.Provider, config)

	if config.Debug {
		log.Info("Database Microservice is running on DEBUG mode")
		log.SetLevel(log.DebugLevel)
	}
}

func Run() error {
	db, err := database.NewDatabaseConnection()
	if err != nil {
		return err
	}

	do.ProvideValue[*database.DatabaseConnection](di.Provider, db)

	svc := grpcimpl.NewServer()

	return pkggrpc.NewServer(config.Server.Port, func(s *grpc.Server) {
		proto.RegisterDatabaseServiceServer(s, svc)
	}).Start(context.Background())

}
