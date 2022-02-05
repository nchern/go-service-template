package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nchern/go-service-template/go-service-template/pkg/cli"
	"github.com/nchern/go-service-template/go-service-template/pkg/srv"
)

const (
	svcName = "go-service-template"

	shortHelp = "go-service-template is a ..."

	longHelp = "Long help"
)

var ()

func init() {
	cli.Init(
		svcName,
		longHelp,
		shortHelp,
		serveForever,
	)
}

func serveForever() {
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
		log.Fatalf("fatal: %s", err)
	}
}
