package usecases

import (
	"context"
	"fmt"
	rpc "github.com/yonisaka/urlshortener/api/go/grpc"
	"github.com/yonisaka/urlshortener/internal/entities/repository"
	"go.uber.org/zap"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func (u *urlShortenerUsecase) ListURLShortener(ctx context.Context, params *repository.ListURLShortenerParams) (*rpc.ListURLShortenerResponse, error) {
	ctx, span := u.trace.StartSpan(ctx, "UC.ListURLShortener", nil)
	defer span.End()

	urlShorteners, err := u.urlShortenerRepo.ListURLShortener(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to list url shorteners: %w", err)
	}

	list := &rpc.ListURLShortenerResponse{
		UrlShortener: urlShorteners,
	}

	return list, nil
}

func (u *urlShortenerUsecase) CreateURLShortener(ctx context.Context, params *repository.CreateURLShortenerParams) (*rpc.URLShortener, error) {
	ctx, span := u.trace.StartSpan(ctx, "UC.CreateURLShortener", nil)
	defer span.End()

	shortenedURL, err := u.generateShortenedURL()
	if err != nil {
		u.logger.Warn("failed to generate shortened url", zap.Error(err))
		return nil, fmt.Errorf("failed to generate shortened url: %w", err)
	}

	params.ShortenedURL = shortenedURL

	return u.urlShortenerRepo.CreateURLShortener(ctx, params)
}

func (u *urlShortenerUsecase) GetShortenedURL(ctx context.Context, params *repository.GetShortenedURLParams) (*rpc.URLShortener, error) {
	ctx, span := u.trace.StartSpan(ctx, "UC.GetShortenedURL", nil)
	defer span.End()

	return u.urlShortenerRepo.GetShortenedURL(ctx, params)
}

func (u *urlShortenerUsecase) generateShortenedURL() (string, error) {
	baseURL := os.Getenv("BASE_URL")
	codeSize := os.Getenv("CODE_SIZE")

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	codeSizeInt, err := strconv.Atoi(codeSize)
	if err != nil {
		return "", fmt.Errorf("failed to convert code size to int: %w", err)
	}

	rand.Seed(time.Now().UnixNano())
	shortCode := make([]byte, codeSizeInt)
	for i := range shortCode {
		shortCode[i] = charset[rand.Intn(len(charset))]
	}

	return fmt.Sprintf("%s/%s", baseURL, string(shortCode)), nil
}
