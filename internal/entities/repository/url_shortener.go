package repository

import (
	"context"
	rpc "github.com/yonisaka/urlshortener/api/go/grpc"
	"time"
)

// ListURLShortenerParams is a parameter for ListURLShortener.
type ListURLShortenerParams struct {
	UserID        int64
	StartDateTime time.Time
	EndDateTime   time.Time
}

// CreateURLShortenerParams is a parameter for CreateURLShortener.
type CreateURLShortenerParams struct {
	UserID       int64
	URL          string
	ShortenedURL string
	DateTime     time.Time
}

// GetShortenedURLParams is a parameter for GetShortenedURL.
type GetShortenedURLParams struct {
	URL string
}

// URLShortenerRepo is a repository for URLShortener.
type URLShortenerRepo interface {
	// ListURLShortener returns a list of URLShortener.
	ListURLShortener(ctx context.Context, params *ListURLShortenerParams) ([]*rpc.URLShortener, error)
	// CreateURLShortener creates a URLShortener.
	CreateURLShortener(ctx context.Context, params *CreateURLShortenerParams) (*rpc.URLShortener, error)
	// GetShortenedURL returns a URLShortener.
	GetShortenedURL(ctx context.Context, params *GetShortenedURLParams) (*rpc.URLShortener, error)
}
