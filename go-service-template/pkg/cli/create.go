package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd is the particular implementation of `create` subcommand.
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create: an example of a cli command")
	},
}
