package internal

import (
	"context"

	pkggrpc "github.com/kitanoyoru/gigaservices/pkg/grpc"
	"github.com/kitanoyoru/gigaservices/platform/db/internal/database"
	"github.com/kitanoyoru/gigaservices/platform/db/internal/delivery/grpc"
	"github.com/kitanoyoru/gigaservices/platform/db/internal/di"
	"github.com/kitanoyoru/gigaservices/platform/db/pkg/cfg"
	"github.com/kitanoyoru/gigaservices/platform/db/pkg/proto"
	"github.com/samber/do"
)

type App struct {
	config *cfg.DatabaseConfig

	db *database.DatabaseConnection
}

func NewApp() *App {
	config := do.MustInvoke[*cfg.Config](di.Provider)
	return &App{
		config: config.Database,
	}
}

func (app *App) Initialize() error {
	db, err := database.NewDatabaseConnection()
	if err != nil {
		return err
	}

	app.db = db
	do.ProvideValue[*database.DatabaseConnection](di.Provider, db)

	return nil
}

func (app *App) Run() {
	svc := grpc.NewServer()

	return pkggrpc.NewServer(app.config.Port, func(s *grpc.Server) {
		proto.RegisterDatabaseServiceServer(s, svc)
	}).Start(context.Background())
}
