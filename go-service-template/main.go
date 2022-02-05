package main

import (
	"log"

	"github.com/nchern/go-service-template/go-service-template/pkg/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		log.Fatalf("fatal: %s", err)
	}
}
