
.PHONY: vendor

vendor:
	go mod tidy && go mod vendor

build:
	go build -o ./bin/api-service

test:
	go test ./...

fmt:
	go fmt ./...