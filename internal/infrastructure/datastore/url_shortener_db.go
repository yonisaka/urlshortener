package datastore

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	rpc "github.com/yonisaka/urlshortener/api/go/grpc"
	"github.com/yonisaka/urlshortener/internal/entities/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

const (
	// ErrNotFound is an error for indicates record not found.
	ErrNotFound = "error not found"
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
// with the given params UserID, StartDateTime, EndDateTime.
func (r *urlShortenerRepo) ListURLShortener(ctx context.Context, params *repository.ListURLShortenerParams) ([]*rpc.URLShortener, error) {
	query := `SELECT user_id, original_url, shortened_url, datetime 
				FROM url_shorteners 
				WHERE user_id = $1 
			  	AND datetime >= $2::timestamptz AND datetime <= $3::timestamptz`

	rows, err := r.dbSlave.Query(ctx, query, params.UserID, params.StartDateTime, params.EndDateTime)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to query url shorteners: %s", err))
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
			return nil, status.Error(codes.Internal, fmt.Sprintf("failed to scan url shortener: %s", err))
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
// with the given params UserID, URL, ShortenedURL, DateTime.
func (r *urlShortenerRepo) CreateURLShortener(ctx context.Context, params *repository.CreateURLShortenerParams) (*rpc.URLShortener, error) {
	var id int64
	var err error

	err = r.dbSlave.QueryRow(ctx, "SELECT id FROM users WHERE id = $1", params.UserID).Scan(&id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, status.Error(codes.NotFound, ErrNotFound)
	}

	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to check if user exists: %s", err))
	}

	// check if url already exists
	var shortenedURL string
	err = r.dbSlave.QueryRow(ctx, "SELECT shortened_url FROM url_shorteners WHERE original_url = $1", params.URL).Scan(&shortenedURL)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to check if url already exists: %s", err))
	}

	if shortenedURL != "" {
		return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("url: %s already exists with shortened url: %s", params.URL, shortenedURL))
	}

	err = r.dbMaster.QueryRow(ctx,
		"INSERT INTO url_shorteners (user_id, original_url, shortened_url, datetime) VALUES ($1, $2, $3, $4) RETURNING id",
		params.UserID, params.URL, params.ShortenedURL, params.DateTime).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to create url: %s", err))
	}

	return &rpc.URLShortener{
		UserId:       params.UserID,
		OriginalUrl:  params.URL,
		ShortenedUrl: params.ShortenedURL,
		Datetime:     timestamppb.New(params.DateTime),
	}, nil
}

// GetShortenedURL returns a URLShortener.
// with the given params URL (original_url).
func (r *urlShortenerRepo) GetShortenedURL(ctx context.Context, params *repository.GetShortenedURLParams) (*rpc.URLShortener, error) {
	// t is a temporary struct for storing the result of the query.
	var t struct {
		UserID       int64
		OriginalURL  string
		ShortenedURL string
		Datetime     time.Time
	}

	err := r.dbSlave.QueryRow(ctx, "SELECT user_id, original_url, shortened_url, datetime FROM url_shorteners WHERE original_url = $1", params.URL).
		Scan(&t.UserID, &t.OriginalURL, &t.ShortenedURL, &t.Datetime)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, status.Error(codes.NotFound, ErrNotFound)
	}

	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to get url: %s", err))
	}

	return &rpc.URLShortener{
		UserId:       t.UserID,
		OriginalUrl:  t.OriginalURL,
		ShortenedUrl: t.ShortenedURL,
		Datetime:     timestamppb.New(t.Datetime),
	}, nil
}
