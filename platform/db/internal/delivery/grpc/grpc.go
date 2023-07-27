package grpc

import (
	"github.com/kitanoyoru/gigaservices/platform/db/internal/database"
	"github.com/kitanoyoru/gigaservices/platform/db/internal/di"
	"github.com/kitanoyoru/gigaservices/platform/db/internal/models"
	"github.com/kitanoyoru/gigaservices/platform/db/pkg/proto"
	"github.com/samber/do"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *Server) CreateCustomer(ctx context.Context, req *proto.CreateCustomerRequest) (*proto.CreateCustomerResponse, error) {
	customer := models.Customer{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	dbCustomer, err := s.db.customers.AddCustomer(ctx, &customer)
	if err != nil {
		return nil, status.Error(codes.Internal, err)
	}

	return &proto.CreateCustomerResponse{
		Customer: dbCustomer,
	}, nil
}
