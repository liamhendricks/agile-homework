package app

import (
	"agile-homework/src/clients/swapi"
	"agile-homework/src/clients/swapi/models"
	"context"
)

type ServiceContainer struct {
	Config      Config
	SwapiClient SwapiClient
}

type SwapiClient interface {
	Characters(ctx context.Context, search, page string) (swapi.SwapiResponse[models.Character], error)
	Planet(ctx context.Context, id string) (models.Planet, error)
	Species(ctx context.Context, id string) (models.Species, error)
	Starship(ctx context.Context, id string) (models.Starship, error)
}

func GetApp(config Config) *ServiceContainer {
	return &ServiceContainer{Config: config}
}
