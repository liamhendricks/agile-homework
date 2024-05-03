package models

const (
  SwapiBaseUrl = "https://swapi.dev/api"
  AllCharactersPath = "/people/"
  AllPlanetsPath = "/planets/"
  AllSpeciesPath = "/species/"
  AllStarshipsPath = "/starships/"
)

type Gettable interface {
  Character | PreloadedCharacter | Planet | Species | Starship
}

