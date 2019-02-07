.PHONY: install-deps
install-deps:
	go get github.com/jteeuwen/go-bindata/...


.PHONY: install
install:
	go-bindata go-service-template/...
	go install .
