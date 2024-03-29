package day11

import (
	"aoc2021/util"
	"fmt"
)

func Part1(inputFile string) string {
	data := octopi(util.GetPlot(inputFile))
	count := 0
	for i := 0; i < 100; i++ {
		count += data.cycle()
	}
	return fmt.Sprint(count)
}
