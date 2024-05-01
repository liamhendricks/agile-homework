package main

import (
	"fmt"
	"agile-homework/src/cmd"
  "os"
)

func main() {
  if err := cmd.RootCommand.Execute(); err != nil {
		fmt.Println(err)
    os.Exit(-1)
	}
}
