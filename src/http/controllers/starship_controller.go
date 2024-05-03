package controllers

import (
	"errors"
	"net/http"

	"agile-homework/src/app"
	"agile-homework/src/clients/swapi"

	"github.com/gin-gonic/gin"
)

type StarshipController struct {
	config      app.Config
	swapiClient app.SwapiClient
}

func NewStarshipController(c app.Config, sc app.SwapiClient) *StarshipController {
	return &StarshipController{
		swapiClient: sc,
		config:      c,
	}
}

func (cc *StarshipController) GetStarship(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, errors.New("bad request"))
		return
	}

	res, err := cc.swapiClient.Starship(c, id)
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
