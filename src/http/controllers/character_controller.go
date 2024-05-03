package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"agile-homework/src/app"
	"agile-homework/src/clients/swapi"
	"agile-homework/src/clients/swapi/models"

	"github.com/gin-gonic/gin"
)

const (
	baseCharacterPath = "/api/characters"
	basePlanetPath    = "/api/planets"
	baseStarshipPath  = "/api/starships"
	baseSpeciesPath   = "/api/species"
)

type CharacterController struct {
	config      app.Config
	swapiClient app.SwapiClient
	preload     chan interface{}
}

func NewCharacterController(c app.Config, sc app.SwapiClient) *CharacterController {
	return &CharacterController{
		swapiClient: sc,
		config:      c,
	}
}

func (cc *CharacterController) GetAllCharacters(c *gin.Context) {
	search, _ := c.GetQuery("search")
	page, _ := c.GetQuery("page")

	res, err := cc.swapiClient.Characters(c, search, page)
	if err != nil {
		var notFoundErr swapi.SwapiNotFoundError
		if errors.As(err, &notFoundErr) {
			c.JSON(err.(swapi.SwapiNotFoundError).Code, err)
			return
		}

		c.JSON(http.StatusBadRequest, err)
		return
	}

	// add our proxy urls as the pagination
  if res.Next != nil {
    *res.Next = fmt.Sprintf("%s%s%s", cc.config.BaseUrl, baseCharacterPath, *res.Next)
  }

  if res.Previous != nil {
    *res.Previous = fmt.Sprintf("%s%s%s", cc.config.BaseUrl, baseCharacterPath, *res.Previous)
  }

	c.JSON(http.StatusOK, res)
	return
}

// preload all data
func (cc *CharacterController) PreloadCharacterData(c *gin.Context) {
	search, _ := c.GetQuery("search")
	page, _ := c.GetQuery("page")
  var preloadedResp swapi.SwapiResponse[models.PreloadedCharacter]

	res, err := cc.swapiClient.Characters(c, search, page)
	if err != nil {
		var notFoundErr swapi.SwapiNotFoundError
		if errors.As(err, &notFoundErr) {
			c.JSON(err.(swapi.SwapiNotFoundError).Code, err)
			return
		}

		c.JSON(http.StatusBadRequest, err)
		return
	}
  preloadedResp.Next = res.Next
  preloadedResp.Count = res.Count
  preloadedResp.Results = cc.preloadCharacterData(c, res)

  c.JSON(http.StatusOK, preloadedResp)
  return
}

func (cc *CharacterController) preloadCharacterData(
  c *gin.Context,
  data swapi.SwapiResponse[models.Character],
) []models.PreloadedCharacter {
  var results []models.PreloadedCharacter
  pcc := make(chan models.PreloadedCharacter) 
  for _, character := range data.Results {
    go cc.fetchCharacterDataAsync(c, character, pcc)
  }

  // read the results from the channel
  for result := range pcc {
    results = append(results, result)
    if len(results) == len(data.Results) {
      close(pcc)
    }
  }

  return results
}

func (cc *CharacterController) fetchCharacterDataAsync(
  c *gin.Context,
  char models.Character,
  result chan models.PreloadedCharacter,
) {
  var preloadedCharacter models.PreloadedCharacter
  preloadedCharacter.Homeworld, _ = cc.swapiClient.Planet(c, char.Homeworld)

  for _, id := range char.Species {
    s, _ := cc.swapiClient.Species(c, id)
    preloadedCharacter.Species = append(preloadedCharacter.Species, s)
  }
  for _, id := range char.Starship {
    s, _ := cc.swapiClient.Starship(c, id)
    preloadedCharacter.Starship = append(preloadedCharacter.Starship, s)
  }
  result <- preloadedCharacter
  return
}
