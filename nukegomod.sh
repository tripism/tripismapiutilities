#!/bin/bash
export GOPRIVATE=github.com/tripism
rm -rf go.mod
rm -rf go.sum
go mod init
go build
go test
go mod tidy

