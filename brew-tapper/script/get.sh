#!/bin/sh
set -e

TAR_FILE="/tmp/tapper.tgz"
RELEASES_URL="https://github.com/softleader/homebrew-tap/releases"
test -z "$TMPDIR" && TMPDIR="$(mktemp -d)"

last_version() {
  curl -sL -o /dev/null -w %{url_effective} "$RELEASES_URL/latest" |
    rev |
    cut -f1 -d'/'|
    rev
}

download() {
  test -z "$VERSION" && VERSION="$(last_version)"
  test -z "$VERSION" && {
    echo "Unable to get tapper version." >&2
    exit 1
  }
  rm -f "$TAR_FILE"
  curl -s -L -o "$TAR_FILE" \
    "$RELEASES_URL/download/$VERSION/tapper.tgz"
}

download
tar -xf "$TAR_FILE" -C "$TMPDIR"
"${TMPDIR}/tapper" "$@"