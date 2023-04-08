package day11

import (
	"aoc2021/util"
	"fmt"
)

func Part2(inputFile string) string {
	data := octopi(util.GetPlot(inputFile))
	target, count := len(data)*len(data[0]), 1 // start with one to offset that we are not incrementing after the final cycle
	for data.cycle() != target {
		count++
	}
	return fmt.Sprint(count)
}
