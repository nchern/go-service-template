package cli

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd *cobra.Command
)

// Init initialises the root command of the utility
//  - defaultAction will be executed if no sub-command is given
func Init(appName string, longHelp string, shortHelp string, defaultAction func()) {
	rootCmd = &cobra.Command{
		Use:   appName,
		Long:  longHelp,
		Short: shortHelp,
		Run: func(cmd *cobra.Command, args []string) {
			defaultAction()
		},
	}
}

// Execute is an entry point to CLI
func Execute() error {
	return rootCmd.Execute()
}
