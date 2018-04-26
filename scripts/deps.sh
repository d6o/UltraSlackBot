#!/usr/bin/env sh

set -e

glide install

cd ./plugins/

for d in *; do
    if [[ -d ${d} ]]; then
        cd ${d}
        if [ -f glide.yaml ]; then
        glide install
        fi
        cd ..
    fi
done
