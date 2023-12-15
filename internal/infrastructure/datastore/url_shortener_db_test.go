package datastore_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	rpc "github.com/yonisaka/urlshortener/api/go/grpc"
	"github.com/yonisaka/urlshortener/internal/di"
	"github.com/yonisaka/urlshortener/internal/entities/repository"
	"github.com/yonisaka/urlshortener/internal/infrastructure/datastore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	var code int

	defer func() {
		os.Exit(code)
	}()

	_ = os.Setenv("APP_ENV", "test")
	_ = os.Setenv("IS_REPLICA", "false")

	code = m.Run()
}

func TestURLShortenerRepo_CreateURLShortener(t *testing.T) {
	type args struct {
		ctx    context.Context
		params *repository.CreateURLShortenerParams
	}

	type test struct {
		args       args
		want       *rpc.URLShortener
		wantErr    error
		beforeFunc func(*testing.T)
		afterFunc  func(*testing.T)
	}

	db := datastore.GetDatabaseMaster()

	tests := map[string]func(t *testing.T) test{
		"Given valid query of Create url shortener, When query executed successfully, Return no error": func(t *testing.T) test {
			userID := int64(402)
			url := "https://www.test.com"
			shortenedURL := "https://localhost/3k4j5l6"
			datetime := time.Now().UTC()

			args := args{
				ctx: context.Background(),
				params: &repository.CreateURLShortenerParams{
					UserID:       userID,
					URL:          url,
					ShortenedURL: shortenedURL,
					DateTime:     datetime,
				},
			}

			want := &rpc.URLShortener{
				UserId:       userID,
				OriginalUrl:  url,
				ShortenedUrl: shortenedURL,
				Datetime:     timestamppb.New(datetime),
			}

			return test{
				args:    args,
				want:    want,
				wantErr: nil,
				beforeFunc: func(t *testing.T) {
					t.Helper()

					// Remove existing data, if any.
					_, err := db.Exec(context.Background(), "DELETE FROM users WHERE id = $1", userID)
					assert.NoError(t, err)

					// Insert test data.
					_, err = db.Exec(context.Background(), "INSERT INTO users (id, name) VALUES ($1, $2)", userID, "test")
					assert.NoError(t, err)
				},
				afterFunc: func(t *testing.T) {
					t.Helper()

					var err error

					// Clear data.
					_, err = db.Exec(context.Background(), "DELETE FROM url_shorteners WHERE user_id = $1", userID)
					assert.NoError(t, err)

					_, err = db.Exec(context.Background(), "DELETE FROM users WHERE id = $1", userID)
					assert.NoError(t, err)
				},
			}
		},
		"Given valid query of Create url shortener, When query executed successfully with no User found, Return an error": func(t *testing.T) test {
			userID := int64(404)
			url := "https://www.test.com"
			shortenedURL := "https://localhost/3k4j5l6"
			datetime := time.Now().UTC()

			args := args{
				ctx: context.Background(),
				params: &repository.CreateURLShortenerParams{
					UserID:       userID,
					URL:          url,
					ShortenedURL: shortenedURL,
					DateTime:     datetime,
				},
			}

			return test{
				args:    args,
				want:    nil,
				wantErr: status.Error(codes.NotFound, datastore.ErrNotFound),
				beforeFunc: func(t *testing.T) {
					t.Helper()

					_, err := db.Exec(context.Background(), "DELETE FROM users WHERE id = $1", userID)
					assert.NoError(t, err)
				},
			}
		},
	}

	for name, fn := range tests {
		t.Run(name, func(t *testing.T) {
			tt := fn(t)

			if tt.beforeFunc != nil {
				tt.beforeFunc(t)
			}

			if tt.afterFunc != nil {
				defer tt.afterFunc(t)
			}

			sut := di.GetURLShortenerRepo()

			got, err := sut.CreateURLShortener(tt.args.ctx, tt.args.params)

			if !assert.ErrorIs(t, err, tt.wantErr) {
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestURLShortenerRepo_ListURLShortener(t *testing.T) {
	type args struct {
		ctx    context.Context
		params *repository.ListURLShortenerParams
	}

	type test struct {
		args       args
		want       []*rpc.URLShortener
		wantErr    error
		beforeFunc func(*testing.T)
		afterFunc  func(*testing.T)
	}

	db := datastore.GetDatabaseMaster()

	tests := map[string]func(t *testing.T) test{
		"Given valid query of Get List url shortener, When query executed successfully, Return no error": func(t *testing.T) test {
			userID := int64(402)
			url := "https://www.test.com"
			shortenedURL := "https://localhost/3k4j5l6"

			// 2023-02-12 02:35:38 +0000 UTC
			datetime := &timestamppb.Timestamp{
				Seconds: 1676169338,
				Nanos:   0,
			}
			args := args{
				ctx: context.Background(),
				params: &repository.ListURLShortenerParams{
					UserID:        userID,
					StartDateTime: datetime.AsTime().Add(-1 * time.Hour),
					EndDateTime:   datetime.AsTime().Add(1 * time.Hour),
				},
			}

			want := []*rpc.URLShortener{
				{
					UserId:       userID,
					OriginalUrl:  url,
					ShortenedUrl: shortenedURL,
					Datetime:     datetime,
				},
			}

			return test{
				args:    args,
				want:    want,
				wantErr: nil,
				beforeFunc: func(t *testing.T) {
					t.Helper()

					// Remove existing data, if any.
					_, err := db.Exec(context.Background(), "DELETE FROM url_shorteners WHERE user_id = $1", userID)
					assert.NoError(t, err)

					_, err = db.Exec(context.Background(), "DELETE FROM users WHERE id = $1", userID)
					assert.NoError(t, err)

					// Insert test data.
					_, err = db.Exec(context.Background(), "INSERT INTO users (id, name) VALUES ($1, $2)", userID, "test")
					assert.NoError(t, err)

					_, err = db.Exec(context.Background(), "INSERT INTO url_shorteners (datetime, user_id, original_url, shortened_url) VALUES ($1, $2, $3, $4)", datetime.AsTime(), userID, url, shortenedURL)
					assert.NoError(t, err)
				},
				afterFunc: func(t *testing.T) {
					t.Helper()

					// Clear data.
					_, err := db.Exec(context.Background(), "DELETE FROM url_shorteners WHERE user_id = $1", userID)
					assert.NoError(t, err)

					_, err = db.Exec(context.Background(), "DELETE FROM users WHERE id = $1", userID)
					assert.NoError(t, err)
				},
			}
		},
	}

	for name, fn := range tests {
		t.Run(name, func(t *testing.T) {
			tt := fn(t)

			if tt.beforeFunc != nil {
				tt.beforeFunc(t)
			}

			if tt.afterFunc != nil {
				defer tt.afterFunc(t)
			}

			sut := di.GetURLShortenerRepo()

			got, err := sut.ListURLShortener(tt.args.ctx, tt.args.params)

			if !assert.ErrorIs(t, err, tt.wantErr) {
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestURLShortenerRepo_GetShortenedURL(t *testing.T) {
	type args struct {
		ctx    context.Context
		params *repository.GetShortenedURLParams
	}

	type test struct {
		args       args
		want       *rpc.URLShortener
		wantErr    error
		beforeFunc func(*testing.T)
		afterFunc  func(*testing.T)
	}

	db := datastore.GetDatabaseMaster()

	tests := map[string]func(t *testing.T) test{
		"Given valid query of Get url shortened, When query executed successfully, Return no error": func(t *testing.T) test {
			url := "https://www.test.com"
			shortenedURL := "https://localhost/3k4j5l6"
			userID := int64(401)

			// 2023-02-12 02:35:38 +0000 UTC
			datetime := &timestamppb.Timestamp{
				Seconds: 1676169338,
				Nanos:   0,
			}

			args := args{
				ctx: context.Background(),
				params: &repository.GetShortenedURLParams{
					URL: url,
				},
			}

			want := &rpc.URLShortener{
				UserId:       userID,
				OriginalUrl:  url,
				ShortenedUrl: shortenedURL,
				Datetime:     datetime,
			}

			return test{
				args:    args,
				want:    want,
				wantErr: nil,
				beforeFunc: func(t *testing.T) {
					t.Helper()

					// Remove existing data, if any.
					_, err := db.Exec(context.Background(), "DELETE FROM users WHERE id = $1", userID)
					assert.NoError(t, err)

					// Insert test data.
					_, err = db.Exec(context.Background(), "INSERT INTO users (id, name) VALUES ($1, $2)", userID, "test")
					assert.NoError(t, err)

					_, err = db.Exec(context.Background(), "INSERT INTO url_shorteners (datetime, user_id, original_url, shortened_url) VALUES ($1, $2, $3, $4)", datetime.AsTime(), userID, url, shortenedURL)
					assert.NoError(t, err)
				},
				afterFunc: func(t *testing.T) {
					t.Helper()

					// Clear data.
					_, err := db.Exec(context.Background(), "DELETE FROM url_shorteners WHERE user_id = $1", userID)
					assert.NoError(t, err)

					_, err = db.Exec(context.Background(), "DELETE FROM users WHERE id = $1", userID)
					assert.NoError(t, err)
				},
			}
		},
		"Given valid query of Get url shortened, When query executed successfully with no URL found, Return an error": func(t *testing.T) test {
			url := "https://www.test.com"
			userID := int64(401)

			args := args{
				ctx: context.Background(),
				params: &repository.GetShortenedURLParams{
					URL: url,
				},
			}

			return test{
				args:    args,
				want:    nil,
				wantErr: status.Error(codes.NotFound, datastore.ErrNotFound),
				beforeFunc: func(t *testing.T) {
					t.Helper()

					// Remove existing data, if any.
					_, err := db.Exec(context.Background(), "DELETE FROM url_shorteners WHERE user_id = $1", userID)
					assert.NoError(t, err)

					_, err = db.Exec(context.Background(), "DELETE FROM users WHERE id = $1", userID)
					assert.NoError(t, err)
				},
			}
		},
	}

	for name, fn := range tests {
		t.Run(name, func(t *testing.T) {
			tt := fn(t)

			if tt.beforeFunc != nil {
				tt.beforeFunc(t)
			}

			if tt.afterFunc != nil {
				defer tt.afterFunc(t)
			}

			sut := di.GetURLShortenerRepo()

			got, err := sut.GetShortenedURL(tt.args.ctx, tt.args.params)

			if !assert.ErrorIs(t, err, tt.wantErr) {
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
