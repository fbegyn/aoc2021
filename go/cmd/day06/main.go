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
	input := strings.Split(lines[0], ",")

	tracker := [9]uint64{}
	for _, nr := range input {
		tracker[helpers.Atoi(nr)]++
	}

	for i := 0; i < 256; i++ {
		nextDay := [9]uint64{}
		for j := 8; j >= 0; j-- {
			nextDay[j] = tracker[(j+1)%9]
			if j == 0 {
				nextDay[6] += tracker[j]
			}
		}
		tracker = nextDay
		if i == 79 {
			count := uint64(0)
			for _, fish := range tracker {
				count += fish
			}
			fmt.Printf("solution for part 1: %d\n", count)
		}
	}

	count := uint64(0)
	for _, fish := range tracker {
		count += fish
	}
	fmt.Printf("solution for part 2: %d\n", count)
}
