package main

import (
	"os"

	"github.com/fbegyn/aoc2021/go/helpers"
)

func main() {
	file := os.Args[1]
	input := helpers.OpenFile(file)
	defer input.Close()
}

func part1() {
}

func part2() {
}
