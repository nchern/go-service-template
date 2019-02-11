NAME=go-service-template
TAG=latest
IMAGE_NAME=$(NAME):$(TAG)
OUT=$(NAME)

# Enrich the commands below with appropriate logic if necessary
# Introducing a shell script(e.g.into ./scripts/ folder) could be concidered as a good practice

.PHONY: install-deps
install-deps:
	@go get -d ./...

.PHONY: build
build:
	 @go build -o bin/$(OUT) .

.PHONY: install
install: test
	 @go install ./...

.PHONY: test
test:
	@go test ./...

.PHONY: docker-build
docker-build:
	@docker build -t $(IMAGE_NAME) .
