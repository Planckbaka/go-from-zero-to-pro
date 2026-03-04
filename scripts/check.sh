#!/bin/bash

# check.sh - Run all checks locally
# Usage: ./scripts/check.sh

set -e

echo "🔧 Running go fmt..."
go fmt ./...

echo "🔍 Running go vet..."
go vet ./...

echo "🧪 Running tests..."
go test ./... -v

echo "📦 Running go build..."
go build ./...

echo "✅ All checks passed!"
