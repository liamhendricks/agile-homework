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

type PreloadedCharacter struct {
	Name      string     `json:"name"`
	Homeworld Planet     `json:"homeworld"`
	Species   []Species  `json:"species"`
	Starship  []Starship `json:"starship"`
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
		species = append(species, getID(sp, AllSpeciesPath))
	}
	for _, ss := range char.Starship {
		ships = append(ships, getID(ss, AllStarshipsPath))
	}

	c.Homeworld = getID(char.Homeworld, AllPlanetsPath)
	c.Species = species
	c.Starship = ships
	c.Name = char.Name
	return nil
}

func getID(s string, path string) string {
	return strings.TrimRight(strings.Replace(s, SwapiBaseUrl+path, "", 1), "/")
}
