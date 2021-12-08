package main

import (
	"fmt"
	"math/bits"
	"os"
	"strings"

	"github.com/fbegyn/aoc2021/go/helpers"
)

func main() {
	// throw the input lines into an array, without the empty lines
	file := os.Args[1]
	lines := helpers.InputToLines(file)

	fmt.Printf("solution to par 1: %d\n", part1(lines))
	fmt.Printf("solution to par 2: %d\n", part2(lines))
}

func part1(lines []string) (count int) {
	for _, line := range lines {
		split := strings.Split(line, " | ")
		outputs := strings.Fields(split[1])
		for _, out := range outputs {
			if len(out) == 2 || len(out) == 3 || len(out) == 4 || len(out) == 7 {
				count++
			}
		}
	}
	return
}

func part2(lines []string) (sum int) {
	for _, line := range lines {
		split := strings.Split(line, " | ")

		matchedChars := [10]int{}
		signals := strings.Fields(split[0])
		for _, signal := range signals {
			switch len(signal) {
			case 2: // 1
				matchedChars[1] = charToInt(signal)
			case 4: // 4
				matchedChars[4] = charToInt(signal)
			case 3: // 7
				matchedChars[7] = charToInt(signal)
			case 7: // 8
				matchedChars[8] = charToInt(signal)
			}
		}

		outputs := strings.Fields(split[1])
		for _, output := range outputs {
			payload := charToInt(output)
			sum *= 10
			switch len(output) {
			case 2:
				sum += 1
			case 3:
				sum += 7
			case 4:
				sum += 4
			case 5:
				switch {
				// if the output is 5 long and matches with 1 (a,b)
				// only possible number displayed is a 3
				case payload&matchedChars[1] == matchedChars[1]:
					sum += 3
				// if the field share 3 bits with a 4, it's a 5
				// matched e,f,b => leaves d and c
				case bits.OnesCount(uint(payload&matchedChars[4])) == 3:
					sum += 5
				// match through exclusion
				default:
					sum += 2
				}
			case 6:
				switch {
				// a 6 shares a,b with the number 1 on the display
				case payload&matchedChars[1] != matchedChars[1]:
					sum += 6
				// a 9 shares a,e,f,b with number 4, also exclusion
				case payload&matchedChars[4] == matchedChars[4]:
					sum += 9
				default:
				}
			case 7:
				sum += 8
			}
		}
	}
	return sum
}

func charToInt(field string) (result int) {
	for i := range field {
		result |= 1 << (field[i] - 'a')
	}
	return
}
