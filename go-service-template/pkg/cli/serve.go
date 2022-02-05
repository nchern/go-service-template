package cli

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/nchern/go-service-template/go-service-template/pkg/srv"
	"github.com/spf13/cobra"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "runs the service and serves forewer",
		// Args:  Set args, e.g. cobra.ExactArgs(1),

		// SilenceErrors: true,
		// SilenceUsage:  true,

		Run: func(cmd *cobra.Command, args []string) {
			serveForever()
		},
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serveForever() {
	go srv.Start(appName)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	for range c {
		srv.Stop()
		break
	}
}
