package main

import (
	"fmt"
	"os"
	"time"

	"github.com/akshaybabloo/dnode/cmd"
)

var (
	version = "0.0.1-dev"
	buildDate = time.Now().String()
)

func main() {
	rootCmd := cmd.NewRootCmd(version, buildDate)
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
