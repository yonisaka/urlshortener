package datastore

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	rpc "github.com/yonisaka/urlshortener/api/go/grpc"
	"github.com/yonisaka/urlshortener/internal/entities/repository"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

var (
	// ErrNotFound is an error for indicates record not found.
	ErrNotFound = errors.New("error not found")
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
	query := `SELECT user_id, original_url, shortened_url, datetime 
				FROM url_shorteners 
				WHERE user_id = $1 
			  	AND datetime >= $2::timestamptz AND datetime <= $3::timestamptz`

	rows, err := r.dbSlave.Query(ctx, query, params.UserID, params.StartDateTime, params.EndDateTime)
	if err != nil {
		return nil, fmt.Errorf("failed to query url shorteners: %w", err)
	}
	defer rows.Close()

	var urlShorteners []*rpc.URLShortener

	type data struct {
		UserID       int64
		OriginalURL  string
		ShortenedURL string
		Datetime     time.Time
	}

	for rows.Next() {
		var d data
		if err := rows.Scan(&d.UserID, &d.OriginalURL, &d.ShortenedURL, &d.Datetime); err != nil {
			return nil, fmt.Errorf("failed to scan url shortener: %w", err)
		}

		urlShorteners = append(urlShorteners, &rpc.URLShortener{
			UserId:       d.UserID,
			OriginalUrl:  d.OriginalURL,
			ShortenedUrl: d.ShortenedURL,
			Datetime:     timestamppb.New(d.Datetime),
		})
	}

	return urlShorteners, nil
}

// CreateURLShortener creates a URLShortener.
func (r *urlShortenerRepo) CreateURLShortener(ctx context.Context, params *repository.CreateURLShortenerParams) (*rpc.URLShortener, error) {
	var id int64
	var err error

	err = r.dbSlave.QueryRow(ctx, "SELECT id FROM users WHERE id = $1", params.UserID).Scan(&id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("user id: %d not found: %w", params.UserID, ErrNotFound)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to check if user exists: %w", err)
	}

	var shortenedURL string
	err = r.dbSlave.QueryRow(ctx, "SELECT shortened_url FROM url_shorteners WHERE original_url = $1", params.URL).Scan(&shortenedURL)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("failed to check if url already exists: %w", err)
	}

	if shortenedURL != "" {
		return nil, fmt.Errorf("url: %s already exists with shortened url: %s", params.URL, shortenedURL)
	}

	err = r.dbMaster.QueryRow(ctx,
		"INSERT INTO url_shorteners (user_id, original_url, shortened_url, datetime) VALUES ($1, $2, $3, $4) RETURNING id",
		params.UserID, params.URL, params.ShortenedURL, params.DateTime).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("failed to create url: %w", err)
	}

	return &rpc.URLShortener{
		UserId:       params.UserID,
		OriginalUrl:  params.URL,
		ShortenedUrl: params.ShortenedURL,
		Datetime:     timestamppb.New(params.DateTime),
	}, nil
}

// GetShortenedURL returns a URLShortener.
func (r *urlShortenerRepo) GetShortenedURL(ctx context.Context, params *repository.GetShortenedURLParams) (*rpc.URLShortener, error) {
	type data struct {
		UserID       int64
		OriginalURL  string
		ShortenedURL string
		Datetime     time.Time
	}

	var d data

	err := r.dbSlave.QueryRow(ctx, "SELECT user_id, original_url, shortened_url, datetime FROM url_shorteners WHERE original_url = $1", params.URL).
		Scan(&d.UserID, &d.OriginalURL, &d.ShortenedURL, &d.Datetime)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("url: %s not found: %w", params.URL, ErrNotFound)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get url: %w", err)
	}

	return &rpc.URLShortener{
		UserId:       d.UserID,
		OriginalUrl:  d.OriginalURL,
		ShortenedUrl: d.ShortenedURL,
		Datetime:     timestamppb.New(d.Datetime),
	}, nil
}
