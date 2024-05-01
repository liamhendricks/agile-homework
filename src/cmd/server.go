package cmd

import (
  "github.com/gin-gonic/gin"
  "github.com/spf13/cobra"
  "agile-homework/src/app"
  "agile-homework/src/http/routes"
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

		gin.SetMode(gin.DebugMode)
		router := gin.New()
		routes.Init(router, container)

    err := router.Run(config.HttpPort)
		if err != nil {
			panic(err)
		}
	},
}
