#! /usr/bin/env sh

DAY=$1
DIR=day$(printf "%02d" ${DAY})

mkdir -p ./inputs/${DIR}/

curl \
  -fSL -o ./inputs/${DIR}/input.txt \
  -H "Cookie: ${AOCCOOKIE}" \
  https://adventofcode.com/2020/day/${DAY}/input
