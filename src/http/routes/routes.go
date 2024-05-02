package routes

import (
	"agile-homework/src/app"
	"agile-homework/src/http/controllers"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine, container *app.ServiceContainer) {
  characterController := controllers.NewCharacterController(container.SwapiClient)

  router.GET("/health", health()) 

  api := router.Group("/api")
  characterRoutes := api.Group("/characters")
  characterRoutes.GET("/all", characterController.GetAllCharacters)
}

func health() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.String(200, "OK")
    return
  }
}
