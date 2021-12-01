package main

import (
	"fmt"
	"os"

	"github.com/fbegyn/aoc2021/go/helpers"
)

func main() {
	file := os.Args[1]
	input := helpers.OpenFile(file)
	defer input.Close()

	depths, err := helpers.LinesToInts(input)
	if err != nil {
		panic(err)
	}

	sol1, sums := part1(depths)
	fmt.Println(sums)
	sol2, _ := helpers.IncDecCount(sums)

	fmt.Printf("Solution for part 1: %d\n", sol1)
	fmt.Printf("Solution for part 2: %d\n", sol2)
}

func part1(depths []int) (incCount int, sums []int) {
	previous := -1000

	for ind, depth := range depths {
		if ind < len(depths)-2 {
			sums = append(sums, depths[ind]+depths[ind+1]+depths[ind+2])
		}
		if ind == 0 {
			continue
		}
		if previous < depth {
			incCount++
		}
		previous = depth
	}

	return
}
