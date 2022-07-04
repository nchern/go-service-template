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
// go-service-template/pkg/cli/serve.go
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

var _goServiceTemplateDockerfile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\x4f\x4b\xc3\x40\x10\xc5\xef\xf3\x29\x06\x7a\xd1\x43\x12\x8b\x17\x11\xbc\xa8\x11\x8a\x34\x09\xb1\xa2\x3d\x6e\x36\xe3\x66\x74\xff\xb1\x3b\x15\x3f\xbe\x44\x2b\xa8\xe4\xb8\xfb\x1e\xf3\x7e\xbf\xbb\xbe\xdd\xa2\x09\x56\x79\x73\xb9\x2e\xd7\xe7\xa8\x32\x0e\x07\xb6\x23\x25\x80\xa7\xb6\xbf\xbf\xdd\xf4\x58\x99\x50\xe5\xa4\x2b\xc3\x32\x1d\x86\x52\x07\x57\x79\x3d\x51\xf2\x95\x09\x45\xa6\xf4\xce\x9a\x0a\x21\x17\xad\x12\x5a\xfa\x83\x9b\xb6\xdb\x63\x89\x25\xc0\x0a\xaf\xe7\xf3\x28\x13\x61\x4c\xe1\x95\xb4\x20\xfb\xcc\x23\xa1\xf2\xc8\x5e\x28\x39\x1a\x59\x09\xa1\x0e\x5e\x14\x7b\x4a\xd0\x3f\x36\xe8\xd4\x1b\xcd\x4d\x51\xd6\x16\x23\xc5\xfc\xf3\x00\xf8\x2f\xf1\x8b\x1c\xbe\xa7\x8b\xe2\x25\x05\x77\x75\x34\xfb\x12\x1a\x78\x11\x1f\x17\xf9\x61\x85\x5b\xc5\x1e\x8f\x01\xc6\x90\x04\xea\xe7\xae\x7d\xa8\xf1\xe2\x6c\x8e\x4f\x42\x14\x0e\x5e\xd9\x53\x74\x24\x89\x75\x46\xfa\x88\x21\xb3\x37\x7f\xdb\x6b\x80\xba\xd9\xf5\xfb\xae\xdd\x34\xbb\xe5\xb5\xcf\x00\x00\x00\xff\xff\x84\xf3\x8f\xa0\x95\x01\x00\x00")

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

	info := bindataFileInfo{name: "go-service-template/Dockerfile", size: 405, mode: os.FileMode(420), modTime: time.Unix(1577372654, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _goServiceTemplateMakefile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x92\x31\x8f\xdb\x30\x0c\x85\xe7\xe8\x57\x3c\xe0\x32\x24\x83\x9d\xdd\xc0\x01\x09\xd0\x20\xbd\xe1\xce\x1d\x7c\x43\xa7\x42\x91\x78\xb2\x50\x59\x34\x24\x3a\x87\xfc\xfb\x42\x4e\xd2\x2b\xdc\x85\x26\x1f\x3e\xf2\x91\xb0\xde\x0e\xaf\xc7\x67\xc7\x55\xa6\x74\xf1\x86\x2a\xa1\x61\x0c\x5a\x48\x75\x87\xd3\x73\x49\xb2\xa8\x97\xd7\xc3\xe9\xf8\x6b\x26\xd7\x9b\xf2\xd9\x36\xeb\x4d\x77\x38\x6d\x55\xfb\xde\x3d\x24\xa5\x9e\xd0\xb5\xdf\xda\x06\xc7\x98\xbc\xe9\x21\x3d\xc1\xf0\x30\xe8\x68\x33\xce\x14\xf8\x13\x9f\x5e\x7a\xe8\x71\x4c\x3c\x26\xaf\x85\x10\xd8\x79\x03\xff\x81\x48\x86\x72\xd6\xe9\xfa\x77\xca\x4b\x94\xc4\x76\x32\x3e\x3a\x68\xe4\x9e\x42\x40\x36\xc9\x8f\xb2\xa1\xda\xd5\x3e\x0a\xa3\xde\xdd\x94\xbc\xc3\x07\x07\x4b\x69\x0b\xc3\x53\xb0\x38\x17\xeb\x68\xbc\xa5\x44\x16\x3a\x43\xc3\x31\x5b\x8c\x49\x1b\xf1\x86\x94\xaa\x7f\x7c\x6f\xdf\x7e\x36\xf0\x31\x8b\x0e\xa1\xb2\x34\x66\xf5\x6f\xd1\xa8\xd5\x13\xa6\x4c\xb8\xf2\x94\x60\x7a\xce\x14\x61\x69\xa4\x68\x29\x9a\x2b\x06\x1d\xb5\xa3\x81\xa2\x20\x5f\xb3\xd0\xa0\x56\x7b\xc7\x70\x24\xa8\x2c\xea\x5d\x5d\xd7\x5f\x2e\xc1\x47\x51\x25\x34\x6a\x85\xbd\xe3\x92\x2e\x99\x0b\x89\xba\x90\xdc\xe1\x15\x1c\x17\x69\x49\x9d\x27\x1f\xac\x9a\xe3\xad\x65\x9e\x77\x93\x51\x31\xce\x3e\xee\xd6\x9b\xf6\xbd\xdb\xa2\xfe\xef\xca\xc7\x81\x0d\xe6\xff\x7a\x6b\xbd\x6b\x4b\xa3\x99\x28\xe1\x6e\x53\xd0\x52\xa2\x4a\xda\xd0\x92\x36\x1c\x45\xfb\x48\xa9\xf2\x83\x76\xa4\x16\x75\xa3\x56\x7b\xcb\xe6\x37\xa5\xc7\xa6\x82\xf5\xe6\xeb\x5d\x95\x65\xff\x04\x00\x00\xff\xff\x7a\x74\xd1\x88\x8a\x02\x00\x00")

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

	info := bindataFileInfo{name: "go-service-template/Makefile", size: 650, mode: os.FileMode(420), modTime: time.Unix(1644671047, 0)}
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

var _goServiceTemplateMainGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xcd\xcd\x6a\x85\x30\x10\xc5\xf1\x75\xe6\x29\xa6\x81\x42\x02\x55\xf7\x16\x97\xed\x7b\xa4\xc3\x24\x0e\xe6\x8b\x18\xcb\x85\x8b\xef\x7e\xd1\xf5\x5d\x1d\xf8\xf3\x83\x53\x1d\x6d\x2e\x30\x26\x27\x19\x40\x52\x2d\xad\xa3\x01\xa5\x63\x09\x1a\x40\xe9\x20\x7d\x3d\xfe\x46\x2a\x69\xca\xb4\x72\xcb\x53\x28\xc3\xce\xed\x5f\x88\x87\xce\xa9\x46\xd7\xf9\x6d\xab\x5b\x98\x28\x8a\x06\x0b\xe0\x8f\x4c\xf7\x87\xb1\xf8\x04\x25\x1e\xb9\x35\x9c\x17\xa4\x28\xe3\xcf\x83\xe9\xe8\x6c\xec\xf7\x5d\x3f\x16\xcc\x12\x2f\xa6\x62\x09\xe3\xaf\xeb\x2e\x7a\xa3\xfd\xb5\x33\x7e\xee\xfa\xeb\x62\x16\xd4\x09\x27\xbc\x02\x00\x00\xff\xff\x9e\x11\x36\x51\xc0\x00\x00\x00")

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

	info := bindataFileInfo{name: "go-service-template/main.go", size: 192, mode: os.FileMode(420), modTime: time.Unix(1644061206, 0)}
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

var _goServiceTemplatePkgCliCliGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\x4f\x4b\xc3\x30\x18\x87\xcf\x79\x3f\xc5\x8f\x1e\xa4\x95\xad\x45\xbc\x15\x76\x1a\x03\x85\xe1\x41\xf1\x24\x1e\xb2\x2c\xeb\x82\xcd\x1f\xde\xa4\x43\x19\xfd\xee\x92\xd9\x4d\x45\xbc\x86\x87\xdf\xf3\xbc\x09\x52\xbd\xc9\x4e\x43\xf5\x86\xc8\xd8\xe0\x39\xa1\x24\x51\x74\x26\xed\x87\x4d\xad\xbc\x6d\x62\xd8\xdd\xdc\x36\xca\x6f\x58\x16\x54\x11\x29\xef\xe2\x09\x92\x21\x3c\x48\xab\xb1\x40\xd1\xf9\x79\xd4\x7c\x30\x4a\xcf\x93\xb6\xa1\x97\x49\x17\x44\x22\xee\x3d\xa7\x3b\xdd\x87\x7f\x18\x98\x08\x89\xba\xae\x33\xdc\x7b\xd7\x9d\xd9\xb5\x77\x1d\xf6\xba\x0f\x27\xe3\x41\x72\xf6\xb1\xf7\x69\x69\xb7\xb8\x3e\xb5\xd4\x4b\x6f\xad\x74\xdb\x0c\xec\x06\xa7\x60\x9c\x49\x65\x85\xe3\x37\xb8\xc0\xd5\x2f\xf4\x48\x42\x3c\x47\xdd\x02\x98\xda\x67\x24\x44\x76\xb5\xc0\x59\x9f\x9f\x9e\x72\x77\x8b\x4b\x7e\x7e\x7b\x1c\xdc\xaa\x45\x36\x95\xea\x4f\xc4\x0c\x92\xbb\x88\x97\xd7\x98\xd8\xb8\xae\x82\x66\xf6\x9c\x5b\x84\x60\x9d\x06\x76\x50\x76\x5b\xe7\xad\xb2\x22\x21\xc6\x19\x89\x91\x46\xa2\xa6\xc1\xea\x5d\xab\x61\xfa\x0b\x07\xed\x12\x7f\x20\x78\xe3\x12\x92\xc7\x72\x7d\xff\x75\xdd\x44\x95\x3f\xa6\xa7\xe1\xe9\xda\xfa\x42\xd0\x48\x9f\x01\x00\x00\xff\xff\xc6\x5a\xde\x0d\xd7\x01\x00\x00")

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

	info := bindataFileInfo{name: "go-service-template/pkg/cli/cli.go", size: 471, mode: os.FileMode(420), modTime: time.Unix(1644061486, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _goServiceTemplatePkgCliServeGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x53\x41\x6b\xdc\x3c\x10\x3d\x4b\xbf\x62\xf0\xe1\xc3\xfe\x70\x6c\x42\x6f\x86\x1c\x42\x48\xcb\x1e\x9a\x43\x36\x39\x95\x1e\x66\xe5\xb1\x2d\xd6\x96\xd4\x19\x79\x9b\x50\xf6\xbf\x17\xc9\x9b\xc0\x42\x2f\x86\xd1\x3c\xbd\x37\xef\x59\x13\xd0\x1c\x71\x24\x30\xb3\xd5\xda\x2e\xc1\x73\x84\x52\xab\xc2\x4b\x91\xbf\xad\xd8\xd1\xe1\x9c\x0a\x79\x17\x83\xf3\x5c\x68\xad\x8a\xd1\xc6\x69\x3d\x34\xc6\x2f\xad\x33\x13\xb1\x6b\x47\x7f\x23\xc4\x27\x6b\xe8\x26\xd2\x12\x66\x8c\xf4\xcf\xb3\x70\x1c\x5b\xe1\x53\x71\x4d\x22\x61\xb8\xfd\xd2\x1a\x7f\x60\x2c\x74\xa5\x75\xdb\xc2\xcb\x64\x05\xac\x00\x3a\xa0\x37\x5c\xc2\x4c\xe0\x07\xc0\x34\x29\xc8\x7a\x30\x7e\x59\xd0\xf5\x1d\x24\x05\x02\x5e\x9d\x40\x9c\x68\x2b\x19\xa2\x07\x34\x86\x42\x84\xe0\x45\xec\x61\x26\x60\xfa\xb5\x92\x44\x49\xe4\xbb\x08\x38\x8b\x87\xe4\x58\x48\x00\xe1\x80\x62\x0d\x48\xe4\xd5\xc4\x95\x09\xbc\xcb\x4a\x17\x19\xe9\xd2\xad\x1b\xf0\x8e\x60\xb0\x33\x41\x20\xfe\xe8\x6d\x1d\xec\xfb\xac\x7f\x39\x4c\x03\xb0\xf7\xf1\xb3\xb6\x0e\xac\xb3\x11\x86\xd5\x99\x68\x7d\xaa\x80\xd0\x4c\x9f\x80\x44\xdb\x68\x7d\x42\x4e\x3f\x20\xdb\x78\x58\x7a\xb8\x83\xff\x72\x2c\xcd\xc3\x86\xfb\xa3\x95\x7a\x15\xea\x00\xa0\xc8\xa0\xa2\xd6\x4a\xed\x27\xcf\xb1\x83\xe2\x2a\x06\x6b\x08\x12\x73\x86\x09\x0c\x9e\xe9\x37\x71\xc6\xb7\x2d\xdc\xf3\x28\x1d\xc0\x9e\x22\x20\x8f\x52\x03\x35\x63\x03\x9b\xd6\xe3\x1b\x9a\x98\x00\xe5\x6d\x55\xeb\x0d\xbf\xb7\x33\x39\x43\x8f\xcc\x9e\xa5\x83\xc8\x2b\xd5\x57\x9d\x57\xc1\x31\xcd\xb5\x75\xb4\x52\xcf\xab\xeb\xb2\xdf\xd2\x2c\x3d\xfc\x7f\x65\xa3\xce\xa2\xf0\xe3\xa7\x44\xb6\x6e\xac\x20\xf9\xda\x5c\x7f\xf5\x4c\x27\xe2\xb2\xd2\x4a\x9d\x6b\xad\xce\xe9\x41\x24\x9a\x1c\x60\x99\xa1\x29\xda\x87\xa5\x6f\xee\xfb\xfe\x42\x58\x7e\x24\x56\xe9\xf3\x05\x7e\xcd\x96\xae\x8d\x1e\x84\x4f\xcd\x3e\x22\xc7\x12\x43\x78\xc2\x85\x2a\xad\x95\x81\xee\x0e\x16\x3c\x52\x69\x26\x74\xe0\xa5\xd9\xe7\x67\x5f\xc3\x6d\xa5\xd5\xb6\x02\xcd\x93\x8f\x76\x78\x2f\x4d\x9d\xfa\x3b\x17\x89\x79\x0d\xb1\x86\xcb\x5e\x34\xfb\xdd\xb7\x97\xc7\xe7\xef\x95\x56\x83\x67\x60\x74\x69\xad\xb2\xaf\x4d\xd3\x87\x6c\xe9\xc0\x84\xc7\xe4\xea\xac\xff\x06\x00\x00\xff\xff\x68\xf7\x08\x40\x7d\x03\x00\x00")

func goServiceTemplatePkgCliServeGoBytes() ([]byte, error) {
	return bindataRead(
		_goServiceTemplatePkgCliServeGo,
		"go-service-template/pkg/cli/serve.go",
	)
}

func goServiceTemplatePkgCliServeGo() (*asset, error) {
	bytes, err := goServiceTemplatePkgCliServeGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "go-service-template/pkg/cli/serve.go", size: 893, mode: os.FileMode(420), modTime: time.Unix(1644061923, 0)}
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

var _goServiceTemplateScriptsBuildSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\xd4\x4f\xca\xcc\xd3\x2f\xce\xe0\xe2\x52\x56\x08\xc8\x49\x4c\x4e\x55\x48\xcc\xab\x54\x28\xce\x2f\xc8\xc8\x2c\x2e\xc9\x4c\x4e\x2c\x49\x4d\x51\x48\x2a\xcd\xcc\x49\x51\xc8\xc9\x4f\xcf\x4c\x56\xc8\x48\x2d\x4a\xe5\x02\x04\x00\x00\xff\xff\xc8\x02\xdd\x43\x36\x00\x00\x00")

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

	info := bindataFileInfo{name: "go-service-template/scripts/build.sh", size: 54, mode: os.FileMode(420), modTime: time.Unix(1644165859, 0)}
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
	"go-service-template/pkg/cli/serve.go": goServiceTemplatePkgCliServeGo,
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
				"serve.go": &bintree{goServiceTemplatePkgCliServeGo, map[string]*bintree{}},
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

