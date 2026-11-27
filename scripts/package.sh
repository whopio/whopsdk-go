#!/usr/bin/env bash
# whopsdk-go's release build: the release workflow (publish-main.yml) runs it before it tags a version,
# with VERSION and TAG. A Go module is released by its tag, so this checks, tests and consumes the tree
# as a module; it uploads nothing. golangci-lint and Go come from the workflow.
set -euo pipefail
cd "$(dirname "$0")/.."

# Check the module path and the version the client reports.
module=$(jq -r .package .github/release/release.json)
declared=$(go mod edit -json | jq -r .Module.Path)
if [ "$declared" != "$module" ]; then
  echo "::error::go.mod declares ${declared}, but ${VERSION} is released as ${module}."
  exit 1
fi
for stamp in "\"X-Fern-SDK-Version\", \"v${VERSION}\"" "\"User-Agent\", \"${module}/${VERSION}\""; do
  if ! grep -qF "$stamp" core/request_option.go; then
    echo "::error::core/request_option.go does not send ${stamp}."
    exit 1
  fi
done

# Verify the module and its dependencies.
go mod verify
go mod download

# CI, as ci.yml runs it.
bash scripts/ci.sh

# Consume the module as a caller would.
(
  module=$(jq -r .package .github/release/release.json)
  workspace="$PWD"
  probe=$(mktemp -d)
  cd "$probe"
  go mod init whopsdkgoprobe
  go mod edit -replace "${module}=${workspace}" -require "${module}@${TAG}"
  sed "s|MODULE|${module}|g" > main.go <<'GO'
package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	whopsdk "MODULE"
	"MODULE/client"
	"MODULE/option"
)

type capture struct{ header http.Header }

func (c *capture) Do(request *http.Request) (*http.Response, error) {
	c.header = request.Header.Clone()
	return &http.Response{StatusCode: http.StatusOK, Header: http.Header{}, Body: io.NopCloser(strings.NewReader("{}")), Request: request}, nil
}

func main() {
	sent := &capture{}
	whop := client.NewWhop(option.WithToken("probe"), option.WithHTTPClient(sent))
	_, _ = whop.Users.Me(context.Background(), &whopsdk.MeUsersRequest{})
	version, agent := sent.header.Get("X-Fern-SDK-Version"), sent.header.Get("User-Agent")
	fmt.Printf("on the wire: X-Fern-SDK-Version=%q User-Agent=%q\n", version, agent)
	if version != "v"+os.Args[1] || !strings.HasSuffix(agent, "/"+os.Args[1]) {
		os.Exit(1)
	}
}
GO
  go mod tidy
  go run . "$VERSION"
)
