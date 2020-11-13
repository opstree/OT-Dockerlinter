# Base image for building binary
FROM golang:1.15 as builder
LABEL VERSION=v0.1 \
      ARCH=AMD64 \
      DESCRIPTION="A dockerfile linting tool" \
      MAINTAINER="OpsTree Solutions"
WORKDIR /go/src/dockerfile-inspector
COPY go.mod go.mod
RUN go mod download
COPY . /go/src/dockerfile-inspector
RUN GIT_COMMIT=$(git rev-parse --short HEAD) && \
    VERSION=$(cat ./VERSION) && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -ldflags "-X main.GitCommit=${GIT_COMMIT} -X main.Version=${VERSION}" -o ot-docker-linter

# Copying binary to distroless
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /go/src/dockerfile-inspector/ot-docker-linter .
USER nonroot:nonroot
ENTRYPOINT ["/ot-docker-linter"]
