#!/bin/bash
set -eu
set -o pipefail

INSTALL=
while getopts 'i' OPTS
do
    case "$OPTS" in
        i)  # Install
            INSTALL=1 ;;
    esac
done
shift $((OPTIND - 1))

VERSION="$(git describe --tags --abbrev=0 || :)"
REVISION="$(git rev-parse --short HEAD || :)"
LDFLAGS="-X 'main.version=${VERSION}' -X 'main.revision=${REVISION}'"

if (( INSTALL )); then
    go install -ldflags "$LDFLAGS"
else
    go build -ldflags "$LDFLAGS"
fi
