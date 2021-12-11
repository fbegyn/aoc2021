package main

import (
	"fmt"
	"os"

	"github.com/fbegyn/aoc2021/go/helpers"
)

func main() {
	file := os.Args[1]
	lines := helpers.InputToLines(file)

	energyLevels := [100]int{}
	for y, line := range lines {
		for x, r := range line {
			energyLevels[y*10+x] = int(r - '0')
		}
	}
	levels2 := [100]int{}
	copy(levels2[:], energyLevels[:])

	part1 := 0
	run(energyLevels[:], func(step, flashes int) bool {
		part1 += flashes

		return step == 99
	})

	var part2 int
	run(levels2[:], func(step, flashes int) bool {
		if flashes == 100 {
			part2 = step + 1
			return true
		}
		return false
	})

	fmt.Printf("solution for part 1: %d\n", part1)
	fmt.Printf("solution for part 2: %d\n", part2)
}

var coord = [8][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
	{1, 1},
	{-1, -1},
	{1, -1},
	{-1, 1},
}

func run(energyLevels []int, stop func(step, flashes int) bool) int {
	for i := 0; ; i++ {
		// increase energy levels by 1
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				energyLevels[y*10+x]++
			}
		}

		// levels > 9 flash
		flashes := 0
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				if energyLevels[y*10+x] > 9 {
					flashes += flash(energyLevels, x, y)
				}
			}
		}

		// all levels that flashes reset to 0
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				if energyLevels[y*10+x] <= -1 {
					energyLevels[y*10+x] = 0
				}
			}
		}

		if stop(i, flashes) {
			return flashes
		}
	}
}

func flash(levels []int, x, y int) int {
	// if already flashed, return 0
	if levels[y*10+x] == -1 {
		return 0
	}

	// flash the current octopus
	flashes := 1
	levels[y*10+x] = -1

	for _, sel := range coord {
		xsel, ysel := x+sel[0], y+sel[1]
		// if we go out of range, skip
		if xsel < 0 || xsel >= 10 || ysel < 0 || ysel >= 10 {
			continue
		}

		// if the selected octopus already flashed, skip
		if levels[ysel*10+xsel] == -1 {
			continue
		}

		// increase the level of the octopus by 1
		levels[ysel*10+xsel]++

		// if the octopus energy level goes > 9, flash it as well
		if levels[ysel*10+xsel] > 9 {
			flashes += flash(levels, xsel, ysel)
		}
	}
	return flashes
}

func Part2() int {
	return 0
}
