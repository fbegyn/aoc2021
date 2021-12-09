package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/fbegyn/aoc2021/go/helpers"
)

func main() {
	file := os.Args[1]
	scanner := bufio.NewScanner(helpers.OpenFile(file))
	heightMap := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		scanned := []int{}
		for _, r := range line {
			scanned = append(scanned, int(r-'0'))
		}
		heightMap = append(heightMap, scanned)
	}

	part1, points := Part1(heightMap)
	fmt.Printf("solution for part 1: %d\n", part1)
	part2 := Part2(heightMap, points)
	fmt.Printf("solution for part 2: %d\n", part2)
}

func Part1(heightMap [][]int) (sum int, lowPoints [][2]int) {
	ylimit, xlimit := len(heightMap)-1, len(heightMap[0])-1
	for y := range heightMap {
		for x := range heightMap[y] {
			before := true
			if x > 0 {
				before = heightMap[y][x] < heightMap[y][x-1]
			}
			after := true
			if x < xlimit {
				after = heightMap[y][x] < heightMap[y][x+1]
			}
			above := true
			if y > 0 {
				above = heightMap[y][x] < heightMap[y-1][x]
			}
			below := true
			if y < ylimit {
				below = heightMap[y][x] < heightMap[y+1][x]
			}
			if before && after && above && below {
				lowPoints = append(lowPoints, [2]int{y, x})
				sum += 1 + heightMap[y][x]
			}
		}
	}
	return
}

func Part2(heightMap [][]int, points [][2]int) int {
	visited := make([][]bool, len(heightMap))
	surfaces := []int{}
	for i := range visited {
		visited[i] = make([]bool, len(heightMap[0]))
	}
	for _, low := range points {
		surfaces = append(surfaces, Flood(heightMap, visited, low[0], low[1]))
	}

	sort.Ints(surfaces)
	largest := len(surfaces) - 1
	return surfaces[largest] * surfaces[largest-1] * surfaces[largest-2]
}

func Flood(heightMap [][]int, visited [][]bool, y, x int) (surface int) {
	if heightMap[y][x] == 9 {
		return 0
	}

	if heightMap[y][x] < 9 && !visited[y][x] {
		surface += 1
		visited[y][x] = true
	}

	if x > 0 && !visited[y][x-1] {
		surface += Flood(heightMap, visited, y, x-1)
	}
	if x < len(heightMap[0])-1 && !visited[y][x+1] {
		surface += Flood(heightMap, visited, y, x+1)
	}
	if y > 0 && !visited[y-1][x] {
		surface += Flood(heightMap, visited, y-1, x)
	}
	if y < len(heightMap)-1 && !visited[y+1][x] {
		surface += Flood(heightMap, visited, y+1, x)
	}

	return
}
