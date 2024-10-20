NAME       := ecmanager
PACKAGE    := github.com/ESELDevelopment/$(NAME)
VERSION := make

format:
	@gofmt -s -w .

start:
	@go run main.go

test:
	@go test ./...

build_version:
	@go build -ldflags="-w -s -X ${PACKAGE}/cmd/info.version=${VERSION}" ${PACKAGE}
