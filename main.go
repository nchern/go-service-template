package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func must(err error) {
	if err != nil {
		log.Fatalf("FATAL: %s\n", err)
	}
}

func writeFile(path string, body string) error {
	return ioutil.WriteFile(path, []byte(body), 0644)
}

func replaceInFile(fileName string, what string, to string) error {
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	return writeFile(fileName, strings.Replace(string(body), what, to, -1))
}

var svcName = flag.String("create", "", "service name to create")

func main() {
	log.SetFlags(log.Lshortfile)
	flag.Parse()

	serviceName := strings.TrimSpace(*svcName)
	if serviceName == "" {
		flag.Usage()
		os.Exit(0)
	}

	create(serviceName)
}
