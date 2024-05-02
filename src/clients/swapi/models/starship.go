package models

type Starship struct {
	Name          string `json:"name"`
	CargoCapacity string `json:"cargo_capacity"`
	StarshipClass string `json:"starship_class"`
}
