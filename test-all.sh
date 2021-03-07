#!/bin/bash
go mod verify || exit 1
go mod tidy -v || exit 1
depm list --json | docker run --rm -i sonatypecommunity/nancy:latest sleuth -n || exit 1
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.38.0 golangci-lint run --enable gosec ./... || exit 1
go test ./...
