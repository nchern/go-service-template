package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/nchern/go-service-template/go-service-template/pkg/cli"
	"github.com/nchern/go-service-template/go-service-template/pkg/srv"
)

const (
	svcName = "go-service-template"

	shortHelp = "%s is a ..."
)

var ()

func init() {
	cli.Init(
		svcName,
		"Long help",
		fmt.Sprintf(shortHelp, svcName),
		runAndWait,
	)
}

func runAndWait() {
	go srv.Start(svcName)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	for range c {
		srv.Stop()
		break
	}
}

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
