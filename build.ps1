#!/usr/bin/env pwsh

go fmt ./...
go generate
go build
