package clients

import (
  "net/http"
  "context"
  "agile-homework/src/clients/swapi/models"
)

const (
  swapiBaseUrl = "https://swapi.dev/api"
)

// interact with swapi.dev api
type SwapiClient struct {
  baseUrl string
  httpClient *http.Client
}

func NewSwapiClient() *SwapiClient {
  return &SwapiClient{
    baseUrl: swapiBaseUrl,
    httpClient: http.DefaultClient,
  }
}

func (c *SwapiClient) Characters(ctx context.Context, query string) ([]models.Character, error) {
  var characters []models.Character

  return characters, nil
}
