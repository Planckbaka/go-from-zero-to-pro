#!/usr/bin/env bash
set -euo pipefail

go_files=$(find . -type f -name '*.go' -not -path './vendor/*')

if [[ -n "$go_files" ]]; then
  unformatted=$(printf '%s\n' "$go_files" | xargs gofmt -l)
  if [[ -n "$unformatted" ]]; then
    echo "These files are not gofmt-formatted:"
    echo "$unformatted"
    exit 1
  fi
fi

go test ./...
