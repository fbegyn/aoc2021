package main

import (
	"testing"

	"github.com/fbegyn/aoc2021/go/helpers"
)

func BenchmarkPart1(b *testing.B) {
	lines := helpers.InputToLines("../../../inputs/day08/input.txt")

	for i := 0; i < b.N; i++ {
		part1(lines)
	}
}

func BenchmarkPart2(b *testing.B) {
	lines := helpers.InputToLines("../../../inputs/day08/input.txt")

	for i := 0; i < b.N; i++ {
		part2(lines)
	}
}
