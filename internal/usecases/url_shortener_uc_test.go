package usecases_test

import (
	"github.com/yonisaka/urlshortener/internal/di"
	"github.com/yonisaka/urlshortener/internal/entities/repository"
	"github.com/yonisaka/urlshortener/internal/usecases"
)

type fields struct {
	urlShortenerRepo repository.URLShortenerRepo
}

func sut(f fields) usecases.URLShortenerUsecase {
	return usecases.NewURLShortenerUsecase(
		f.urlShortenerRepo,
		di.GetTracer().Tracer(),
		di.GetLogger(),
	)
}
