#!/usr/bin/env bash
# whopsdk-go's release check: the release workflow (publish-main.yml) runs it, with VERSION and TAG,
# after the tag is released. The tag is the release; this is its distribution: a request for the
# version makes proxy.golang.org fetch it, and go get proves it resolves at the module path.
set -euo pipefail
cd "$(dirname "$0")/.."
export GOPROXY=https://proxy.golang.org GOFLAGS=-mod=mod

# Resolve the version from proxy.golang.org.
module=$(jq -r .package .github/release/release.json)
consumer=$(mktemp -d)
cd "$consumer"
go mod init whopsdkgoconsumer
for attempt in $(seq 1 20); do
  if go get "${module}@${TAG}" > go-get.log 2>&1; then
    echo "${module}@${TAG} resolves from proxy.golang.org after ${attempt} attempt(s)"
    exit 0
  fi
  sleep 15
done
cat go-get.log
echo "::error::${TAG} is tagged, but go get ${module}@${TAG} still fails: proxy.golang.org has not caught up, or cannot read the repository."
exit 1
