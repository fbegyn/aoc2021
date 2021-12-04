package main

import (
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
	// throw the input into a go channel
	file := os.Args[1]
	input := make(chan string)
	go helpers.StreamLines(file, input)

	// read the first line and parse into int array
	callStr := <-input
	numbersStr := strings.Split(callStr, ",")
	numbers := make([]int, len(numbersStr))
	for ind, str := range numbersStr {
		number, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		numbers[ind] = number
	}

	// throw away the empty line after the first
	<-input

	// load in the bingo boards
	boards := []board{}
	c := 0
	temp := [5][5]int{}
	for line := range input {
		row := [5]int{}
		split := strings.Fields(line)
		j := 0
		for _, st := range split {
			nr, err := strconv.Atoi(st)
			if err != nil {
				panic(err)
			}
			row[j] = nr
			j++
		}

		temp[c] = row
		c++
		if 5 <= c {
			<-input
			t := board{temp, [5][5]bool{}}
			boards = append(boards, t)
			c = 0
		}
	}

	// time to play bingo!
	winners := make(map[int]bool) // keep track of who already won, so we can skip
	order := make([][2]int, len(boards))
	count := 0
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

		// once all boards finished, stop
		if count >= len(boards) {
			break
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
