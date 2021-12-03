package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fbegyn/aoc2021/go/helpers"
)

func main() {
	file := os.Args[1]
	input := helpers.OpenFile(file)
	defer input.Close()

	reportsStr := helpers.LinesToStrings(input)
	reports := make([]uint, len(reportsStr))
	for ind, report := range reportsStr {
		parse, err := strconv.ParseInt(report, 2, 64)
		if err != nil {
			panic(err)
		}
		reports[ind] = uint(parse)
	}

	gamma := uint(0)
	size := len(reportsStr[0])
	mask := uint(0)
	for i := uint(0); i < uint(size); i++ {
		mask = mask | (0b1 << i)
	}
	for i := size - 1; i >= 0; i-- {
		ones, zeros, _, _ := OnesZerosCount(reports, i)
		if ones > zeros {
			gamma = gamma | (0b1 << i)
		}
	}

	epsilon := ^gamma & mask

	fmt.Printf("gamma:   %d\n", gamma)
	fmt.Printf("epsilon: %d\n", epsilon)
	fmt.Printf("Solution for part 1: %d\n", gamma*epsilon)

	oxygen := CalcRating(reports, size-1, false)
	co2 := CalcRating(reports, size-1, true)
	fmt.Printf("oxygen:  %d\n", oxygen)
	fmt.Printf("co2:     %d\n", co2)
	fmt.Printf("Solution for part 2: %d\n", oxygen*co2)
}

func OnesZerosCount(inputs []uint, n int) (zeros, ones int, oneSl, zeroSl []uint) {
	for _, inp := range inputs {
		cmp := inp >> n
		cmp &= 0b1
		if cmp == 0b1 {
			ones++
			oneSl = append(oneSl, inp)
		} else {
			zeros++
			zeroSl = append(zeroSl, inp)
		}
	}
	return
}

func CalcRating(reports []uint, n int, co2 bool) uint {
	zeros, ones, oneSl, zeroSl := OnesZerosCount(reports, n)

	if co2 {
		if len(zeroSl) == 1 {
			return zeroSl[0]
		}

		if zeros < ones {
			return CalcRating(zeroSl, n-1, co2)
		}
		if ones < zeros {
			return CalcRating(oneSl, n-1, co2)
		}
		return CalcRating(zeroSl, n-1, co2)
	} else {
		if len(oneSl) == 1 {
			return oneSl[0]
		}

		if zeros < ones {
			return CalcRating(oneSl, n-1, co2)
		}
		if ones < zeros {
			return CalcRating(zeroSl, n-1, co2)
		}
		    return CalcRating(oneSl, n-1, co2)
	}
}
