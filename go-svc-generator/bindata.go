// Code generated by go-bindata.
// sources:
// go-service-template/Dockerfile
// go-service-template/Makefile
// go-service-template/README.md
// go-service-template/deploy/README.md
// go-service-template/docker-compose.yml
// go-service-template/main.go
// go-service-template/pkg/README.md
// go-service-template/pkg/cli/cli.go
// go-service-template/pkg/srv/srv.go
// go-service-template/scripts/build.sh
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _goServiceTemplateDockerfile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xcd\x4a\xec\x40\x10\x85\xf7\xf5\x14\x05\xb3\xb9\x77\x91\xc4\xec\x44\x70\xa3\x46\x18\x64\x92\x10\x47\x74\x96\x9d\x4e\xd9\x29\xed\x3f\xba\x6b\xc4\xc7\x97\xe8\x08\x2a\x59\x76\x9f\x43\x9d\xef\xbb\x1d\xba\x1d\x9a\x60\x95\x37\x17\x75\x59\xd7\xa8\x32\x8e\x47\xb6\x13\x25\x80\xc7\x6e\xb8\xbb\xd9\x0e\x58\x99\x50\xe5\xa4\x2b\xc3\x32\x1f\xc7\x52\x07\x57\x79\x3d\x53\xf2\x95\x09\x45\xa6\xf4\xc6\x9a\x0a\x21\x17\xad\x12\x5a\xfb\x83\xeb\xae\x3f\x60\x89\x25\xc0\x06\xaf\x96\xf3\x28\x33\x61\x4c\xe1\x85\xb4\x20\xfb\xcc\x13\xa1\xf2\xc8\x5e\x28\x39\x9a\x58\x09\xa1\x0e\x5e\x14\x7b\x4a\x30\x3c\xb4\xe8\xd4\x2b\x2d\x4d\x51\xd6\x16\x13\xc5\xfc\xfd\x00\xf8\x2b\xf1\x83\x1c\xbe\xa6\x8b\xe2\x39\x05\x77\x79\x32\xfb\x14\x1a\x79\x15\x1f\x57\xf9\x61\x83\x3b\xc5\x1e\x4f\x01\xc6\x90\x04\x9a\xa7\xbe\xbb\x6f\xf0\xfc\x6c\x89\xff\x85\x28\x1c\xbc\xb2\xff\xd1\x91\x24\xd6\x19\xe9\x3d\x86\xcc\xde\xfc\x6e\xd7\x00\x4d\xbb\x1f\x0e\x7d\xb7\x6d\xf7\xeb\x6b\x1f\x01\x00\x00\xff\xff\xb5\xa6\xf8\x2f\x95\x01\x00\x00")

func goServiceTemplateDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_goServiceTemplateDockerfile,
		"go-service-template/Dockerfile",
	)
}

