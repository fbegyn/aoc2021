#!/usr/bin/env bash

days=()

for file in ./bin/*; do
    day=$(basename $file)
    days+=("${file} ../inputs/${day}/input.txt")
done

echo $days

hyperfine --warmup 10 --export-csv aoc2020.csv "${days[@]}"
