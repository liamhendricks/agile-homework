package cmd

import (
  "github.com/spf13/cobra"
  "github.com/spf13/viper"
)

var RootCommand = &cobra.Command{
    Use:   "api",
    Short: "Root command for api.",
}

func init() {
  viper.AutomaticEnv()
}
