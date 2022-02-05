package main

import (
	"bufio"
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	templateName = "go-service-template"
)

type sources string

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

func cleanImporsFromFile(path string) error {
	srcFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	lines := []string{}
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

	return writeFile(path, strings.Join(lines, "\n"))
}

func fixImports(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		isGoSourceFile := strings.HasSuffix(path, ".go") && !info.IsDir()
		if !isGoSourceFile {
			return nil
		}

		if err := cleanImporsFromFile(path); err != nil {
			return fmt.Errorf("cleanImporsFromFile: %w", err)
		}

		err = exec.Command("goimports", "-w", path).Run()
		if err != nil {
			return fmt.Errorf("goimports: %w", err)
		}
		return nil
	})
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

func create(serviceName string) {
	currentDir := "."

	restoreDir, notErr := ioutil.TempDir(currentDir, "template-")
	must(notErr)

	defer os.RemoveAll(restoreDir)

	must(RestoreAssets(restoreDir, templateName))

	templateDir := filepath.Join(restoreDir, templateName)

	must(os.Rename(templateDir, serviceName))

	must(fixDockerfile(serviceName))

	must(fixImports(serviceName))

	must(substituteTemplateName(serviceName))
}
