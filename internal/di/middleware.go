package di

import (
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

// GetMiddleware get the grpc middlewares.
func GetMiddleware() []grpc.ServerOption {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
	}

	return opts
}
