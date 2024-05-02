package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"agile-homework/src/app"
	"agile-homework/src/clients/swapi"

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
}

func NewCharacterController(c app.Config, sc app.SwapiClient) *CharacterController {
	return &CharacterController{
		swapiClient: sc,
		config:      c,
	}
}

func (cc *CharacterController) GetAllCharacters(c *gin.Context) {
	searchQuery, _ := c.GetQuery("search")

	res, err := cc.swapiClient.Characters(c, searchQuery)
	if err != nil {
		var notFoundErr swapi.SwapiNotFoundError
		if errors.As(err, &notFoundErr) {
			c.JSON(err.(swapi.SwapiNotFoundError).Code, err)
			return
		}

		c.JSON(http.StatusBadRequest, err)
		return
	}

	// add our proxy url as the pagination
	*res.Next = fmt.Sprintf("%s%s%s", cc.config.BaseUrl, baseCharacterPath, *res.Next)

	c.JSON(http.StatusOK, res)
	return
}

