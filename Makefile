build:
	@go build -o bin/gochat ./cmd/server

run: build
	@./bin/gochat -config config/local.yaml

test:
	@go test -v ./...
