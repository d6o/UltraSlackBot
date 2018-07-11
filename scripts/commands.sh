#!/usr/bin/env sh

set -e

cd ./commands/

find . -name \*.so -delete || true

for d in *; do
     (go build -buildmode=plugin $d/*.go)
done

cd ..