package pkg

import (
	"errors"

	"github.com/meilisearch/meilisearch-go"
)

func InitDB() (client *meilisearch.Client, err error) {
	client = meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://localhost:7700",
		APIKey: "admin123",
	})

	if client == nil {
		return nil, errors.New("error when try to connect to meilisearch")
	}
	return
}
