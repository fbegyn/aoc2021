package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fbegyn/aoc2021/go/helpers"
)

func main() {
	file := os.Args[1]
	lines := helpers.InputToLines(file)
	errScore := 0
	var incomplete []string

	for _, line := range lines {
		var queue []rune
		complete := false

		for _, c := range line {
			var expected rune

			if strings.Contains("({[<", string(c)) {
				queue = append([]rune{closeRune(c)}, queue...)
				continue
			}

			if len(queue) == 0 {
				errScore += errorValue(c)
				complete = true
			}

			expected, queue = queue[0], queue[1:]
			if expected != c {
				errScore += errorValue(c)
				complete = true
			}
		}

		if complete {
			continue
		}

		if len(queue) > 0 {
			incomplete = append(incomplete, string(queue))
		}
	}

	var complScores []int
	for _, option := range incomplete {
		completeScore := 0
		for _, c := range option {
			completeScore *= 5
			completeScore += completeValue(c)
		}
		complScores = append(complScores, completeScore)
	}

	fmt.Printf("solution for part 1: %d\n", errScore)
	fmt.Printf("solution for part 2: %d\n", complScores[len(complScores)/2])
}

func closeRune(r rune) rune {
	switch r {
	case '{':
		return '}'
	case '(':
		return ')'
	case '[':
		return ']'
	case '<':
		return '>'
	default:
		panic(fmt.Errorf("no closing rune found for: %c", r))
	}
}

func errorValue(r rune) int {
	switch r {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	default:
		panic(fmt.Errorf("no value found for: %c", r))
	}
}

func completeValue(r rune) int {
	switch r {
	case ')':
		return 1
	case ']':
		return 2
	case '}':
		return 3
	case '>':
		return 4
	default:
		panic(fmt.Errorf("no value found for: %c", r))
	}
}
