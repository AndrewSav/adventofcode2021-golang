package day24

import (
	"aoc2021/util"
)

func Part2(inputFile string) string {
	inp := [14]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	return solve(util.ReadInput(inputFile), inp)
}
