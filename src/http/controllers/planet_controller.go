package controllers

import (

	"errors"
  "net/http"

	"agile-homework/src/app"
	"agile-homework/src/clients/swapi"

	"github.com/gin-gonic/gin"
)

type PlanetController struct {
	config      app.Config
	swapiClient app.SwapiClient
}

func NewPlanetController(c app.Config, sc app.SwapiClient) *PlanetController {
	return &PlanetController{
		swapiClient: sc,
		config:      c,
	}
}

func (cc *PlanetController) GetPlanet(c *gin.Context) {
	id := c.Param("id")
  if id == "" {
		c.JSON(http.StatusBadRequest, errors.New("bad request"))
		return
  }

	res, err := cc.swapiClient.Planet(c, id)
	if err != nil {
		var notFoundErr swapi.SwapiNotFoundError
		if errors.As(err, &notFoundErr) {
			c.JSON(err.(swapi.SwapiNotFoundError).Code, err)
			return
		}

		c.JSON(http.StatusBadRequest, err)
		return
	}

  c.JSON(http.StatusOK, res)
  return
}
