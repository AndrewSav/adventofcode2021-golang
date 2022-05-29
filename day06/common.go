package day06

import (
	"aoc2021/util"
	"fmt"
	"strings"
)

func tick(m [9]int) (result [9]int) {
	for i, n := range m {
		if i > 0 {
			result[i-1] += n
		} else {
			result[6] += n
			result[8] += n
		}
	}
	return
}

// This is an optimisation, which in this case was not nescessary at all
// Does the same as calling `tick` 7 times, but faster. Not used
func tick7(m [9]int) (result [9]int) {
	result = m
	for i, n := range m {
		if i > 6 {
			result[i] -= n
			result[i%7] += n
		} else {
			result[i+2] += n
		}
	}
	return
}

func solve(inputFile string, days int) string {
	lines := util.ReadInput(inputFile)
	ff := util.Atoi(strings.Split(lines[0], ","))
	m := [9]int{}
	for _, f := range ff {
		m[f]++
	}
	for i := 0; i < days; i++ {
		m = tick(m)
	}
	count := 0
	for _, f := range m {
		count += f
	}
	return fmt.Sprint(count)
}
