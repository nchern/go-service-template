.PHONY: install-deps
install-deps:
	go get github.com/jteeuwen/go-bindata/...

.PHONY: install
install:
	go-bindata -o go-svc-generator/bindata.go ./go-service-template/...
	cd go-svc-generator && go install .
