package grpc

import (
	"github.com/kitanoyoru/gigaservices/platform/db/internal/database"
	"github.com/kitanoyoru/gigaservices/platform/db/internal/di"
	"github.com/kitanoyoru/gigaservices/platform/db/pkg/proto"
	"github.com/samber/do"
)

var _ proto.DatabaseServiceServer = (*Server)(nil)

type Server struct {
	proto.UnimplementedDatabaseServiceServer

	db *database.DatabaseConnection
}

func NewServer() *Server {
	db := do.MustInvoke[*database.DatabaseConnection](di.Provider)
	return &Server{
		db: db,
	}
}
