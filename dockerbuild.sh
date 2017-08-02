#!/bin/bash -e

TAG=patients-api:$(git describe --long --always --dirty)
echo Building $TAG

CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -v -o patients-api

docker build -t $TAG .
