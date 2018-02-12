#!/usr/bin/env bash
set -e
mkdir -p bin
go build -o bin/example cmd/main.go