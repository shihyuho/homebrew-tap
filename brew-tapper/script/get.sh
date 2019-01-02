#!/bin/sh
set -e

TAR_FILE="/tmp/tapper.tgz"
RELEASES_URL="http://ci.softleader.com.tw:8082/tapper"
test -z "$TMPDIR" && TMPDIR="$(mktemp -d)"

download() {
  rm -f "$TAR_FILE"
  curl -s -L -o "$TAR_FILE" "$RELEASES_URL"
}

download
tar -xf "$TAR_FILE" -C "$TMPDIR"
"${TMPDIR}/tapper" "$@"