package handlers

import (
	"context"

	"go-micro.dev/v4/logger"

	pb "github.com/kitanoyoru/kita/apps/emailservice/pkg/proto"
)

type Email struct{}

func (s *Email) SendOrderConfirmation(ctx context.Context, in *pb.SendOrderConfirmationRequest, out *pb.Empty) error {
	logger.Infof("A request to send order confirmation email to %s has been received.", in.Email)
	return nil
}