func goServiceTemplateDockerfile() (*asset, error) {
	bytes, err := goServiceTemplateDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "go-service-template/Dockerfile", size: 405, mode: os.FileMode(420), modTime: time.Unix(1552344077, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _goServiceTemplateMakefile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x92\x31\x6f\x23\x21\x10\x85\x6b\xf3\x2b\x9e\x64\x17\x76\xb1\xb8\x5f\x29\x92\x5d\x58\xbe\x14\x89\xaf\x70\x8a\xab\x4e\x18\x26\x2c\x3a\x96\x59\xc1\xac\x23\xff\xfb\x13\xbb\xb1\x12\x39\xcd\xc0\x7b\xfa\x60\xde\x08\x5e\xf7\x2f\x87\x27\xcf\x4d\xa1\x7c\x0d\x96\x1a\xa1\x7e\x88\x46\x48\x9d\xf7\xc7\xa7\xba\x29\xa2\x9e\x5f\xf6\xc7\xc3\xdf\x89\x5c\xad\xeb\xb2\x69\x57\xeb\xf3\xfe\xb8\x51\xa7\xb7\xf3\xdd\x52\x6a\x89\x43\xca\xc1\x76\x90\x8e\x60\xb9\xef\x4d\x72\x05\x17\x8a\xfc\x81\x8f\x20\x1d\xcc\x30\x64\x1e\x72\x30\x42\x88\xec\x83\x45\x78\x47\x22\x4b\xa5\x98\x7c\x53\x4b\x3c\x27\xc9\xec\x46\x1b\x92\x87\x41\xe9\x28\x46\x14\x9b\xc3\x20\x6b\xd2\x5e\x87\x24\x0c\xbd\x9d\x9d\xb2\xc5\x3b\x47\x47\x79\x03\xcb\x63\x74\xb8\xd4\xa6\xc9\x06\x47\x99\x1c\x4c\x81\x81\x67\x76\x18\xb2\xb1\x12\x2c\x29\xa5\x7f\xff\x3a\xbd\xfe\x69\x11\x52\x11\x13\x63\xe3\x68\x28\xea\xbb\x68\xd5\x62\x89\xb1\x10\x6e\x3c\x66\xd8\x8e\x0b\x25\x38\x1a\x28\x39\x4a\xf6\x86\xde\x24\xe3\xa9\xa7\x24\x28\xb7\x22\xd4\xab\xc5\xce\x33\x3c\x09\x1a\x07\xbd\xd5\x5a\x7f\x75\x89\x21\x89\xaa\xa5\x55\x0b\xec\x3c\xf7\x24\xa6\x4a\xca\x8f\xe4\x95\x44\x5d\x69\xe2\x3c\x57\xf5\x08\x5c\xc6\x10\x9d\x9a\xea\xe7\x65\xb3\x85\x86\x71\x09\x69\xbb\x5a\x9f\xde\xce\x1b\xe8\x1f\x23\xde\xa7\x6b\x31\x3d\xe4\x7c\xf4\xd3\x7b\x6c\x32\x11\xb5\xcc\x81\xa6\xc9\xaa\x44\x93\x8d\xa5\x47\xda\xb1\xfd\x47\xb9\x99\x93\x7d\x17\xad\x5a\xec\x66\x7d\xcf\x28\x58\xad\xbf\xbe\x50\x8d\xf9\x3f\x00\x00\xff\xff\xf4\x12\x29\x1f\x75\x02\x00\x00")

func goServiceTemplateMakefileBytes() ([]byte, error) {
	return bindataRead(
		_goServiceTemplateMakefile,
		"go-service-template/Makefile",
	)
}

func goServiceTemplateMakefile() (*asset, error) {
	bytes, err := goServiceTemplateMakefileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "go-service-template/Makefile", size: 629, mode: os.FileMode(420), modTime: time.Unix(1552345179, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _goServiceTemplateReadmeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x4c\x49\x51\x48\xcc\xc9\x51\xc8\x4b\x4d\x4e\x2d\x2e\x4e\x2c\xaa\x54\x28\x28\xca\xcf\x4a\x4d\x2e\xd1\x2d\x4a\x4d\xcc\x29\x49\x4d\x51\xc8\xcc\x4b\xcb\x57\xc8\x48\x2d\x4a\xe5\x02\x04\x00\x00\xff\xff\x41\x19\xbb\xd9\x2c\x00\x00\x00")

func goServiceTemplateReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_goServiceTemplateReadmeMd,
		"go-service-template/README.md",
	)
}

func goServiceTemplateReadmeMd() (*asset, error) {
	bytes, err := goServiceTemplateReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "go-service-template/README.md", size: 44, mode: os.FileMode(420), modTime: time.Unix(1552344077, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _goServiceTemplateDeployReadmeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x0a\xc8\x49\x4c\x4e\x55\x48\xcc\xab\x54\x48\x49\x2d\xc8\xc9\xaf\xcc\x4d\xcd\x2b\xd1\x2d\x4a\xcd\x49\x2c\x49\x4d\x51\xc8\xc9\x4f\xcf\x4c\x56\xc8\x48\x2d\x4a\xe5\x02\x04\x00\x00\xff\xff\x39\x5a\x8f\xad\x28\x00\x00\x00")

func goServiceTemplateDeployReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_goServiceTemplateDeployReadmeMd,
		"go-service-template/deploy/README.md",
	)
}

func goServiceTemplateDeployReadmeMd() (*asset, error) {
	bytes, err := goServiceTemplateDeployReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "go-service-template/deploy/README.md", size: 40, mode: os.FileMode(420), modTime: time.Unix(1552344077, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _goServiceTemplateDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x08\x28\xca\x2f\xcb\x4c\x49\x55\x48\x54\x48\xc9\x4f\xce\x4e\x2d\xd2\x4d\xce\xcf\x2d\xc8\x2f\x4e\x55\x48\xcb\x2f\x52\x28\xc9\x48\x55\x28\x4e\x2d\x2a\xcb\x4c\x4e\xe5\x02\x04\x00\x00\xff\xff\x3a\x2b\xe6\x96\x2b\x00\x00\x00")

func goServiceTemplateDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_goServiceTemplateDockerComposeYml,
		"go-service-template/docker-compose.yml",
	)
}

func goServiceTemplateDockerComposeYml() (*asset, error) {
	bytes, err := goServiceTemplateDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "go-service-template/docker-compose.yml", size: 43, mode: os.FileMode(420), modTime: time.Unix(1552344077, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _goServiceTemplateMainGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\xc1\x8a\xdc\x30\x0c\x40\xcf\xd6\x57\xa8\x86\x82\x0d\x59\x0f\xbd\xb6\xcc\xa1\x87\x69\x3b\xd0\xee\xa1\x53\xe8\xd9\xeb\x75\x12\x13\xc7\x36\xb2\x12\x76\x29\xfb\xef\xc5\xc9\xec\xf6\x32\xb7\x5e\x02\x52\x24\xbd\x87\xe4\x62\xdd\x64\x07\x8f\xb3\x0d\x09\x20\xcc\x25\x13\xa3\x02\x21\xfb\x99\x25\x08\x99\xeb\xfe\x3d\xd4\x30\x24\x1b\x5b\x50\x9f\xab\xb3\x31\x4a\x00\x21\x87\xc0\xe3\xf2\x60\x5c\x9e\x0f\xc9\x8d\x9e\xd2\x61\xc8\x77\xd5\xd3\x1a\x9c\xbf\x63\x3f\x97\x68\xd9\xdf\xcc\x95\x69\x38\xb8\x18\xe4\x7f\x0f\xa9\xb4\x4a\xd0\x00\x2e\xa7\xba\xa9\xd7\xd5\xdd\xdb\xd9\xe3\x11\xe5\x8d\x9e\xa6\x5d\xc7\x4c\xfc\xcd\xc7\xd2\x6a\xde\x57\x0c\x15\x2d\x1a\x63\xb6\x39\xab\x25\x54\x1a\xa0\x5f\x92\xc3\x90\x02\x2b\x8d\x7f\x40\xb8\x18\xcc\xb9\x45\x20\x5e\x09\x1d\x08\x21\xbf\xe7\x34\xe0\xe8\x63\x91\x2d\xec\x67\x36\x97\x42\x21\x71\xaf\xde\x28\x1d\x5e\x1b\x74\x2b\xa1\x25\x7d\x4e\x8f\xbf\x6d\xe0\x0e\x84\x86\x97\x2b\xe9\x5f\x7a\xe7\x0d\x19\x2b\xad\xe6\xc2\x96\x58\xbd\xf6\x03\x08\x87\x1f\x8f\x38\xdb\xc9\x2b\x37\xda\x84\xb9\x9a\xcb\x76\x99\x0e\x3f\x68\x10\xfb\x95\xcc\x7d\xe6\xd0\x3f\x2b\xd7\xb5\xff\xe7\xc4\x9e\x68\x29\xdc\xe1\xf5\x74\xe6\x72\xfe\xfa\xeb\xf4\xf3\x87\x06\xd1\x67\x42\xb2\x69\xf0\xe8\x1a\x55\xec\xcc\x5c\x94\x06\x21\x1e\xc8\xdb\x09\xc4\xcb\x9b\x64\x7b\x25\xbb\x5e\xe8\xd1\x13\x35\x97\xb6\x98\xd3\x93\x77\x0b\x7b\xa5\x3f\x6d\xd9\x77\x47\x4c\x21\x6e\xf3\xda\x42\xbe\x6c\x0b\x89\x49\x35\x59\x7e\xf4\x44\x5d\x2b\x6b\x84\x5c\xcd\xe9\x29\xb0\x6a\xee\x0d\xf3\x37\x00\x00\xff\xff\x64\x88\xaa\x8e\x8f\x02\x00\x00")

func goServiceTemplateMainGoBytes() ([]byte, error) {
	return bindataRead(
		_goServiceTemplateMainGo,
		"go-service-template/main.go",
	)
}

func goServiceTemplateMainGo() (*asset, error) {
	bytes, err := goServiceTemplateMainGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "go-service-template/main.go", size: 655, mode: os.FileMode(420), modTime: time.Unix(1552344077, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _goServiceTemplatePkgReadmeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x0a\xc8\x49\x4c\x4e\x55\x48\x2c\x28\xc8\xc9\x4c\x4e\x2c\xc9\xcc\xcf\xd3\x2d\x4a\xcd\x49\x2c\x49\x4d\x51\x48\xcf\x57\x28\x48\x4c\xce\x4e\x4c\x4f\x2d\x56\xc8\x48\x2d\x4a\xe5\x02\x04\x00\x00\xff\xff\xcf\x59\x5b\x47\x2b\x00\x00\x00")

func goServiceTemplatePkgReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_goServiceTemplatePkgReadmeMd,
		"go-service-template/pkg/README.md",
	)
}

func goServiceTemplatePkgReadmeMd() (*asset, error) {
	bytes, err := goServiceTemplatePkgReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "go-service-template/pkg/README.md", size: 43, mode: os.FileMode(420), modTime: time.Unix(1552344077, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _goServiceTemplatePkgCliCliGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x91\x41\x6b\xdc\x30\x10\x85\xcf\xd2\xaf\x78\xf8\x50\xec\xb2\xb1\x29\xbd\x2d\xf4\x10\x96\x40\x03\xa1\x87\x96\x9e\x4a\x0f\x5a\x79\xec\x1d\x2a\x4b\x46\x1a\xa7\x09\x61\xff\x7b\x91\xbd\xeb\xae\x29\xf4\xa8\xd1\xd3\x7b\x4f\xdf\x8c\xc6\xfe\x32\x3d\xc1\x3a\xd6\x9a\x87\x31\x44\x41\xa9\x55\xd1\xb3\x9c\xa6\x63\x6d\xc3\xd0\xa4\xb1\xfb\xf0\xb1\xb1\xe1\x18\x4d\xa1\x2b\xad\x9f\x4d\xcc\x12\xeb\xf8\x30\xb4\xf8\x84\x77\xf3\x55\x7d\x08\xc3\x60\x7c\xfb\xa6\x95\xfa\x9e\x68\x0f\xa0\xb0\x8e\x8b\x9d\x56\xea\xdb\x29\x44\xd9\xa3\x78\x78\x19\x43\xa2\x04\xbb\x48\xe1\xd8\x13\xd8\x0b\xc5\xce\x58\xca\xd2\xb3\xd6\x2a\x86\x20\xd9\xf9\xfd\xc6\x37\x27\x37\x0d\x1e\x3d\x0b\xd8\xb3\xb0\x71\x9c\xbd\xe4\x44\xc8\x2f\x56\xd3\xd0\xcd\xb3\x49\xd8\xb1\xbc\xe6\x37\xb8\x43\x4b\x9d\x99\x9c\xdc\x5b\xe1\xe0\xf1\x9b\x9d\xc3\x91\x40\x2f\x64\x27\xa1\x16\xdc\xc1\x07\xa4\xe9\x78\x77\x75\xe1\x84\x9e\x9f\xc9\xeb\x6e\xf2\x76\x4e\x2d\xcd\x38\x7e\x31\x03\x21\x49\x64\xdf\xef\xe0\x82\xef\x3f\x93\x1b\xd7\x41\xca\xdf\xdc\x4c\xb6\xb1\xd9\xaa\xac\x2a\xbc\xfd\xfd\xe3\x7f\xe8\x5d\xe2\x32\xbf\xa7\xe0\xfb\x3d\xd6\xc0\x1b\xa4\x6b\x64\x9e\x7d\x9d\xfc\x7e\x09\xb1\xff\xe0\xdb\xc1\xc4\x3e\xe1\xc7\xcf\xa5\xda\x5c\x42\xa9\x4d\xbf\xb2\xd2\x4a\x9d\xe7\x25\x5c\xfb\xd5\xf7\x6d\x7b\x31\x28\x97\x7d\x57\x5a\xab\xa6\xc1\x72\xb8\xbd\xae\xeb\xba\xd2\xe7\x79\x47\x0f\x0b\xd6\xcc\xd0\x78\x90\x97\xf8\x8a\x31\xb0\x17\x48\xc0\xe1\xe9\x71\x61\x7a\x51\x95\x15\x28\xc6\x10\x67\x2a\x24\x53\xf4\xb8\x86\xaf\x0a\x7d\xd6\x7f\x02\x00\x00\xff\xff\x49\x79\xc8\x91\xa8\x02\x00\x00")

func goServiceTemplatePkgCliCliGoBytes() ([]byte, error) {
	return bindataRead(
		_goServiceTemplatePkgCliCliGo,
		"go-service-template/pkg/cli/cli.go",
	)
}

func goServiceTemplatePkgCliCliGo() (*asset, error) {
	bytes, err := goServiceTemplatePkgCliCliGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "go-service-template/pkg/cli/cli.go", size: 680, mode: os.FileMode(420), modTime: time.Unix(1552345391, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _goServiceTemplatePkgSrvSrvGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\x5d\x4a\x05\x31\x0c\x46\x9f\xed\x2a\x3e\x0a\xc2\x5c\x90\xde\x9d\x88\xe0\x0a\x42\x4d\x6b\xb1\x93\x94\xa4\x33\x22\xe2\xde\xa5\xfe\x3c\xdc\xd7\x84\xef\x9c\x33\x28\xbf\x51\x65\xb8\x9d\x21\xb4\x7d\xa8\x4d\xc4\xae\x35\x86\x70\xbd\xe2\x79\x92\x4d\xd8\x21\x8e\xf9\xca\xa0\xd1\xe0\x6c\x27\x5b\x28\x87\xe4\xdf\xf7\x26\xb4\x33\x7c\x5a\x93\x7a\xc1\x67\xb8\xeb\x5a\xd3\x93\x35\x99\x65\x8b\x8f\xaa\x03\xf7\xbe\xf0\x68\xbe\x48\xd2\xa4\xc6\x07\xac\xd1\x25\x7c\xfd\x59\x74\x60\xb0\x15\xb5\xdd\x41\xbd\x43\x38\xb3\x3b\xd9\x07\x74\xb0\xd1\x6c\x2a\x8e\xa2\x86\x6a\x94\xb9\x1c\xfd\x27\xa3\x65\xc6\x64\x32\xbc\xe8\xbb\xa4\xff\x24\x1d\xdb\x6d\x46\x97\x2d\xae\xf3\x68\x52\x53\x4a\x71\x69\xbf\x03\x00\x00\xff\xff\xac\x43\x32\x89\xf7\x00\x00\x00")

func goServiceTemplatePkgSrvSrvGoBytes() ([]byte, error) {
	return bindataRead(
		_goServiceTemplatePkgSrvSrvGo,
		"go-service-template/pkg/srv/srv.go",
	)
}

func goServiceTemplatePkgSrvSrvGo() (*asset, error) {
	bytes, err := goServiceTemplatePkgSrvSrvGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "go-service-template/pkg/srv/srv.go", size: 247, mode: os.FileMode(420), modTime: time.Unix(1552344077, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _goServiceTemplateScriptsBuildSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\xc1\x0d\x80\x20\x0c\x05\xd0\x3b\x53\x7c\xc3\x00\xac\xe2\x0a\xa5\x34\xf6\x27\x0d\x18\x8b\x07\xb7\xf7\xd5\xa3\x75\xce\x96\x5e\x00\xa0\x54\x9c\x21\x6a\x90\xf9\x21\xd7\xed\xcc\x4d\x95\x6d\x03\xfd\x65\x0c\xc4\xba\xa8\x70\x7b\xac\xfc\x01\x00\x00\xff\xff\x6b\xdc\xfd\xbd\x3a\x00\x00\x00")

func goServiceTemplateScriptsBuildShBytes() ([]byte, error) {
	return bindataRead(
		_goServiceTemplateScriptsBuildSh,
		"go-service-template/scripts/build.sh",
	)
}

func goServiceTemplateScriptsBuildSh() (*asset, error) {
	bytes, err := goServiceTemplateScriptsBuildShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "go-service-template/scripts/build.sh", size: 58, mode: os.FileMode(420), modTime: time.Unix(1552344077, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"go-service-template/Dockerfile": goServiceTemplateDockerfile,
	"go-service-template/Makefile": goServiceTemplateMakefile,
	"go-service-template/README.md": goServiceTemplateReadmeMd,
	"go-service-template/deploy/README.md": goServiceTemplateDeployReadmeMd,
	"go-service-template/docker-compose.yml": goServiceTemplateDockerComposeYml,
	"go-service-template/main.go": goServiceTemplateMainGo,
	"go-service-template/pkg/README.md": goServiceTemplatePkgReadmeMd,
	"go-service-template/pkg/cli/cli.go": goServiceTemplatePkgCliCliGo,
	"go-service-template/pkg/srv/srv.go": goServiceTemplatePkgSrvSrvGo,
	"go-service-template/scripts/build.sh": goServiceTemplateScriptsBuildSh,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"go-service-template": &bintree{nil, map[string]*bintree{
		"Dockerfile": &bintree{goServiceTemplateDockerfile, map[string]*bintree{}},
		"Makefile": &bintree{goServiceTemplateMakefile, map[string]*bintree{}},
		"README.md": &bintree{goServiceTemplateReadmeMd, map[string]*bintree{}},
		"deploy": &bintree{nil, map[string]*bintree{
			"README.md": &bintree{goServiceTemplateDeployReadmeMd, map[string]*bintree{}},
		}},
		"docker-compose.yml": &bintree{goServiceTemplateDockerComposeYml, map[string]*bintree{}},
		"main.go": &bintree{goServiceTemplateMainGo, map[string]*bintree{}},
		"pkg": &bintree{nil, map[string]*bintree{
			"README.md": &bintree{goServiceTemplatePkgReadmeMd, map[string]*bintree{}},
			"cli": &bintree{nil, map[string]*bintree{
				"cli.go": &bintree{goServiceTemplatePkgCliCliGo, map[string]*bintree{}},
			}},
			"srv": &bintree{nil, map[string]*bintree{
				"srv.go": &bintree{goServiceTemplatePkgSrvSrvGo, map[string]*bintree{}},
			}},
		}},
		"scripts": &bintree{nil, map[string]*bintree{
			"build.sh": &bintree{goServiceTemplateScriptsBuildSh, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

