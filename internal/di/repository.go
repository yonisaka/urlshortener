package di

import (
	"github.com/yonisaka/urlshortener/internal/entities/repository"
	"github.com/yonisaka/urlshortener/internal/infrastructure/datastore"
)

// GetBaseRepo returns BaseRepo instance.
func GetBaseRepo() *datastore.BaseRepo {
	return datastore.NewBaseRepo(datastore.GetDatabaseMaster(), datastore.GetDatabaseSlave())
}

// GetURLShortenerRepo returns URLShortenerRepo instance.
func GetURLShortenerRepo() repository.URLShortenerRepo {
	return datastore.NewURLShortenerRepo(GetBaseRepo())
}
