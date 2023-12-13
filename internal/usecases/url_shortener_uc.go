package usecases

import (
	"context"
	rpc "github.com/yonisaka/urlshortener/api/go/grpc"
	"github.com/yonisaka/urlshortener/internal/entities/repository"
	"github.com/yonisaka/urlshortener/pkg/logging"
	"github.com/yonisaka/urlshortener/pkg/trace"
)

//go:generate rm -f ./url_shortener_uc_mock.go
//go:generate mockgen -destination url_shortener_uc_mock.go -package usecases -mock_names URLShortenerUsecase=GoMockURLShortenerUsecase -source url_shortener_uc.go

// URLShortenerUsecase is a usecase for URLShortener.
type URLShortenerUsecase interface {
	// ListURLShortener returns a list of URLShortener.
	ListURLShortener(ctx context.Context, params *repository.ListURLShortenerParams) (*rpc.ListURLShortenerResponse, error)
	// CreateURLShortener creates a URLShortener.
	CreateURLShortener(ctx context.Context, params *repository.CreateURLShortenerParams) (*rpc.URLShortener, error)
	// GetShortenedURL returns a URLShortener.
	GetShortenedURL(ctx context.Context, params *repository.GetShortenedURLParams) (*rpc.URLShortener, error)
}

// compile time interface implementation check.
var _ URLShortenerUsecase = (*urlShortenerUsecase)(nil)

// NewURLShortenerUsecase returns a new URLShortenerUsecase.
func NewURLShortenerUsecase(
	urlShortenerRepo repository.URLShortenerRepo,
	trace trace.Tracer,
	logger logging.Logger,
) URLShortenerUsecase {
	return &urlShortenerUsecase{
		urlShortenerRepo: urlShortenerRepo,
		trace:            trace,
		logger:           logger,
	}
}

type urlShortenerUsecase struct {
	urlShortenerRepo repository.URLShortenerRepo
	trace            trace.Tracer
	logger           logging.Logger
}
