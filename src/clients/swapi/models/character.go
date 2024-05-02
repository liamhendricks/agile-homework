package models

import (
  "encoding/json"
  "strings"
)

type Character struct {
	Name      string   `json:"name"`
	Homeworld string   `json:"homeworld"`
	Species   []string `json:"species"`
	Starship  []string `json:"starships"`
}

// override Unmarshal to strip swapi urls
func (c *Character) UnmarshalJSON(data []byte) error {
	type _character struct {
    Name      string   `json:"name"`
    Homeworld string   `json:"homeworld"`
    Species   []string `json:"species"`
    Starship  []string `json:"starships"`
	}

	var char _character
	if err := json.Unmarshal(data, &char); err != nil {
		return err
	}

  var species []string
  var ships []string

  for _, sp := range char.Species {
    species = append(species, strings.Replace(sp, SwapiBaseUrl + AllSpeciesPath, "", 1))
  }
  for _, ss := range char.Starship {
    ships = append(ships, strings.Replace(ss, SwapiBaseUrl + AllStarshipsPath, "", 1))
  }

	c.Homeworld = strings.Replace(char.Homeworld, SwapiBaseUrl + AllPlanetsPath, "", 1)
  c.Species = species
  c.Starship = ships
  c.Name = char.Name
  return nil
}
