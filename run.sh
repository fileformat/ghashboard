#!/bin/bash
#
# script to run on localhost
#

set -o errexit
set -o pipefail
set -o nounset

if [ -f ".env" ]; then
  export $(cat .env)
fi

rm -f test.md
#go run ghashboard.go repos.go workflows.go templates.go test.md

go run *.go Dashboard.md


