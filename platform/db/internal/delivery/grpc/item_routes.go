package grpc

import (
	"github.com/kitanoyoru/gigaservices/platform/db/internal/models"
	"github.com/kitanoyoru/gigaservices/platform/db/pkg/proto"
	"github.com/samber/lo"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateItem(ctx context.Context, req *proto.CreateItemRequest) (*proto.CreateItemResponse, error) {
	item := models.Item{
		CustomerId: req.CustomerId,
		Title:      req.Title,
		Price:      req.Price,
	}

	dbItem, err := s.db.Items.AddItem(ctx, &item)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.CreateItemResponse{
		Item: dbItem.ToProto(),
	}, nil
}

func (s *Server) GetItemById(ctx context.Context, req *proto.GetItemByIdRequest) (*proto.GetItemByIdResponse, error) {
	dbItem, err := s.db.Items.GetItemById(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetItemByIdResponse{
		Item: dbItem.ToProto(),
	}, nil
}

func (s *Server) ListItems(ctx context.Context, req *proto.ListItemsRequest) (*proto.ListItemsResponse, error) {
	dbItems, err := s.db.Items.GetAllItems(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.ListItemsResponse{
		Items: lo.Map(dbItems, func(item *models.Item, _ int) *proto.Item {
			return item.ToProto()
		}),
	}, nil
}
