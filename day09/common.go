package day09

import (
	"aoc2021/util"
	"strings"
)

type plot [][]int

func getPlot(inputFile string) plot {
	lines := util.ReadInput(inputFile)
	var data = make(plot, len(lines))
	for i, l := range lines {
		data[i] = util.Atoi(strings.Split(l, ""))
	}
	return data
}
