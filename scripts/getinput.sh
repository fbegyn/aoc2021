#! /usr/bin/env sh

DAY=$1
DIR=day$(printf "%02d" ${DAY})

mkdir -p ./inputs/${DIR}/

curl "https://adventofcode.com/2021/day/${DAY}/input" \
  -H 'authority: adventofcode.com' \
  -H 'pragma: no-cache' \
  -H 'cache-control: no-cache' \
  -H 'upgrade-insecure-requests: 1' \
  -H "referer: https://adventofcode.com/2021/day/${DAY}" \
  -H "cookie: ${AOCCOOKIE}" \
  --compressed \
  -fSL -o ./inputs/${DIR}/input.txt \
