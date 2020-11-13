.PHONY: all build lint check-fmt vet golangci-lint

GIT_COMMIT := $(shell git rev-parse --short HEAD)
IMAGE_NAME ?= quay.io/opstree/ot-dockerlinter
VERSION := $(shell cat ./VERSION)

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -ldflags "-X main.GitCommit=${GIT_COMMIT} -X main.Version=${VERSION}" -o ot-docker-linter

image:
	docker build -t ${IMAGE_NAME}:${VERSION} -f Dockerfile .

check-fmt:
	test -z "$(shell gofmt -l .)"

lint:
	OUTPUT="$(shell go list ./...)"; golint -set_exit_status $$OUTPUT

vet:
	VET_OUTPUT="$(shell go list ./...)"; GO111MODULE=on go vet $$VET_OUTPUT

test:
	go test -v -coverprofile=coverage.txt ./...

golangci-lint:
	golangci-lint run ./...
