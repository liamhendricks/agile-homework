package routes

import (
	"agile-homework/src/app"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine, container *app.ServiceContainer) {
  router.GET("health", health()) 
}

func health() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.String(200, "OK")
    return
  }
}
