.PHONY: run test build

run:
	go run cmd/hapoon/main.go
	
test:
	go test -cover ./...

build:
	GOOS=linux GOARCH=amd64 go build -o build/hapoon-api cmd/hapoon/main.go
	