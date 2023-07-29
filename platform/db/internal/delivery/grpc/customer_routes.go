package grpc

import (
	"github.com/kitanoyoru/gigaservices/platform/db/internal/models"
	"github.com/kitanoyoru/gigaservices/platform/db/pkg/proto"
	"github.com/samber/lo"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCustomer(ctx context.Context, req *proto.CreateCustomerRequest) (*proto.CreateCustomerResponse, error) {
	customer := models.Customer{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	dbCustomer, err := s.db.Customers.AddCustomer(ctx, &customer)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.CreateCustomerResponse{
		Customer: dbCustomer.ToProto(),
	}, nil
}

func (s *Server) GetCustomerById(ctx context.Context, req *proto.GetCustomerByIdRequest) (*proto.GetCustomerByIdResponse, error) {
	dbCustomer, err := s.db.Customers.GetCustomerById(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetCustomerByIdResponse{
		Customer: dbCustomer.ToProto(),
	}, nil
}

func (s *Server) GetCustomerByEmail(ctx context.Context, req *proto.GetCustomerByEmailRequest) (*proto.GetCustomerByEmailResponse, error) {
	dbCustomer, err := s.db.Customers.GetCustomerByEmail(ctx, req.Email)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetCustomerByEmailResponse{
		Customer: dbCustomer.ToProto(),
	}, nil
}

func (s *Server) ListCustomers(ctx context.Context, req *proto.ListCustomersRequest) (*proto.ListCustomersResponse, error) {
	dbCustomers, err := s.db.Customers.GetAllCustomers(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.ListCustomersResponse{
		Customers: lo.Map(dbCustomers, func(customer *models.Customer, _ int) *proto.Customer {
			return customer.ToProto()
		}),
	}, nil
}
