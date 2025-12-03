build:
	@go build -o bin/hopkey cmd/hopkey/main.go

fmt:
	@go fmt ./...
