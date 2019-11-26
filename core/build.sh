#!/bin/bash

set -e  # Exit on error.

NOW="_test_AA/$(date +%Y%m%d%H%M%S).dir"
mkdir -p "$NOW"
mv -iv a_*.go "$NOW" || :
[ -x ../joker ] && cp -aiv ../joker "$NOW"
(git log -n 1; git status) > "$NOW/git.txt"
ln -sfTv "$NOW" _test_AA/LATEST

time=$(which time)

set -x  # Echo commands

# Build gen_data before generating code that would otherwise be
# unnecessarily compiled into it.
$time go build -o gen_data/gen_data gen_data/gen_data.go

$time go run gen_code/gen_code.go

$time ./gen_data/gen_data

$time go fmt a_*.go

JUSTVET=true

$JUSTVET && (cd ..; $time go vet ./...)

$JUSTVET || (cd ..; KEEP_A_FILES=true $time ./run.sh "$@")
