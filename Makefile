GOCMD=go
GOTEST=$(GOCMD) test
BUILD_DIR=build
BINARY_NAME=gotcher

run:
	$(GOCMD) run .

build:
	mkdir -p $(BUILD_DIR)/
	$(GOCMD) build -ldflags="-s -w" -gcflags=all="-l" -o $(BUILD_DIR)/$(BINARY_NAME) main.go

clean:
	rm -fr ./$(BUILD_DIR)

test: 
	$(GOTEST) -v ./... $(OUTPUT_OPTIONS)

docker-build:
	docker build --rm -t $(BINARY_NAME) .

docker-run:
	docker run -p ${PORT}:${PORT} -t $(BINARY_NAME)

.PHONY: all test build