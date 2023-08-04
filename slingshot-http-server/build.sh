#!/bin/bash
export TAG="v0.0.0"
#env CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o slingshot-${TAG}-darwin-arm64
#env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o slingshot-${TAG}-darwin-amd64
env CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o slingshot-${TAG}-linux-arm64
#env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o slingshot-${TAG}-linux-amd64
