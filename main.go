package main

import (
	"bufio"
	"flag"
	"fmt"
	"go/build"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	templateName = "go-service-template"
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

type sources string

func (src sources) mainGo() string {
	return filepath.Join(string(src), "main.go")
}

func (src sources) dockerfile() string {
	return filepath.Join(string(src), "Dockerfile")
}

func substituteTemplateName(serviceName string) error {
	return filepath.Walk(serviceName, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".git") {
			return fmt.Errorf("must not contain git subdirs, found: '%s'", info.Name())
		}
		if info.IsDir() {
			return nil
		}
		body, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		return writeFile(path, strings.Replace(string(body), templateName, serviceName, -1))
	})
}

func cleanImports(dir string) error {
	lines := []string{}
	src := sources(dir)
	srcFile, err := os.Open(src.mainGo())
	if err != nil {
		return err
	}
	defer srcFile.Close()

	scanner := bufio.NewScanner(srcFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, templateName) && strings.HasPrefix(strings.TrimSpace(line), "\"github.com") {
			continue
		}

		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return writeFile(src.mainGo(), strings.Join(lines, "\n"))
}

func fixDockerfile(serviceName string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	cwd = strings.TrimPrefix(cwd, build.Default.GOPATH+"/src/")
	src := sources(serviceName)

	dockerCwd := fmt.Sprintf("%s/%s", cwd, serviceName)
	dockerCwdTemplate := "github.com/nchern/go-service-template/go-service-template"

	return replaceInFile(src.dockerfile(), dockerCwdTemplate, dockerCwd)
}

var svcName = flag.String("create", "", "service name to create")

func main() {
	log.SetFlags(log.Lshortfile)
	flag.Parse()
	currentDir := "."

	serviceName := strings.TrimSpace(*svcName)
	if serviceName == "" {
		flag.Usage()
		os.Exit(0)
	}

	must(RestoreAssets(currentDir, templateName))
	defer os.RemoveAll(templateName)

	must(os.Rename(templateName, serviceName))

	must(cleanImports(serviceName))

	must(fixDockerfile(serviceName))

	must(substituteTemplateName(serviceName))

	must(exec.Command("goimports", "-w", sources(serviceName).mainGo()).Run())
}
