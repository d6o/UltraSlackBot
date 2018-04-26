#!/usr/bin/env sh

go clean

find . -name \*.so -delete || true

rm -f ultraslackbot

