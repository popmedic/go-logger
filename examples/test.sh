#!/usr/bin/env bash

set -e

echo -e "starting tests..."
go test github.com/popmedic/go-logger/log/...  -race -coverprofile=coverage.txt -covermode=atomic
echo -e "tests done."