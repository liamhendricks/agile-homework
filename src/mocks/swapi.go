package mocks

import (
	"context"

	"agile-homework/src/clients/swapi"
	"agile-homework/src/clients/swapi/models"
)

type MockSwapiClient struct {
	ExpectedCharacters swapi.SwapiResponse[models.Character]
	ExpectedPlanets    models.Planet
	ExpectedSpecies    models.Species
	ExpectedStarships  models.Starship
	ExpectedErr        error
}

func (ms *MockSwapiClient) Characters(ctx context.Context, search, page string) (swapi.SwapiResponse[models.Character], error) {
	return ms.ExpectedCharacters, ms.ExpectedErr
}
func (ms *MockSwapiClient) Planet(ctx context.Context, id string) (models.Planet, error) {
	return ms.ExpectedPlanets, ms.ExpectedErr
}
func (ms *MockSwapiClient) Species(ctx context.Context, id string) (models.Species, error) {
	return ms.ExpectedSpecies, ms.ExpectedErr
}
func (ms *MockSwapiClient) Starship(ctx context.Context, id string) (models.Starship, error) {
	return ms.ExpectedStarships, ms.ExpectedErr
}
