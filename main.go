package main

import (
	"agile-homework/src/cmd"
	"fmt"
	"os"
)

func main() {
	if err := cmd.RootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
