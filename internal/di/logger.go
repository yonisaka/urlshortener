package di

import (
	"log"

	"github.com/yonisaka/urlshortener/pkg/logging"
)

// GetLogger get the logger wrapper.
func GetLogger() logging.Logger {
	l, err := logging.NewLogger()
	if err != nil {
		log.Fatal(err)
	}

	return l
}
