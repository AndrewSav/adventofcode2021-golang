package day01

import (
	"aoc2021/util"
	"fmt"
)

func solve(inputFile string, shift int) string {
	lines := util.Atoi(util.ReadInput(inputFile))
	count := 0
	for i, s := range lines {
		if i > shift && lines[i-shift-1] < s {
			count++
		}
	}
	return fmt.Sprintf("%d", count)
}
