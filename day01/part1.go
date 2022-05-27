package day01

import (
	"aoc2021/util"
	"fmt"
)

func Part1(inputFile string) string {
	lines := util.Atoi(util.ReadInput(inputFile))
	count := 0
	for i, s := range lines {
		if i > 0 && lines[i-1] < s {
			count++
		}
	}
	return fmt.Sprintf("%d", count)
}
