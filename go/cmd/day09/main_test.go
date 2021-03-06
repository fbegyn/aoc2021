package main

import (
	"bufio"
	"testing"

	"github.com/fbegyn/aoc2021/go/helpers"
)

func TestPart1(t *testing.T) {
	file := "../../../inputs/day09/test.txt"
	scanner := bufio.NewScanner(helpers.OpenFile(file))
	heightMap := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		scanned := []int{}
		for _, r := range line {
			scanned = append(scanned, int(r-'0'))
		}
		heightMap = append(heightMap, scanned)
	}

	answer, _ := Part1(heightMap)
	if answer != 15 {
		t.Errorf("Part1 got %d, wants %d", answer, 15)
	}
}

func TestPart2(t *testing.T) {
	file := "../../../inputs/day09/test.txt"
	scanner := bufio.NewScanner(helpers.OpenFile(file))
	heightMap := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		scanned := []int{}
		for _, r := range line {
			scanned = append(scanned, int(r-'0'))
		}
		heightMap = append(heightMap, scanned)
	}

	_, points := Part1(heightMap)
	part2 := Part2(heightMap, points)
	if part2 != 1134 {
		t.Errorf("Part1 got %d, wants %d", part2, 1134)
	}
}
