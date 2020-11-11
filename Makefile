GIT_COMMIT := $(shell git rev-parse --short HEAD)

build:
	go build -ldflags "-X main.GitCommit=${GIT_COMMIT}" -o dockerfile-inspector
