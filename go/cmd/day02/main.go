package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fbegyn/aoc2021/go/helpers"
)

func main() {
	file := os.Args[1]

	commands := make(chan string)
	go helpers.StreamLines(file, commands)

	x1, depth1 := 0, 0
	x, depth, aim := 0, 0, 0

	for command := range commands {
		split := strings.Split(command, " ")
		operation, operandStr := split[0], split[1]

		operand, err := strconv.Atoi(operandStr)
		if err != nil {
			panic(err)
		}

		switch operation {
		case "forward":
			x1 += operand
			x += operand
			depth += aim * operand
		case "up":
			depth1 -= operand
			aim -= operand
		case "down":
			depth1 += operand
			aim += operand
		}
	}

	fmt.Printf("Solution for part 1: %d\n", depth1*x1)
	fmt.Printf("Solution for part 2: %d\n", depth*x)
}
