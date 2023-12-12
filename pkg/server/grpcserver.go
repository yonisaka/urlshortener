package server

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

// HandlerRegister is a function to register gRPC handler to the given gRPC server.
type HandlerRegister func(*grpc.Server)

// NewGRPCServer returns Server implementation of gRPC server.
func NewGRPCServer(port string, register HandlerRegister, opts ...grpc.ServerOption) (Server, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return nil, fmt.Errorf("failed to open listener port %s: %w", port, err)
	}

	s := &gRPCServer{
		lis:    lis,
		server: grpc.NewServer(opts...),
	}

	register(s.server)

	return s, nil
}

type gRPCServer struct {
	lis    net.Listener
	server *grpc.Server
}

func (s *gRPCServer) Run() error {
	if err := s.server.Serve(s.lis); err != nil {
		return fmt.Errorf("failed to start a server: %w", err)
	}

	return nil
}

func (s *gRPCServer) GracefulStop() {
	s.server.GracefulStop()
}
