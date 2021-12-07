package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/fbegyn/aoc2021/go/helpers"
)

func main() {
	file := os.Args[1]
	input := helpers.InputToLines(file)
	posStr := strings.Split(input[0], ",")
	positions := make([]int, len(posStr))
	for ind := range posStr {
		positions[ind] = helpers.Atoi(posStr[ind])
	}

	sort.Ints(positions)
	median := positions[len(positions)/2]
	sum := 0
	part1 := int64(0)
	for _, pos := range positions {
		part1 += helpers.Abs(int64(pos - median))
		sum += pos
	}
	mean := sum/len(positions)
	fmt.Printf("solution to part 1: %d\n", part1)

	candidates := [2]int{mean, mean+1}
	cheapest := 100000000000
	for _, cand := range candidates {
		fuel := 0
		for _, pos := range positions {
			diff := int(helpers.Abs(int64(pos - cand)))
			for j := 0; j < diff; j++ {
				fuel += j + 1
			}
		}
		if fuel < cheapest {
			cheapest = fuel
		}
	}
	fmt.Printf("solution to part 2: %d\n", cheapest)
}
