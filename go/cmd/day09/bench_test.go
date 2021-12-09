package main

import (
	"bufio"
	"testing"

	"github.com/fbegyn/aoc2021/go/helpers"
)

func BenchmarkPart1(b *testing.B) {
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

	for i := 0; i < b.N; i++ {
		Part1(heightMap)
	}
}

func BenchmarkPart2(b *testing.B) {
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
	for i := 0; i < b.N; i++ {
		Part2(heightMap, points)
	}
}
