package cli

import (
	"github.com/spf13/cobra"
)

const (
	appName = "go-service-template"

	shortHelp = "go-service-template is a ..."

	longHelp = "Long help"
)

var (
	rootCmd *cobra.Command
)

func init() {
	rootCmd = &cobra.Command{
		Use:   appName,
		Long:  longHelp,
		Short: shortHelp,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
}

// Execute is an entry point to CLI
func Execute() error {
	return rootCmd.Execute()
}
