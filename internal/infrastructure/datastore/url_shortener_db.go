package datastore

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	rpc "github.com/yonisaka/urlshortener/api/go/grpc"
	"github.com/yonisaka/urlshortener/internal/entities/repository"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type urlShortenerRepo struct {
	*BaseRepo
}

// NewURLShortenerRepo returns a new URLShortenerRepo.
func NewURLShortenerRepo(base *BaseRepo) repository.URLShortenerRepo {
	return &urlShortenerRepo{
		BaseRepo: base,
	}
}

// ListURLShortener returns a list of URLShortener.
func (r *urlShortenerRepo) ListURLShortener(ctx context.Context, params *repository.ListURLShortenerParams) ([]*rpc.URLShortener, error) {
	query := `SELECT original_url, shortened_url, datetime 
				FROM url_shorteners 
				WHERE user_id = $1 
			  	AND datetime >= $2::timestamptz AND datetime <= $3::timestamptz`

	rows, err := r.dbSlave.Query(ctx, query, params.UserID, params.StartDateTime, params.EndDateTime)
	if err != nil {
		return nil, fmt.Errorf("failed to query url shorteners: %w", err)
	}
	defer rows.Close()

	var urlShorteners []*rpc.URLShortener

	for rows.Next() {
		var urlShortener rpc.URLShortener
		if err := rows.Scan(&urlShortener.OriginalUrl, &urlShortener.ShortenedUrl, &urlShortener.Datetime); err != nil {
			return nil, fmt.Errorf("failed to scan url shortener: %w", err)
		}

		urlShorteners = append(urlShorteners, &urlShortener)
	}

	return urlShorteners, nil
}

// CreateURLShortener creates a URLShortener.
func (r *urlShortenerRepo) CreateURLShortener(ctx context.Context, params *repository.CreateURLShortenerParams) (*rpc.URLShortener, error) {
	var id int64
	var err error

	err = r.dbSlave.QueryRow(ctx, "SELECT id FROM url_shorteners WHERE original_url = $1", params.URL).Scan(&id)
	if !errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("failed to check if url already exists: %w", err)
	}

	if id != 0 {
		return nil, fmt.Errorf("url already exists")
	}

	err = r.dbMaster.QueryRow(ctx,
		"INSERT INTO url_shorteners (user_id, original_url, shortened_url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		params.UserID, params.URL, params.ShortenedURL, params.DateTime, params.DateTime).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("failed to create url: %w", err)
	}

	return &rpc.URLShortener{
		OriginalUrl: params.URL,
		Datetime:    timestamppb.New(params.DateTime),
	}, nil
}

// GetShortenedURL returns a URLShortener.
func (r *urlShortenerRepo) GetShortenedURL(ctx context.Context, params *repository.GetShortenedURLParams) (*rpc.URLShortener, error) {
	var urlShortener *rpc.URLShortener

	err := r.dbSlave.QueryRow(ctx, "SELECT original_url, shortened_url, datetime FROM url_shorteners WHERE original_url = $1", params.URL).
		Scan(&urlShortener.OriginalUrl, &urlShortener.ShortenedUrl, &urlShortener.Datetime)
	if err != nil {
		return nil, fmt.Errorf("failed to get url: %w", err)
	}

	return urlShortener, nil
}
