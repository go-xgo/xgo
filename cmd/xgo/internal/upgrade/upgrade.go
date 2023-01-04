package upgrade

import (
	"fmt"
	"github.com/go-xgo/xgo/cmd/xgo/internal/base"
	"github.com/spf13/cobra"
)

// CmdUpgrade represents the upgrade command.
var CmdUpgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the xgo tools",
	Long:  "Upgrade the xgo tools. Example: xgo upgrade",
	Run:   Run,
}

// Run upgrade the xgo tools.
func Run(cmd *cobra.Command, args []string) {
	fmt.Println("Running upgrade")
	if err := base.GoInstall(
		"github.com/go-xgo/xgo/cmd/xgo@latest",
	); err != nil {
		fmt.Println(err)
	}
	if err := base.GoUpdate(
		"github.com/go-xgo/xgo@latest",
	); err != nil {
		fmt.Println(err)
	}
}
