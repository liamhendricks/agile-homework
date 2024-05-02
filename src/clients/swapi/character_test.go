package swapi

import (
	"agile-homework/src/clients/swapi/models"
	"encoding/json"
	"testing"
)

const testJson = `
{
	"count": 82,
	"next": "https://swapi.dev/api/people/?page=2",
	"previous": null,
	"results": [
		{
			"name": "Luke Skywalker",
			"height": "172",
			"mass": "77",
			"hair_color": "blond",
			"skin_color": "fair",
			"eye_color": "blue",
			"birth_year": "19BBY",
			"gender": "male",
			"homeworld": "https://swapi.dev/api/planets/1/",
			"films": [
				"https://swapi.dev/api/films/1/",
				"https://swapi.dev/api/films/2/",
				"https://swapi.dev/api/films/3/",
				"https://swapi.dev/api/films/6/"
			],
			"species": [],
			"vehicles": [
				"https://swapi.dev/api/vehicles/14/",
				"https://swapi.dev/api/vehicles/30/"
			],
			"starships": [
				"https://swapi.dev/api/starships/12/",
				"https://swapi.dev/api/starships/22/"
			],
			"created": "2014-12-09T13:50:51.644000Z",
			"edited": "2014-12-20T21:17:56.891000Z",
			"url": "https://swapi.dev/api/people/1/"
		}
  ]
}`

func TestSwapiResponseUnmarshal(t *testing.T) {
  var swapiResponse SwapiResponse[models.Character]

  err := json.Unmarshal([]byte(testJson), &swapiResponse)
  if err != nil {
    t.Fatal("error during unmarshal")
  }

  if len(swapiResponse.Results) == 0 {
    t.Fatal("no results")
  }

  if swapiResponse.Results[0].Name != "Luke Skywalker" {
    t.Fatal("bad name")
  }

  if len(swapiResponse.Results[0].Starship) == 0 {
    t.Fatal("bad decode")
  }

  if swapiResponse.Results[0].Starship[0] != "12/" {
    t.Fatal("bad decode")
  }

  if swapiResponse.Results[0].Homeworld != "1/" {
    t.Fatal("bad decode")
  }
}
