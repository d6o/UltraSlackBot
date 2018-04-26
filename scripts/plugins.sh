#!/usr/bin/env sh

set -e

find . -name \*.so -delete || true

cd ./plugins/

for d in *; do
     (go build -buildmode=plugin $d/*.go)
done
