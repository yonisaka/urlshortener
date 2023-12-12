package grpchandler

import (
	"context"
	rpc "github.com/yonisaka/urlshortener/api/go/grpc"
	"github.com/yonisaka/urlshortener/internal/entities/repository"
	"github.com/yonisaka/urlshortener/internal/usecases"
)

// URLShortenerServiceServer is URL Shortener service server contract.
type URLShortenerServiceServer interface {
	rpc.URLShortenerServiceServer
}

// NewURLShortenerHandler returns a new URLShortenerServiceServer.
func NewURLShortenerHandler(uc usecases.URLShortenerUsecase) URLShortenerServiceServer {
	return &urlShortenerHandler{
		uc: uc,
	}
}

type urlShortenerHandler struct {
	rpc.UnimplementedURLShortenerServiceServer
	uc usecases.URLShortenerUsecase
}

func (h *urlShortenerHandler) ListURLShortener(ctx context.Context, req *rpc.ListURLShortenerRequest) (*rpc.ListURLShortenerResponse, error) {
	return h.uc.ListURLShortener(ctx, &repository.ListURLShortenerParams{
		UserID:        req.GetUserId(),
		StartDateTime: req.GetStartDatetime().AsTime(),
		EndDateTime:   req.GetEndDatetime().AsTime(),
	})
}

func (h *urlShortenerHandler) CreateURLShortener(ctx context.Context, req *rpc.CreateURLShortenerRequest) (*rpc.URLShortener, error) {
	return h.uc.CreateURLShortener(ctx, &repository.CreateURLShortenerParams{
		UserID:   req.GetUserId(),
		URL:      req.GetUrl(),
		DateTime: req.GetDatetime().AsTime(),
	})
}

func (h *urlShortenerHandler) GetShortenedURL(ctx context.Context, req *rpc.GetShortenedURLRequest) (*rpc.URLShortener, error) {
	return h.uc.GetShortenedURL(ctx, &repository.GetShortenedURLParams{
		URL: req.GetUrl(),
	})
}
