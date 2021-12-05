package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fbegyn/aoc2021/go/helpers"
)

func main() {
	// throw the input lines into an array, without the empty lines
	file := os.Args[1]
	lines := helpers.InputToLines(file)

	fmt.Printf("Solution to part 1: %d\n", checkVents(lines, false))
	fmt.Printf("Solution to part 2: %d\n", checkVents(lines, true))
}

func checkVents(lines []string, diagonal bool) int {
	frequency := make(map[helpers.Point]int)
	for _, line := range lines {
		split := strings.Split(line, " -> ")
		startpoint, endpoint := parsePoint(split[0]), parsePoint(split[1])

		// only do horizontal and vertical lines
		if !diagonal && (startpoint.X != endpoint.X && startpoint.Y != endpoint.Y) {
			continue
		}

		direction := [2]int64{0, 0}
		if startpoint.X < endpoint.X {
			direction[0] = 1
		} else if endpoint.X < startpoint.X {
			direction[0] = -1
		}
		if startpoint.Y < endpoint.Y {
			direction[1] = 1
		} else if endpoint.Y < startpoint.Y {
			direction[1] = -1
		}

		copyStart := startpoint
		for {
			frequency[*copyStart]++
			if copyStart.X == endpoint.X && copyStart.Y == endpoint.Y {
				break
			}
			copyStart.Move(direction)
		}
	}

	overlap := 0
	for _, count := range frequency {
		if count > 1 {
			overlap++
		}
	}
	return overlap
}

func parsePoint(str string) *helpers.Point {
	split := strings.Split(str, ",")
	return helpers.NewPoint(
		int64(helpers.Atoi(split[0])),
		int64(helpers.Atoi(split[1])),
	)
}
