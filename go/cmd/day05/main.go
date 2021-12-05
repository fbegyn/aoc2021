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
		startpoint, endpoint := helpers.ParsePoint(split[0]), helpers.ParsePoint(split[1])

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

		for {
			frequency[*startpoint]++
			if startpoint.X == endpoint.X && startpoint.Y == endpoint.Y {
				break
			}
			startpoint.Move(direction)
		}
	}

	overlap := 0
	for _, count := range frequency {
		if count > 1 { overlap++ }
	}
	return overlap
}

