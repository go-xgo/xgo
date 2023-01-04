package main

import (
	"github.com/go-xgo/xgo/cmd/xgo/internal/run"
	"github.com/go-xgo/xgo/cmd/xgo/internal/upgrade"
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
	rootCmd.AddCommand(upgrade.CmdUpgrade)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
