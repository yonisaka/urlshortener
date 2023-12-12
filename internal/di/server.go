package di

import (
	"log"
	"os"
	"sync"

	rpc "github.com/yonisaka/urlshortener/api/go/grpc"
	"github.com/yonisaka/urlshortener/pkg/di"
	"github.com/yonisaka/urlshortener/pkg/server"

	"google.golang.org/grpc"
)

var (
	grpcServerOnce sync.Once
	btcServer      server.Server
)

// GetURLShortenerGRPCServer returns gRPC server instance for URLShortener service.
func GetURLShortenerGRPCServer() server.Server {
	return getGRPCServer(btcServer, func(server *grpc.Server) {
		h := GetURLShortenerGRPCHandler()
		rpc.RegisterURLShortenerServiceServer(server, h)
	})
}

// getGRPCServer
func getGRPCServer(grpcServer server.Server, register server.HandlerRegister) server.Server {
	grpcServerOnce.Do(func() {
		opts := GetMiddleware()

		port := os.Getenv("SERVER_PORT")

		s, err := server.NewGRPCServer(port, register, opts...)
		if err != nil {
			log.Fatal("gRPC server", err)
		}

		di.RegisterCloser("gRPC server", di.NewCloser(s.GracefulStop))

		grpcServer = s
	})

	return grpcServer
}
