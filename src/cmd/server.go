package cmd

import (
	"agile-homework/src/app"
	"agile-homework/src/clients/swapi"
	"agile-homework/src/http/routes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(serverCommand)
}

var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "Runs the gin server.",
	Run: func(cmd *cobra.Command, args []string) {
    config := app.GetConfig()
    container := app.GetApp(config)
    container.SwapiClient = swapi.NewSwapiClient()

		gin.SetMode(gin.DebugMode)
		router := gin.New()
		routes.Init(router, container)

    err := router.Run(config.HttpPort)
		if err != nil {
			panic(err)
		}
	},
}
