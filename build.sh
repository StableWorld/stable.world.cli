#!/bin/sh
set -eo pipefail

COMMAND=$1
export GOOS=${2:-$(go env GOOS)}
export GOARCH=${3:-$(go env GOARCH)}

echo "Building ./bin/${COMMAND}-$GOOS-$GOARCH"
go build -o "./bin/${COMMAND}-$GOOS-$GOARCH" "commands/${COMMAND}/main.go"
