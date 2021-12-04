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
	file := os.Args[1]
	input := make(chan string)
	go helpers.StreamLines(file, input)

	callStr := <-input
	numbersStr := strings.Split(callStr, ",")
	numbers := make([]int, len(numbersStr))
	for ind, str := range numbersStr {
		number, err := strconv.Atoi(str)
		numbers[ind] = number
		if err != nil {
			panic(err)
		}
	}
	<-input

	boards := []board{}

	c := 0
	temp := [5][5]int{}
	skip := false
	for line := range input {
		if skip {
			skip = false
			continue
		}
		row := [5]int{}
		split := strings.Split(line, " ")
		j := 0
		for _, st := range split {
			if st == "" {
				continue
			}
			nr, err := strconv.Atoi(strings.TrimSpace(st))
			if err != nil {
				panic(err)
			}
			row[j] = nr
			j++
		}

		temp[c] = row
		c++
		if 5 <= c {
			skip = true
			t := board{temp, [5][5]bool{}}
			boards = append(boards, t)
			c = 0
		}
	}

	bingo := false
	winners := make(map[int]bool)
	order := make([][2]int, len(boards))
	count := 0

	for _, call := range numbers {
		for ind := range boards {
			if winners[ind] {
				continue
			}
			boards[ind].Mark(call)
			if bingo = boards[ind].CheckBingo(); bingo {
				fmt.Println("BINGO!")
				boards[ind].Print()
				winners[ind] = true 
				order[count] = [2]int{ind, call}
				count++
			}
		}
		if count >= len(boards) {
			break
		}
	}

	sumFirst := boards[order[0][0]].SumUnmarked()
	fmt.Printf("solution to part 1: %d\n", order[0][1]*sumFirst)
	sumLast := boards[order[len(order)-1][0]].SumUnmarked()
	fmt.Printf("solution to part 1: %d\n", order[len(order)-1][1]*sumLast)
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

func (b* board) Mark(nr int) {
	for yi, y := range b.numbers {
		for xi, x := range y {
			if x == nr {
				b.marked[yi][xi] = true
			}
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
		if xbingo || ybingo {
			return true
		}
	}
	return false
}

func (b *board) SumUnmarked() (sum int) {
	for y, yv := range b.marked {
		for x, xv := range yv {
			if !xv {
				sum += b.numbers[y][x]
			}
		}
	}
	return
}
