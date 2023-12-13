package di

import (
	"github.com/yonisaka/urlshortener/internal/adapters/grpchandler"
)

// GetURLShortenerGRPCHandler returns URLShortenerServiceServer handler.
func GetURLShortenerGRPCHandler() grpchandler.URLShortenerServiceServer {
	return grpchandler.NewURLShortenerHandler(GetURLShortenerUsecase())
}
