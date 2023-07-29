package models

import "github.com/kitanoyoru/gigaservices/platform/db/pkg/proto"

type Customer struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *Customer) ToProto() *proto.Customer {
	return &proto.Customer{
		Id:       c.Id,
		Name:     c.Name,
		Email:    c.Email,
		Password: c.Password,
	}
}
