package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fbegyn/aoc2021/go/helpers"
)

type board struct {
	numbers [5][5]int
	marked [5][5]bool
}

func main() {
	// throw the input lines into an array, without the empty lines
	file := os.Args[1]
	input := helpers.OpenFile(file)
	lines := []string{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
		    lines = append(lines, scanner.Text())
		}
	}

	// read the first line and parse into int array
	numbersStr := strings.Split(lines[0], ",")
	numbers := make([]int, len(numbersStr))
	for ind, str := range numbersStr {
		number, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		numbers[ind] = number
	}

	 // load in the bingo boards
	boards := make([]board, len(lines[1:])/5)
	current := 0
	for index, line := range lines[1:] {
		// skim the row
	 	row := [5]int{}
	 	split := strings.Fields(line)
	 	for elIndex, elementStr := range split {
	 		nr, err := strconv.Atoi(elementStr)
	 		if err != nil {
	 			panic(err)
	 		}
	 		row[elIndex] = nr
	 	}

		boards[current].numbers[index%5] = row
		// if we fill a board, move along to the next
		if ((index) % 5) == 4 { // 0-based index
			current += 1
		}
	 }

	// time to play bingo!
	winners := make(map[int]bool) // keep track of who already won, so we can skip
	order := make([][2]int, len(boards)) // we need the index of the board as wel as the last number called
	// start looking for winners
	for count := 0; count < len(boards); count++ {
		// start calling off numbers
		for _, call := range numbers {
			for ind := range boards {
				// skip over the already won boards
				if winners[ind] {
					continue
				}

				// play bingo!
				boards[ind].Mark(call)
				if boards[ind].CheckBingo() {
					winners[ind] = true 
					order[count] = [2]int{ind, call}
					count++
				}
			}
		}
	}

	sumFirst := boards[order[0][0]].SumUnmarked()
	fmt.Printf("solution to part 1: %d\n", order[0][1]*sumFirst)
	sumLast := boards[order[len(order)-1][0]].SumUnmarked()
	fmt.Printf("solution to part 1: %d\n", order[len(order)-1][1]*sumLast)
}

func (b* board) Mark(nr int) {
	for yi, y := range b.numbers {
		for xi, x := range y {
			if x == nr { b.marked[yi][xi] = true }
		}
	}
}

func (b *board) CheckBingo() bool {
	for i := 0; i < 5; i++ {
		xbingo, ybingo := true, true
		for j := 0; j < 5; j++ {
			xbingo = xbingo && b.marked[i][j]
			ybingo = ybingo && b.marked[j][i]
		}
		if xbingo || ybingo { return true }
	}
	return false
}

func (b *board) SumUnmarked() (sum int) {
	for y, yv := range b.marked {
		for x, xv := range yv {
			if !xv { sum += b.numbers[y][x] }
		}
	}
	return
}

func (b *board) Print() {
	for _, y := range b.numbers {
		for _, x := range y {
			fmt.Printf("%2d ", x)
		}
		fmt.Println()
	}
	fmt.Println()
}
