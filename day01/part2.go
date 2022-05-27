package day01

import (
	"aoc2021/util"
	"fmt"
)

func Part2(inputFile string) string {
	lines := util.Atoi(util.ReadInput(inputFile))
	count := 0
	for i, s := range lines {
		if i > 2 && lines[i-3] < s {
			count++
		}
	}
	return fmt.Sprintf("%d", count)
}
