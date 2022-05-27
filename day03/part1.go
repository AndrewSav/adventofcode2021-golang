package day03

import (
	"aoc2021/util"
	"fmt"
)

func Part1(inputFile string) string {
	lines := util.ReadInput(inputFile)
	dim := len(lines[0])
	counters := make([]int, dim)
	for _, l := range lines {
		for pos := 0; pos < dim; pos++ {
			if l[pos] == '1' {
				counters[pos]++
			}
		}
	}
	var gamma, epsilon int
	for pos := 0; pos < dim; pos++ {
		gamma *= 2
		epsilon *= 2
		if counters[pos] > len(lines)/2 {
			gamma++
		} else {
			epsilon++
		}
	}
	return fmt.Sprint(gamma * epsilon)
}
