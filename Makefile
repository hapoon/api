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
