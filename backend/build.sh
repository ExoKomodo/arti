#! /bin/bash

set -e

pushd "$(dirname "${BASH_SOURCE[0]}")"

go build -mod=mod -v -o ../dist/backend ./cmd/base/main.go

popd
