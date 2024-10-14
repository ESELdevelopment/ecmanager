format:
	@gofmt -s -w .

start:
	@go run main.go

test:
	@go test ./...
