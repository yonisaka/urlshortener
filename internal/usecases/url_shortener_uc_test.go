package usecases_test

import (
	"github.com/yonisaka/urlshortener/internal/di"
	"github.com/yonisaka/urlshortener/internal/entities/repository"
	"github.com/yonisaka/urlshortener/internal/usecases"
)

type fields struct {
	btcRepo repository.URLShortenerRepo
}

func sut(f fields) usecases.URLShortenerUsecase {
	return usecases.NewURLShortenerUsecase(
		f.btcRepo,
		di.GetTracer().Tracer(),
		di.GetLogger(),
	)
}
