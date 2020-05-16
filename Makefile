.PHONY: run test build build-image run-image

run:
	go run main.go
	
test:
	go test -coverprofile=coverage.txt -covermode=atomic ./...

build:
	GOOS=linux GOARCH=amd64 go build -o build/hapoon-api main.go
	
build-image:
	docker build . --file build/package/Dockerfile --tag hapoon/api:latest

run-image:
	docker run --name hapoon-api -d -p 8080:8080 hapoon/api
