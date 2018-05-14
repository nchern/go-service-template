NAME=my-cool-service
TAG=latest
IMAGE_NAME=$(NAME):$(TAG)
OUT=$(NAME)

# Enrich the commands below with appropriate logic if necessary
# Introducing a shell script(e.g.into ./scripts/ folder) could be concidered as a good practice

.PHONY: build
build:
	 GOOS=linux go build -o bin/$(OUT) .

.PHONY: test
test:
	go test ./...

.PHONY: docker-build
docker-build: build
	docker build -t $(IMAGE_NAME) --build-arg SERVICE_NAME=$(NAME) .
