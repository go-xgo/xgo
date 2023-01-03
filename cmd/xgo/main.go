package main

import (
	"go-xgo/xgo/cmd/xgo/internal/run"
	"log"
)
import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:     "xgo",
	Short:   "xgo: A simple toolkit for Go microservices.",
	Long:    `xgo: A simple toolkit for Go microservices.`,
	Version: release,
}

func init() {
	rootCmd.AddCommand(run.CmdRun)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
