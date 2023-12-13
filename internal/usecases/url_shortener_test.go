package usecases_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	rpc "github.com/yonisaka/urlshortener/api/go/grpc"
	"github.com/yonisaka/urlshortener/internal/entities/repository"
	"google.golang.org/protobuf/types/known/timestamppb"
	"os"
	"testing"
	"time"
)

var errInternal = errors.New("error")

func TestMain(m *testing.M) {
	var code int

	defer func() {
		os.Exit(code)
	}()

	_ = os.Setenv("BASE_URL", "localhost")
	_ = os.Setenv("CODE_SIZE", "8")

	code = m.Run()
}

func TestURLShortenerUC_CreateURLShortener(t *testing.T) {
	type args struct {
		ctx    context.Context
		params *repository.CreateURLShortenerParams
	}

	type test struct {
		fields  fields
		args    args
		want    *rpc.URLShortener
		wantErr error
	}

	tests := map[string]func(t *testing.T, ctrl *gomock.Controller) test{
		"Given valid request of Create url shortener, When repository executed successfully, Return no error": func(t *testing.T, ctrl *gomock.Controller) test {
			ctx := context.Background()

			userID := int64(1)
			url := "https://www.test.com"
			shortenedURL := "https://localhost/3k4j5l6"
			now := time.Now()

			args := args{
				ctx: ctx,
				params: &repository.CreateURLShortenerParams{
					UserID:   userID,
					URL:      url,
					DateTime: now,
				},
			}

			want := &rpc.URLShortener{
				UserId:       userID,
				OriginalUrl:  url,
				ShortenedUrl: shortenedURL,
				Datetime:     timestamppb.New(now),
			}

			mockJourneyRepo := repository.NewGoMockURLShortenerRepo(ctrl)
			mockJourneyRepo.EXPECT().CreateURLShortener(args.ctx, args.params).Return(want, nil)

			return test{
				fields: fields{
					btcRepo: mockJourneyRepo,
				},
				args:    args,
				want:    want,
				wantErr: nil,
			}
		},
		"Given valid request of Create url shortener, When repository failed to executed, Return an error": func(t *testing.T, ctrl *gomock.Controller) test {
			ctx := context.Background()

			userID := int64(1)
			url := "https://www.test.com"
			now := time.Now()

			args := args{
				ctx: ctx,
				params: &repository.CreateURLShortenerParams{
					UserID:   userID,
					URL:      url,
					DateTime: now,
				},
			}

			mockJourneyRepo := repository.NewGoMockURLShortenerRepo(ctrl)
			mockJourneyRepo.EXPECT().CreateURLShortener(args.ctx, args.params).Return(nil, errInternal)

			return test{
				fields: fields{
					btcRepo: mockJourneyRepo,
				},
				args:    args,
				want:    nil,
				wantErr: errInternal,
			}
		},
	}

	for name, testFn := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tt := testFn(t, ctrl)

			sut := sut(tt.fields)

			got, err := sut.CreateURLShortener(tt.args.ctx, tt.args.params)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestURLShortenerUC_ListURLShortener(t *testing.T) {
	type args struct {
		ctx    context.Context
		params *repository.ListURLShortenerParams
	}

	type test struct {
		fields  fields
		args    args
		want    *rpc.ListURLShortenerResponse
		wantErr error
	}

	tests := map[string]func(t *testing.T, ctrl *gomock.Controller) test{
		"Given valid request of List url shortener, When repository executed successfully without cache, Return no error": func(t *testing.T, ctrl *gomock.Controller) test {
			ctx := context.Background()

			userID := int64(1)
			url := "https://www.test.com"
			shortenedURL := "https://localhost/3k4j5l6"
			now := time.Now()

			args := args{
				ctx: ctx,
				params: &repository.ListURLShortenerParams{
					UserID:        userID,
					StartDateTime: now,
					EndDateTime:   now,
				},
			}

			urlShorteners := []*rpc.URLShortener{
				{
					UserId:       userID,
					Datetime:     timestamppb.New(now),
					OriginalUrl:  url,
					ShortenedUrl: shortenedURL,
				},
			}

			want := &rpc.ListURLShortenerResponse{
				UrlShortener: urlShorteners,
			}

			mockJourneyRepo := repository.NewGoMockURLShortenerRepo(ctrl)
			mockJourneyRepo.EXPECT().ListURLShortener(args.ctx, args.params).Return(urlShorteners, nil)

			return test{
				fields: fields{
					btcRepo: mockJourneyRepo,
				},
				args:    args,
				want:    want,
				wantErr: nil,
			}
		},
	}

	for name, testFn := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tt := testFn(t, ctrl)

			sut := sut(tt.fields)

			got, err := sut.ListURLShortener(tt.args.ctx, tt.args.params)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestURLShortenerUC_GetShortenedURL(t *testing.T) {
	type args struct {
		ctx    context.Context
		params *repository.GetShortenedURLParams
	}

	type test struct {
		fields  fields
		args    args
		want    *rpc.URLShortener
		wantErr error
	}

	tests := map[string]func(t *testing.T, ctrl *gomock.Controller) test{
		"Given valid request of Get shortened URL, When repository executed successfully without cache, Return no error": func(t *testing.T, ctrl *gomock.Controller) test {
			ctx := context.Background()

			userID := int64(1)
			url := "https://www.test.com"
			shortenedURL := "https://localhost/3k4j5l6"
			now := time.Now()

			args := args{
				ctx: ctx,
				params: &repository.GetShortenedURLParams{
					URL: url,
				},
			}

			want := &rpc.URLShortener{
				UserId:       userID,
				OriginalUrl:  url,
				ShortenedUrl: shortenedURL,
				Datetime:     timestamppb.New(now),
			}

			mockJourneyRepo := repository.NewGoMockURLShortenerRepo(ctrl)
			mockJourneyRepo.EXPECT().GetShortenedURL(args.ctx, args.params).Return(want, nil)

			return test{
				fields: fields{
					btcRepo: mockJourneyRepo,
				},
				args:    args,
				want:    want,
				wantErr: nil,
			}
		},
	}

	for name, testFn := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tt := testFn(t, ctrl)

			sut := sut(tt.fields)

			got, err := sut.GetShortenedURL(tt.args.ctx, tt.args.params)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
