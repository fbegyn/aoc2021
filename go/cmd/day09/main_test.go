package main

import (
	"testing"

	"github.com/fbegyn/aoc2021/go/helpers"
)

func TestPart1(t *testing.T) {
	file := "../../../inputs/day09/test.txt"
	lines := helpers.InputToLines(file)

	answer := Part1(lines)
	if answer != 15 {
		t.Errorf("Part1 got %d, wants %d", answer, 15)
	}
}
