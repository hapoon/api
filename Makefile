BIN := hapoon-api
VERSION := $$(make -s show-version)
CURRENT_REVISION := $(shell git rev-parse --short HEAD)
BUILD_LDFLAGS := "-s -w -X main.revision=$(CURRENT_REVISION)"
GOBIN ?= $(shell go env GOPATH)/bin
export GO111MODULE=on

.PHONY: run
run:
	go run main.go
	
.PHONY: test
test:
	go test -v -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: build
build:
	go build -ldflags=$(BUILD_LDFLAGS) -o $(BIN) .
	
.PHONY: show-version
show-version: $(GOBIN)/gobump
	gobump show -r .

$(GOBIN)/gobump:
	go install github.com/x-motemen/gobump/cmd/gobump@latest

.PHONY: build-image
build-image:
	docker build . --file Dockerfile --tag $(BIN):latest

.PHONY: run-image
run-image:
	docker run --name $(BIN) -d -p 8080:8080 $(BIN)

.PHONY: clean
clean:
	rm -rf $(BIN)
	go clean

.PHONY: gen
gen: oapi-validate
	docker run --rm -v "${PWD}:/local" \
		openapitools/openapi-generator-cli generate \
		-i /local/oas/status.yaml -g go-server \
		-o /local --additional-properties=outputAsLibrary=true,sourceFolder=gen,packageName=api \

.PHONY: oapi-validate
oapi-validate:
	docker run --rm -v "${PWD}:/local" \
		openapitools/openapi-generator-cli validate \
		-i /local/oas/status.yaml

.PHONY: author
author:
	docker run --rm -v "${PWD}:/local" \
		openapitools/openapi-generator-cli author template \
		-g go -o /local/tmp

.PHONY: config-help
config-help:
	docker run --rm -v "${PWD}:/local" \
		openapitools/openapi-generator-cli config-help \
		-g go-server
