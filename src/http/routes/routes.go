package routes

import (
	"agile-homework/src/app"
	"agile-homework/src/http/controllers"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine, container *app.ServiceContainer) {
  characterController := controllers.NewCharacterController(container.Config, container.SwapiClient)
  planetController := controllers.NewPlanetController(container.Config, container.SwapiClient)
  speciesController := controllers.NewSpeciesController(container.Config, container.SwapiClient)
  starshipsController := controllers.NewStarshipController(container.Config, container.SwapiClient)

  router.GET("/health", health()) 

  api := router.Group("/api")
  characterRoutes := api.Group("/characters")
  characterRoutes.GET("", characterController.GetAllCharacters)
  characterRoutes.GET("/preload", characterController.PreloadCharacterData)

  planetRoutes := api.Group("/planets")
  planetRoutes.GET(":id", planetController.GetPlanet)

  speciesRoutes := api.Group("/species")
  speciesRoutes.GET(":id", speciesController.GetSpecies)

  starshipRoutes := api.Group("/starships")
  starshipRoutes.GET(":id", starshipsController.GetStarship)
}

func health() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.String(200, "OK")
    return
  }
}
