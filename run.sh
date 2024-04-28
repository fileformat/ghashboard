#!/bin/bash
#
# script to run on localhost
#

set -o errexit
set -o pipefail
set -o nounset

echo "INFO: starting at $(date -u +%Y-%m-%dT%H:%M:%SZ)"


if [ ! -d "./dist" ]; then
    echo "INFO: creating dist directory"
    mkdir ./dist
fi

if [ -f "./dist/ghashboard" ]; then
    echo "INFO: removing old build of ghashboard"
    rm ./dist/ghashboard
fi

COMMIT=$(git rev-parse HEAD)-local
LASTMOD=$(date -u +%Y-%m-%dT%H:%M:%SZ)
VERSION=local
BUILTBY=run.sh

echo "INFO: building new ghashboard"
go build \
    -ldflags "-X main.commit=$COMMIT -X main.date=$LASTMOD -X main.version=$VERSION -X main.builtBy=$BUILTBY -extldflags '-static'" \
    -o ./dist/ghashboard \
    ./*.go

if [ ! -f "./dist/ghashboard" ]; then
    echo "ERROR: failed to build ghashboard"
    exit 1
fi

echo "INFO: running $(./dist/ghashboard --version)"

if [ -f ".env" ]; then
    echo "INFO: loading environment variables from .env"
    export $(grep -v '^#' .env)
fi

echo "INFO: running ghashboard with arguments: $@"
./dist/ghashboard "$@"

echo "INFO: complete at $(date -u +%Y-%m-%dT%H:%M:%SZ)"
