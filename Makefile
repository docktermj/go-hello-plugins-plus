# Makefile that builds go-plugins, a "go" program.

# PROGRAM_NAME is the name of the GIT repository.
PROGRAM_NAME := $(shell basename `git rev-parse --show-toplevel`)
DOCKER_IMAGE_NAME := local/$(PROGRAM_NAME)
DOCKER_CONTAINER_NAME := $(PROGRAM_NAME)
BUILD_VERSION := $(shell git describe --always --tags --abbrev=0 --dirty)
BUILD_TAG := $(shell git describe --always --tags --abbrev=0)
BUILD_ITERATION := $(shell git log $(BUILD_TAG)..HEAD --oneline | wc -l)

.PHONY: help
help:
	@echo "Build $(PROGRAM_NAME) version $(BUILD_VERSION)-$(BUILD_ITERATION)".
	@echo "To build, run 'make build'"
	@echo "All targets:"
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | xargs


.PHONY: docker-build
docker-build:
	docker build \
		--build-arg PROGRAM_NAME=$(PROGRAM_NAME) \
		--build-arg BUILD_VERSION=$(BUILD_VERSION) \
		--build-arg BUILD_ITERATION=$(BUILD_ITERATION) \
		--tag $(DOCKER_IMAGE_NAME) \
		.


.PHONY: build
build: clean docker-build
	mkdir -p ./target || true
	docker create \
		--name $(DOCKER_CONTAINER_NAME) \
		$(DOCKER_IMAGE_NAME)
	docker cp $(DOCKER_CONTAINER_NAME):/output/. ./target/
	docker cp $(DOCKER_CONTAINER_NAME):/root/gocode/bin/$(PROGRAM_NAME) ./target/
	docker rm --force $(DOCKER_CONTAINER_NAME)


.PHONY: docker-run
docker-run:
	docker run \
	    --interactive \
	    --tty \
	    --name $(DOCKER_CONTAINER_NAME) \
	    $(DOCKER_IMAGE_NAME)


.PHONY: clean
clean:
	docker rm --force $(DOCKER_CONTAINER_NAME) || true
	rm -rf ./target
