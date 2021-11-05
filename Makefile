GOPATH:=$(shell go env GOPATH)
ORG_NAME?=ohmygrpc
SERVICE_NAME?=echo

.PHONY: build
## build: build the application(api)
build:
	go build -o bin/${SERVICE_NAME} cmd/main.go

.PHONY: run
## run: run the application(api)
run:
	go run -v -race cmd/main.go

.PHONY: format
## format: format files
format:
	@go get golang.org/x/tools/cmd/goimports
	goimports -local github.com/${ORG_NAME} -w .
	gofmt -s -w .
	go mod tidy

.PHONY: test
## test: run tests
test:
	@go get github.com/rakyll/gotest
	gotest -p 1 -race -cover -v ./...

.PHONY: coverage
## coverage: run tests with coverage
coverage:
	@go get github.com/rakyll/gotest
	gotest -p 1 -race -coverprofile=coverage.txt -covermode=atomic -v ./...

.PHONY: lint
## lint: check everything's okay
lint:
	@go get github.com/kyoh86/scopelint
	golangci-lint run ./...
	scopelint --set-exit-status ./...
	go mod verify

.PHONY: generate
## generate: generate source code for mocking
generate:
	@go get golang.org/x/tools/cmd/stringer
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen
	go generate ./...

.PHONY: help
## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':'
