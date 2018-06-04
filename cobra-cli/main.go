package main

import (
	"log"

	"github.com/matsu0228/go_sandbox/cobra-cli/cmd"
)

func main() {
	rootCmd := cmd.NewCmdRoot()
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
