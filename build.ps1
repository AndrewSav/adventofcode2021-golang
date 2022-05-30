#!/usr/bin/env pwsh

go fmt ./...
if ($LASTEXITCODE) { exit $LASTEXITCODE }
go generate
if ($LASTEXITCODE) { exit $LASTEXITCODE }
go build
if ($LASTEXITCODE) { exit $LASTEXITCODE }
