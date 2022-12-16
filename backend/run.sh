#! /bin/bash

set -e

pushd "$(dirname "${BASH_SOURCE[0]}")"

go run -mod=mod -v ./cmd/base/main.go

popd
