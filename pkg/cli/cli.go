package cli

import (
	"github.com/spf13/cobra"
)

var (
	cliCmd = &cobra.Command{
		Use:   "cli",
		Short: "Exposes command line interface",
	}

	rootCmd *cobra.Command
)

// Init initialises the root command of the utility
// defaultAction
func Init(appName string, longHelp string, shortHelp string, defaultAction func()) {
	rootCmd = &cobra.Command{
		Use:   appName,
		Long:  longHelp,
		Short: shortHelp,
		Run: func(cmd *cobra.Command, args []string) {
			defaultAction()
		},
	}
	rootCmd.AddCommand(cliCmd)

	cliCmd.AddCommand(createCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
