package main

import (
	"bufio"
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

func readFlag(i int, defaultVal string) string {
	if len(os.Args) <= i {
		return defaultVal
	}
	return os.Args[i]
}

func mainGo(serviceName string) string {
	return filepath.Join(serviceName, "main.go")
}

func dockerfile(serviceName string) string {
	return filepath.Join(serviceName, "Dockerfile")
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
	srcFile, err := os.Open(mainGo(dir))
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

	return writeFile(mainGo(dir), strings.Join(lines, "\n"))
}

func fixDockerfile(serviceName string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	cwd = strings.TrimPrefix(cwd, build.Default.GOPATH+"/src/")

	dockerCwd := fmt.Sprintf("%s/%s", cwd, serviceName)
	dockerCwdTemplate := "github.com/nchern/go-service-template/go-service-template"

	return replaceInFile(dockerfile(serviceName), dockerCwdTemplate, dockerCwd)
}

func showHelp() {
	fmt.Printf("Usage go-service-template create <service-name> [<currentDir>]")
	os.Exit(0)
}

func parseFlags() (currentDir string, serviceName string, err error) {
	currentDir = readFlag(3, ".")
	serviceName = "my-svc"
	return
}
func main() {
	log.SetFlags(log.Lshortfile)

	currentDir, serviceName, _ := parseFlags()

	must(RestoreAssets(currentDir, templateName))

	must(os.Rename(templateName, serviceName))

	must(cleanImports(serviceName))

	must(fixDockerfile(serviceName))

	must(substituteTemplateName(serviceName))

	must(exec.Command("goimports", "-w", mainGo(serviceName)).Run())
}
