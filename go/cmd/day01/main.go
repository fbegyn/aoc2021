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

	previous, previousSum := 0, 0
	sum := 0
	incCount1, incCount2 := 0, 0

	for ind, depth := range depths {
		if ind < len(depths)-2 {
			sum = depths[ind] + depths[ind+1] + depths[ind+2]
		}

		// skip the rest for the first element, not relevant
		if ind == 0 {
			previous = depth
			previousSum = sum
			continue
		}

		// check if the previous depth is smaller than the current
		if previous < depth {
			incCount1++
		}

		// check if the previous sliding window sum is smaller than the current
		if previousSum < sum {
			incCount2++
		}
		previous = depth
		previousSum = sum
	}

	fmt.Printf("Solution for part 1: %d\n", incCount1)
	fmt.Printf("Solution for part 2: %d\n", incCount2)
}
