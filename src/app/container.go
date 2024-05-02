package app

import (
  "context"
  "agile-homework/src/clients/swapi/models"
)

type ServiceContainer struct {
  config Config
  SwapiClient SwapiClient
}

type SwapiClient interface {
  Characters(ctx context.Context, query string) ([]models.Character, error)
}

func GetApp(config Config) *ServiceContainer {
  return &ServiceContainer{config: config}
}
