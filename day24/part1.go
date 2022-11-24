package day24

import (
	"aoc2021/util"
)

func Part1(inputFile string) string {
	inp := [14]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	return solve(util.ReadInput(inputFile), inp)
}
