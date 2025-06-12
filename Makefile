GOCMD=go
GOTEST=$(GOCMD) test
BUILD_DIR=build
BINARY_NAME=gotcher

VERSION := $(shell git describe --tags --always --dirty 2>/dev/null)
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_DATE := $(shell date -u +"%Y-%m-%d-%H:%M:%S")
LDFLAGS := -X 'github.com/yairp7/gotcher/internal/version.Version=$(VERSION)' \
           -X 'github.com/yairp7/gotcher/internal/version.GitCommit=$(COMMIT)' \
           -X 'github.com/yairp7/gotcher/internal/version.BuildDate=$(BUILD_DATE)'

run:
	$(GOCMD) run .

build:
	mkdir -p $(BUILD_DIR)/
	$(GOCMD) build -ldflags="$(LDFLAGS)" -o $(BUILD_DIR)/$(BINARY_NAME) main.go

clean:
	rm -fr ./$(BUILD_DIR)

test: 
	$(GOTEST) -v ./... $(OUTPUT_OPTIONS)

docker-build:
	docker build --rm -t $(BINARY_NAME) .

docker-run:
	docker run -p ${PORT}:${PORT} -t $(BINARY_NAME)

.PHONY: all test build

.PHONY: install
install:
	$(GOCMD) install -ldflags="$(LDFLAGS)"

.PHONY: release
release:
	@if [ ! "$(TAG)" ]; then echo "TAG is required - example: make release TAG=v1.0.0"; exit 1; fi
	@echo "Creating new release $(TAG)"
	git tag -a $(TAG) -m "Release $(TAG)"
	git push origin $(TAG)

.DEFAULT_GOAL := build