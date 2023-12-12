package di

import "github.com/yonisaka/urlshortener/internal/usecases"

// GetURLShortenerUsecase returns URLShortenerUsecase instance.
func GetURLShortenerUsecase() usecases.URLShortenerUsecase {
	return usecases.NewURLShortenerUsecase(
		GetURLShortenerRepo(),
		GetTracer().Tracer(),
		GetLogger(),
	)
}
