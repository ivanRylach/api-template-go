.PHONY: vendor

vendor:
	go mod tidy ; go mod vendor

build:
	go build -mod=vendor -o ./bin/api-service

test:
	go test ./...

fmt:
	go fmt ./...