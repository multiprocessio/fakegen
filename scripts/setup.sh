#!/usr/bin/env bash

set -eux

curl -L -o words.txt https://github.com/dwyl/english-words/blob/master/words.txt?raw=true
python3 scripts/togo.py
rm words.txt
gofmt -w -s ./
