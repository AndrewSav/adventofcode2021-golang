package day03

import (
	"aoc2021/util"
	"fmt"
)

func Part1(inputFile string) string {
	lines := util.ReadInput(inputFile)
	dim := len(lines[0])
	var gamma, epsilon int
	for pos := 0; pos < dim; pos++ {
		gamma <<= 1
		epsilon <<= 1
		ones, zeroes := split(lines, pos)
		if len(ones) > len(zeroes) {
			gamma++
		} else {
			epsilon++
		}
	}
	return fmt.Sprint(gamma * epsilon)
}
