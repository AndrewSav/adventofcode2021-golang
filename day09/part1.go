package day09

import (
	"aoc2021/util"
	"fmt"
)

func Part1(inputFile string) string {
	data := util.GetPlot(inputFile)
	count := 0
	for y, dd := range data {
		for x, d := range dd {
			lowest := true
			if lowest && x > 0 && dd[x-1] <= d {
				lowest = false
			}
			if lowest && x < len(dd)-1 && dd[x+1] <= d {
				lowest = false
			}
			if lowest && y > 0 && data[y-1][x] <= d {
				lowest = false
			}
			if lowest && y < len(data)-1 && data[y+1][x] <= d {
				lowest = false
			}
			if lowest {
				count += d + 1
			}
		}
	}
	return fmt.Sprint(count)
}
