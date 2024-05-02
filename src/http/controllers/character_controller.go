package controllers

import (
	"agile-homework/src/app"
	"agile-homework/src/clients/swapi/models"
	"github.com/gin-gonic/gin"
)

type CharacterController struct {
  swapiClient app.SwapiClient
}

func NewCharacterController(sc app.SwapiClient) *CharacterController {
  return &CharacterController{
    swapiClient: sc,
  }
}

type CharactersResponse struct {
	Characters  []models.Character `json:"characters"`
	SearchQuery string             `json:"search_query"`
  Total int
}

func (cc *CharacterController) GetAllCharacters(c *gin.Context) {
  searchQuery, ok := c.GetQuery("search")
  if !ok {
    //
  }

  characters, err := cc.swapiClient.Characters(c, searchQuery)
  if err != nil {
    //
  }

  response := CharactersResponse{
    SearchQuery: searchQuery,
    Characters: characters,
    Total: len(characters),
  }

  c.JSON(200, response)
  return
}
