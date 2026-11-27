#!/usr/bin/env bash
# whopsdk-go's CI: .github/workflows/ci.yml runs it, publish-main.yml runs it before it tags a release,
# and whop-monorepo's SDK pipeline runs it over every build it delivers. golangci-lint, Go and Docker
# come from the caller.
set -euo pipefail
cd "$(dirname "$0")/.."

go build ./...

unformatted=$(gofmt -l .)
if [ -n "$unformatted" ]; then
  echo "::error::These files are not gofmt-clean:"
  printf '%s\n' "$unformatted"
  exit 1
fi

go vet ./...
golangci-lint run

# The generated tests expect a wiremock server, from wiremock/docker-compose.test.yml.
if [ -f wiremock/docker-compose.test.yml ]; then
  compose=(docker compose -p "whopsdk-go-ci-$$" -f wiremock/docker-compose.test.yml)
  trap '"${compose[@]}" down --remove-orphans' EXIT
  "${compose[@]}" up -d --wait
  published="$("${compose[@]}" port wiremock 8080 | head -n1)"
  WIREMOCK_URL="http://localhost:${published##*:}"
  export WIREMOCK_URL
fi
go test ./...
