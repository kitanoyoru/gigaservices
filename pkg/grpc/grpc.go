package grpc

import (
	"context"
	"fmt"
	"net"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/kitanoyoru/gigaservices/pkg/grpc/interceptor"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	channelz "google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	server *grpc.Server
	port   int
}

func NewServer(port int, register func(server *grpc.Server)) *Server {
	interceptors := []grpc.UnaryServerInterceptor{
		interceptor.NewRequestLogger(),
	}

	opts := []grpc.ServerOption{
		middleware.WithUnaryServerChain(interceptors...),
	}

	server := grpc.NewServer(opts...)

	register(server)

	reflection.Register(server)
	channelz.RegisterChannelzServiceToServer(server)

	return &Server{
		server: server,
		port:   port,
	}
}

func (s *Server) Start(ctx context.Context) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("Failed to listen on port %d: %+v", s.port, err)
	}

	errCh := make(chan error, 1)
	go func() {
		if err := s.server.Serve(listener); err != nil {
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		if err != nil {
			logrus.Fatalf("GRPC server has stopped with error %+v", err)
			return err
		}
		return nil
	case <-ctx.Done():
		s.server.GracefulStop()
		return <-errCh
	}
}
