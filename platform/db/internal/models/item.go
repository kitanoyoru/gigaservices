package models

import "github.com/kitanoyoru/gigaservices/platform/db/pkg/proto"

type Item struct {
	Id         string `json:"id,omitempty"`
	CustomerId string `json:"customer_id"`
	Title      string `json:"title"`
	Price      int64  `json:"price"`
}

func (i *Item) ToProto() *proto.Item {
	return &proto.Item{
		Id:         i.Id,
		CustomerId: i.CustomerId,
		Title:      i.Title,
		Price:      i.Price,
	}
}
